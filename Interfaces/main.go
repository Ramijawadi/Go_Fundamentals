package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumference() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

//square methods for area and circumference

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return 4 * s.length
}

// circle methods for area and circumference
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {

	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Circumference: %.2f\n", s.circumference())
}

func main() {
	shapes := []shape{
		square{length: 5},
		circle{radius: 3},
		circle{radius: 2.5},
		square{length: 2},
	}
	for _, s := range shapes {
		printShapeInfo(s)
		fmt.Println("---")
	}
}
