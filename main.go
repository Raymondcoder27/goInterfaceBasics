package main

import "math"

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
