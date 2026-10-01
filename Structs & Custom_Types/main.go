package main

import "fmt"

func main() {

	mybill := newBill("Myy Bill") //appel instance de la facture
	mybill.addItem("milk", 4.5)
	mybill.addItem("eggs", 2.5)
	mybill.addItem("meat", 12.0)
	mybill.updateItem(10)

	fmt.Println(mybill.format())

}
