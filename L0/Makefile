dev:
	docker compose up kafka kafka-ui postgres -d
	docker compose exec kafka kafka-topics.sh --create --topic events --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092
run:
	docker compose up -d --build
	docker compose exec kafka kafka-topics.sh --create --topic events --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092