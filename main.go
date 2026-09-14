package main

import "fmt"

func main() {

	var name string = "Rami"
	age := 30
	var score float64 = 95.5666
	fmt.Println(name, "and his age is ", age, "and his score is ", score)
	// Rami and his age is  30 and his score is  95.5666
	/** Printf example */
	fmt.Printf("Name: %s, Age: %d, Score: %.2f\n", name, age, score)
	/*Name: Rami, Age: 30, Score: 95.57*/

	var str = fmt.Sprintln("My name is ", name, "and my age is ", age, "and my score is ", score)
	fmt.Print(str)
	/*how to save the formatted string into a variable str using Sprintln*/
}
