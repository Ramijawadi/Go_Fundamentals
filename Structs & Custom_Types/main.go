package main

import "fmt"

func main() {

	mybill := newBill("Myy Bill") //appel instance de la facture
	fmt.Println(mybill.format())

}
