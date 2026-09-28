package main

import "fmt"

var points = []int{10, 20, 80, 50}

func Name(n string) {
	fmt.Println("Hello", n)
}

func CalcScore() {

	fmt.Println("the score is ", score) //same here we receive score form main.go and then call the function in main.go to see the result
}
