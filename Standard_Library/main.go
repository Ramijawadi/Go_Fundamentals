package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "Hello my friend"

	fmt.Println(greetings)

	fmt.Println("Length of greetings:", len(greetings))

	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "hi"))
	fmt.Println(strings.ToUpper(greetings))    // to majuscule = HELLO MY FRIEND
	fmt.Println(strings.Index(greetings, "m")) // l index = 6
	fmt.Println(strings.Split(greetings, " ")) //[Hello my friend]

	fmt.Println("original value is ", greetings)

	ages := []int{75, 20, 90, 10, 30}
	sort.Ints(ages) //method sort the table of ages  to [10 20 30 75 90]
	fmt.Println(ages)

	index := sort.SearchInts(ages, 75)
	fmt.Println(index) //index = 3

	names := []string{"Salah", "Alex", "Yuchi", "Bali"}
	sort.Strings(names)
	fmt.Println(names)
}
