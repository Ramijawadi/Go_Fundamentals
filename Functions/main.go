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

func main() {

	sayHello("Rami")
	sayBye("Rami")

}
