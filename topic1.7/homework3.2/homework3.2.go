package main

import (
	"fmt"
	"slices"
	"sort"
)

func addCity(city []string) []string {
	var newCity string
	fmt.Scanln(&newCity)

	city = append(city, newCity)
	sort.Strings(city)

	return city
}

func delCity(city []string) []string {
	var delCity string
	fmt.Scanln(&delCity)

	n := sort.SearchStrings(city, delCity)

	return slices.Delete(city, n, n+1)
}

func searchCity(city []string) int {
	var searchCity string
	fmt.Scanln(&searchCity)

	n := sort.SearchStrings(city, searchCity)

	return n + 1
}

func main() {
	cities := []string{"Москва", "Санкт-Петербург", "Краснодар"}

	sort.Strings(cities)

	fmt.Printf("Список текущих городов: %s\nВпишите новое название города, чтобы добавить его в список:\n", cities)
	cities = addCity(cities)
	fmt.Println(cities)

	fmt.Printf("Впишите название города, чтобы удалить его:\n")
	cities = delCity(cities)
	fmt.Println(cities)

	fmt.Printf("Впишите название города, чтобы узнать его номер в списке:\n")
	fmt.Println(searchCity(cities))
}
