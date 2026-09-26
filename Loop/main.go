package main

import (
	"fmt"
)

func main() {

	// x := 0
	// for x < 7 {
	// 	fmt.Println("la valeur de x =", x)
	// 	x++
	// }

	// for i := 0; i < 7; i++ {
	// 	fmt.Println("la valeur de i = ", i)
	// }

	names := []string{"rami", "alex", "bruno", "amine"}

	// for i := 0; i < len(names); i++ {

	// 	fmt.Println("la valeur des names = ", names[i])
	// }

	for index, value := range names {

		fmt.Printf("la index %v et la valeur  %v \n", index, value)
	}

}
