package main

import (
	"fmt"
)

func sayHello(n string) {
	fmt.Printf("Hello %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func cycleName(n []string, f func(string)) {

	for _, v := range n {
		f(v)
	}
}

func main() {

	cycleName([]string{"Rami", "Alice", "Bob"}, sayHello)
	cycleName([]string{"Rami", "Alice", "Bob"}, sayBye)

}
