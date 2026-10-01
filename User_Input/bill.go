package main

import (
	"fmt"
	"os"
)

type bill struct { // creation structurée d'une facture
	name  string
	items map[string]float64
	paye  float64
}

func newBill(name string) bill { //creation d'une nouvelle facture
	b := bill{
		name:  name,
		items: map[string]float64{},
		paye:  0,
	}
	return b
}

// format retourne une représentation formatée de la facture
func (b *bill) format() string {

	resultat := "Fcature details : \n"
	var total float64 = 0

	//list itmes

	for k, v := range b.items { //parcourir items de la facture

		resultat += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//Afficher the pay

	resultat += fmt.Sprintf("%-25v ...%.0f\n", "paye:", b.paye)

	//total

	resultat += fmt.Sprintf("%-25v ...%.2f\n", "Total:", total+b.paye)
	return resultat

}

//function to update the bill 'facture

func (b *bill) updateItem(paye float64) {
	b.paye = paye

}

//function de add item to the bill facture

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save facture dans un fichier
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("../bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("the bill was saved to file ")
}
