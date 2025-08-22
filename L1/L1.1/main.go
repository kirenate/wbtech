package main

type Human struct {
	Height float64
	Weight int
}

func (r *Human) BMI() float64 {
	return float64(r.Weight) / (r.Height * r.Height / 100)
}

type Action struct {
	Human
}

func main() {
	act := Action{Human{
		Height: 175,
		Weight: 70,
	}}
	println(act.BMI())
}
