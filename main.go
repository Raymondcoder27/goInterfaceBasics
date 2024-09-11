package main

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

}
