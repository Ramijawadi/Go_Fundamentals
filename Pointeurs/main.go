package main

import "fmt"

func updateName(x *string) {
	*x = "rami"
}

func main() {

	name := "sara"
	// updateName(name)

	m := &name

	// fmt.Println("adress de m =", m) // l'adress memoire de name

	// fmt.Println("the value of m is ", *m) //pointeur sur l adress donne la valeur
	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}
