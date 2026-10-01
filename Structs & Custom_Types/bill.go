package main

import "fmt"

type bill struct { // creation structurée d'une facture
	name  string
	items map[string]float64
	paye  float64
}

func newBill(name string) bill { //creation d'une nouvelle facture
	b := bill{
		name:  name,
		items: map[string]float64{"bread": 6.5, "cake": 3.6},
		paye:  0,
	}
	return b
}

// format retourne une représentation formatée de la facture
func (b bill) format() string {

	resultat := "Fcature details : \n"
	var total float64 = 0

	//list itmes

	for k, v := range b.items { //parcourir items de la facture

		resultat += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//total

	resultat += fmt.Sprintf("%-25v ...%.2f\n", "Total:", total)
	return resultat

}
