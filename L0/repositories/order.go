package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Model struct {
	OrderUID string    `json:"order_uid" grom:"primary_key"`
	Order    *Order    `json:"order"`
	Delivery *Delivery `json:"delivery"`
	Payment  *Payment  `json:"payment"`
	Item     *[]Item   `json:"items"`
}

type Order struct {
	OrderUID          string    `json:"order_uid" grom:"primary_key"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	ID       uuid.UUID `json:"id,omitempty" grom:"primary_key"`
	OrderUID string    `json:"order_uid"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Zip      string    `json:"zip"`
	City     string    `json:"city"`
	Address  string    `json:"address"`
	Region   string    `json:"region"`
	Email    string    `json:"email"`
}

type Payment struct {
	ID           uuid.UUID `json:"id,omitempty" grom:"primary_key"`
	OrderUID     string    `json:"order_uid"`
	Transaction  string    `json:"transaction"`
	RequestId    string    `json:"request_id"`
	Currency     string    `json:"currency"`
	Provider     string    `json:"provider"`
	Amount       int       `json:"amount"`
	PaymentDt    int       `json:"payment_dt"`
	Bank         string    `json:"bank"`
	DeliveryCost int       `json:"delivery_cost"`
	GoodsTotal   int       `json:"goods_total"`
	CustomFee    int       `json:"custom_fee"`
}

type Item struct {
	ID          uuid.UUID `json:"id,omitempty" grom:"primary_key"`
	OrderUID    string    `json:"order_uid"`
	ChrtId      int       `json:"chrt_id"`
	TrackNumber string    `json:"track_number"`
	Price       int       `json:"price"`
	Rid         string    `json:"rid"`
	Name        string    `json:"name"`
	Sale        int       `json:"sale"`
	Size        string    `json:"size"`
	TotalPrice  int       `json:"total_price"`
	NmId        int       `json:"nm_id"`
	Brand       string    `json:"brand"`
	Status      int       `json:"status"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOrderTX(ctx context.Context, req *Model) error {

	order := req.Order
	delivery := req.Delivery
	payment := req.Payment
	items := req.Item

	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Table("order").Clauses(clause.OnConflict{DoNothing: true}).Create(&order).Error
		if err != nil {
			return errors.Wrap(err, "failed to save order")
		}

		err = tx.WithContext(ctx).Table("delivery").Clauses(clause.OnConflict{DoNothing: true}).Create(&delivery).Error
		if err != nil {
			return errors.Wrap(err, "failed to save delivery")
		}

		err = tx.WithContext(ctx).Table("payment").Clauses(clause.OnConflict{DoNothing: true}).Create(&payment).Error
		if err != nil {
			return errors.Wrap(err, "failed to save payment")
		}

		err = tx.WithContext(ctx).Table("item").Clauses(clause.OnConflict{DoNothing: true}).Create(&items).Error
		if err != nil {
			return errors.Wrap(err, "failed to save items")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to execute save transaction")
	}

	return nil
}

func (r *Repository) GetOrderTX(ctx context.Context, orderUID string) (*Model, error) {
	req := &Model{OrderUID: orderUID}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var order *Order
		err := tx.WithContext(ctx).Raw(`SELECT * FROM "order" o WHERE o.order_uid = ?`, orderUID).
			Find(&order).Error
		if err != nil {
			return errors.Wrap(err, "failed to find order information")
		}

		var delivery *Delivery
		err = tx.WithContext(ctx).Raw(`SELECT * FROM delivery d WHERE d.order_uid = ?`, orderUID).
			Find(&delivery).Error
		if err != nil {
			return errors.Wrap(err, "failed to find delivery information")
		}

		var payment *Payment
		err = tx.WithContext(ctx).Raw(`SELECT * FROM payment p WHERE p.order_uid = ?`, orderUID).
			Find(&payment).Error
		if err != nil {
			return errors.Wrap(err, "failed to find payment information")
		}

		var items *[]Item
		err = tx.WithContext(ctx).Raw(`SELECT * FROM item i WHERE i.order_uid = ?`, orderUID).
			Find(&items).Error
		if err != nil {
			return errors.Wrap(err, "failed to find item information")
		}

		req.Order = order
		req.Delivery = delivery
		req.Payment = payment
		req.Item = items

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find order model")
	}

	return req, nil
}
