package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) GetInfo() {
	fmt.Printf("\"%s\" by %s, %d\n", b.Title, b.Author, b.Year)
}

func main() {
	var book = Book{"Преступление и наказание", "Достоевский Ф.М.", 1866}

	book.GetInfo()
}
