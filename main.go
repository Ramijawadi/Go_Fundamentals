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

	//you can use this link to find all formatting verbs: https://pkg.go.dev/fmt

	//ints :
	var ageOne int = 31
	var ageTwo = 30
	ageThree := 32

	fmt.Println("the result is ", ageOne, ageTwo, ageThree)

	//bit and memories  int8/int16/int32/int64/int/uint8/uint16/uint32/uint64/uint
	//each type has its own range and usage

	var price1 int8 = 20
	var price2 int16 = 300
	var price3 int32 = 40000
	var price4 int64 = 5000000
	fmt.Println("the prices are ", price1, price2, price3, price4)

	var score1 float32 = 12.5
	var score2 float64 = 188.4 //most useful type more precise
	score3 := 1.5
	fmt.Println("the scores are ", score1, score2, "and score 3 egale", score3)
}
