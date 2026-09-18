package main

import "fmt"

func main() {

	// Array example
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr, "longuer est ", len(arr))

	names := [3]string{"Alice", "Bob", "Charlie"}
	//we can update the table bob with Rami
	names[1] = "Rami"
	fmt.Println("Names:", names, "longuer est ", len(names))

	//same here
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array1:", arr1, "longuer est ", len(arr1))

	//	we can also update elements in arr1
	arr1[2] = 10
	fmt.Println("Updated Array1:", arr1, "longuer est ", len(arr1))

	//slice exemple

	scores := []int{12, 15, 20}
	fmt.Println("Scores:", scores, "longuer est ", len(scores))

	scores[1] = 18
	fmt.Println("Updated Scores:", scores, "longuer est ", len(scores))

	scores = append(scores, 90)
	fmt.Println("Appended Scores:", scores, "longuer est ", len(scores))
}
