package main

import "fmt"

func Mult(a, b int) int {
	return a * b
}

func main() {
	p1, p2 := 2, 10

	c := Mult(p1, p2)

	fmt.Println(p1, "*", p2, "=", c)
}
