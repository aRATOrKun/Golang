package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	a, b := 10, 8

	fmt.Printf("%d + %d = %d\n", a, b, applyOperation(a, b, sum))
	fmt.Printf("%d - %d = %d\n", a, b, applyOperation(a, b, sub))
	fmt.Printf("%d * %d = %d\n", a, b, applyOperation(a, b, mult))
}
