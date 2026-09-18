package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := "Hello word"

	fmt.Println(greetings)

	fmt.Println("Length of greetings:", len(greetings))

	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "hi"))

}
