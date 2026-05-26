package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func main() {
	c1 := Circle{12}
	c2 := Circle{5}
	r1 := Rectangle{5, 6}
	r2 := Rectangle{15, 64}

	figuers := []Shape{c1, c2, r1, r2}

	for _, user := range figuers {
		fmt.Printf("Площадь: %f\n", user.Area())
	}
}
