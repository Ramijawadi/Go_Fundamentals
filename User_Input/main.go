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

	reader := bufio.NewReader(os.Stdin) //NewReader pour les informations ; Stdin lire depuis l'entrée standard
	// fmt.Printf("enter the facture name :  ")
	// name, _ := reader.ReadString('\n') //lire ce que l utulisateur tape et save dans name
	// name = strings.TrimSpace(name)     //supprimer les espaces blancs autour du nom
	name, _ := getUserInput("enter the facture name :  ", reader)
	b := newBill(name)

	fmt.Println("the name of the bill is  ", b.name)
	return b

}

func main() {
	mybill := createBill()
	fmt.Println(mybill)

}
