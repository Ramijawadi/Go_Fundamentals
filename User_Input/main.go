package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Println(name, price)

	case "c":
		paye, _ := getUserInput("enter the paye ($):  ", reader)
		fmt.Println(paye)
	case "s":
		fmt.Println("Save facture selected")

	default:
		fmt.Println("Invalid option")
		Promptoption(b)
	}

}

func main() {
	mybill := createBill()
	Promptoption(mybill)
	fmt.Println(mybill)

}
