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

// square methods
func (s square) circumference() float64 {
	return s.length * 4
}
func (s square) area() float64 {
	return s.length * s.length
}

// circle methods
func (c circle) circumference() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is %0.2f \n", s, s.area())
	fmt.Printf("Circumeference of %T is %0.2f \n", s, s.circumference())
}

func main(){
	var shapes []shape{
		square{length: 13.4},
		circle{radius: 5.6},
		circle{radius: 21},
		square{length: 4}
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("-----------")
	}
}