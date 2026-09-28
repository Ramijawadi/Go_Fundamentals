package main

import "fmt"

func updateName(n string) string {

	n = "David"
	return n //with this return, the updated value can be captured
}

func updateList(x map[string]float64) { //update the  menu directly

	x["coffee"] = 25.33

}

func main() {
	//with group A type ,strings , ints , bools , floats, arrays,structs
	name := "Alice"

	fmt.Println(name)
	//without return, the original variable remains unchanged
	name = updateName(name)

	fmt.Println("this is the value after update:", name)

	//group B type , maps, slices,  functions
	menu := map[string]float64{
		"tea":   5.0,
		"milk":  3.0,
		"sugar": 2.0,
	}

	fmt.Println("this is the list before update:", menu)
	updateList(menu)
	fmt.Println("this menu after update", menu)
}
