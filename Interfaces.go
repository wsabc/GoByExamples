package main

import (
	"fmt"
	"math"
)

func main() {
	r := rect{width: 10, height: 5}
	measure(r)

	c := circle{radius: 4}
	// only implement all methods in interface,
	// a struct can be called a implement of one interface
	//measure(c)
	fmt.Println(c.area())
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return (r.height + r.width) * 2
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}
