package main

import (
	"fmt"
)

func main() {

	age := 45
	fmt.Println(age < 45)
	fmt.Println(age > 45)
	fmt.Println(age == 45)
	fmt.Println(age != 45)

	if age < 30 {
		fmt.Println(" the age is less than 30")
	} else if age < 45 {
		fmt.Println(" the age is less than 45")
	} else if age == 45 {
		fmt.Println(" the age is equal to 45 ")

	} else {
		fmt.Println("the age is more than 45")

	}
}
