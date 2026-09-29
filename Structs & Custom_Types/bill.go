package main

type bill struct {
	name  string
	items map[string]float64
	paye  float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		paye:  0,
	}
	return b
}
