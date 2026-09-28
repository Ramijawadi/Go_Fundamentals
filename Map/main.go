package main

import "fmt"

func main() {

	menu := map[string]float64{ //parcourir une carte en boucle avec range
		//[string]float64 sont les types de [clé] et de valeur du map
		"Burger":  5.99,
		"Fries":   2.99,
		"Soda":    1.99,
		"Salad":   4.99,
		"Dessert": 3.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["Burger"])

	//looping map

	for k, v := range menu {

		fmt.Println(k, "-", v)
	}

	//int as key in map

	phonenumber := map[int]string{

		1234567890: "Rami",
		8596668588: "Alice",
		9891234567: "Bob",
	}

	fmt.Println(phonenumber)
	fmt.Println(phonenumber[1234567890])

	phonenumber[9891234567] = "Wassim"
	fmt.Println(phonenumber)

	phonenumber[8596668588] = "SARA"
	fmt.Println(phonenumber)

}
