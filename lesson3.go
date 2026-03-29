package main

import "fmt"

func СheckAge(age int, name string) {
	if age >= 18 {
		fmt.Printf("Можете проходить в наш бар, %s \n", name)
	} else if age <= 18 && age >= 0 {
		fmt.Println("Тебе пока рано")
	} else {
		fmt.Println("Вы врете")
	}

}

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scanln(&age)
	СheckAge(age, name)
}
