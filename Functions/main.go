package main

import (
	"fmt"
	"strings"
)

// func sayHello(n string) {
// 	fmt.Printf("Hello %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Bye %v \n", n)
// }

// func cycleName(n []string, f func(string)) {

// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 { //float64 is the return type
// 	return math.Pi * r * r // use math library to calculate the area of a circle
// }

func getInitials(n string) (string, string) {

	s := strings.ToUpper(n)        //change to uppercase
	names := strings.Split(s, " ") //split = diviser the string into words by space

	var initials []string //slice to save  the initials of each word

	for _, v := range names { //parcourir chaque mot du nom

		initials = append(initials, v[:1]) //ajouter la première lettre du mot aux initiales

	}

	if len(initials) > 1 {
		return initials[0], initials[1]

	}

	return initials[0], " "
}

func main() {

	// sayHello("mohamed")
	// sayBye("mohamed")

	// cycleName([]string{"Rami", "Alice", "Bob"}, sayHello)
	// cycleName([]string{"Rami", "Alice", "Bob"}, sayBye)

	// a1 := circleArea(7.5) // we pass 7.5 as argument r
	// a2 := circleArea(10.0)

	// fmt.Println(a1, a2)
	// fmt.Printf("Area of circle 1 is %0.3f  and the circle 2 is %0.3f", a1, a2)

	fn1, sn1 := getInitials("Mohamed Rami") //M R
	fmt.Println(fn1, sn1)
	fn2, sn2 := getInitials("Sami alex") //S A
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("Alice") //A
	fmt.Println(fn3, sn3)

}
