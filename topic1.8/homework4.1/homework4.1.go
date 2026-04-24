package main

import "fmt"

func divide(a, b float64) (float64, error) {
	err := fmt.Errorf("Знаменатель не может быть равен нулю")

	if b == 0 {
		return 0, err
	}

	return a / b, nil
}

func main() {
	var a, b float64

	fmt.Println("Напишите числитель:")
	fmt.Scanln(&a)
	fmt.Println("Напишите знаменатель:")
	fmt.Scanln(&b)

	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Результат деления: %f", res)
	}
}
