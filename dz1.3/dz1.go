package main

import "fmt"

func bubbleSort(arr []int) []int {
	for {

		cnt := 0

		for i := range 4 {
			if arr[i] < arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			} else {
				cnt++
			}
		}

		if cnt == 4 {
			return arr
		}

	}
}

func main() {
	var numbers [5]int
	var sum int

	fmt.Println("Введите 5 чисел: ")

	for i := range 5 {
		fmt.Scanln(&numbers[i])
		sum += numbers[i]
	}

	bubbleSort(numbers[0:5])

	fmt.Printf("Отсортированные элементы: %d\n", numbers)
	fmt.Printf("Самое большое число: %d\n", numbers[0])
	fmt.Printf("Самое маленькое число: %d\n", numbers[4])
	fmt.Printf("Среднее-арифметическое: %d\n", sum/5)
}
