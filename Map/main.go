package main

import "fmt"

func main() {

	menu := map[string]float64{

		"Burger":  5.99,
		"Fries":   2.99,
		"Soda":    1.99,
		"Salad":   4.99,
		"Dessert": 3.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["Burger"])

}
