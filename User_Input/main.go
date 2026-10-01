package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Create function for any input from the user

func getUserInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

//function for input from the user

func createBill() bill {
	//reader est un lecteur locale
	reader := bufio.NewReader(os.Stdin) //NewReader pour les informations ; Stdin lire depuis l'entrée standard
	// fmt.Printf("enter the facture name :  ")
	// name, _ := reader.ReadString('\n') //lire ce que l utulisateur tape et save dans name
	// name = strings.TrimSpace(name)     //supprimer les espaces blancs autour du nom
	name, _ := getUserInput("enter the facture name :  ", reader)
	b := newBill(name)

	fmt.Println("the name of the bill is  ", b.name)
	return b

}

func Promptoption(b bill) {

	reader := bufio.NewReader(os.Stdin)
	option, _ := getUserInput("choose option ( a - Create facture , s- save facture, c - add paye ): ", reader)
	fmt.Println("selected option:", option)

	switch option {
	case "a":
		name, _ := getUserInput("enter the facture name :  ", reader)
		price, _ := getUserInput("enter the price :  ", reader)

		p, err := strconv.ParseFloat(price, 64) //parsing string input to float64
		if err != nil {
			fmt.Println("the price must be a valid number")
			Promptoption(b)
		}

		b.addItem(name, p)

		fmt.Println("items added - ", name, price)
		Promptoption(b)

	case "c":
		paye, _ := getUserInput("enter the paye ($):  ", reader)

		p, err := strconv.ParseFloat(paye, 64)
		if err != nil {
			fmt.Println("the paye must be a valid number")
			Promptoption(b)
		}

		b.updateItem(p)

		fmt.Println("paye added  - ", paye)
		Promptoption(b)

	case "s":

		b.save()
		fmt.Println("the  facture / bill is saved  - ", b.name)

	default:
		fmt.Println("Invalid option")
		Promptoption(b)
	}

}

func main() {
	mybill := createBill()
	Promptoption(mybill)
	// fmt.Println(mybill)

}
