package main

import "fmt"

var score = 55

func main() {

	// since all are in same package main we can pass data from one greetings to another

	for _, v := range points {
		fmt.Println(v)
	}

	Name("Alice")

	CalcScore()
}
