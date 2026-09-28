package main

import (
	"fmt"
)

func main() {

	age := 30
	fmt.Println(age < 45)
	fmt.Println(age > 45)
	fmt.Println(age == 45)
	fmt.Println(age != 45)

	if age < 30 {
		fmt.Println(" the age is less than 30")
	} else if age > 30 {
		fmt.Println(" the age is greater than 30")
	} else if age == 30 {
		fmt.Println(" the age is equal to 30 ")

	} else {
		fmt.Println("the age is more than 45")

	}

	names := []string{"rami", "alex", "bruno", "amine"}

	for index, value := range names {

		if index == 1 {

			fmt.Println("Continuing at position:", index)
			continue
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
