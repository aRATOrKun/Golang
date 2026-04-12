package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func vowelLetters(str string) int {
	cnt := 0
	strRune := []rune(str)
	vowel := []rune{'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я'}

	for j := range utf8.RuneCountInString(str) {
		for i := range 10 {
			if strRune[j] == vowel[i] {
				cnt++
			}
		}
	}

	return cnt
}

func capitalizeWords(s string) string {
	strRune := []rune(s)
	strRune[0] = unicode.ToUpper(strRune[0])

	for i := 1; i < utf8.RuneCountInString(s); i++ {
		if strRune[i] == ' ' {
			i++
			strRune[i] = unicode.ToUpper(strRune[i])
		} else {
			strRune[i] = unicode.ToLower(strRune[i])
		}
	}

	return string(strRune)
}

func bracketCount(str string) {
	if strings.Count(str, "(") == strings.Count(str, ")") {
		fmt.Printf("Скобки расставлены верно, %d открывающиеся, %d закрывающиеся", strings.Count(str, "("), strings.Count(str, ")"))
	} else {
		fmt.Printf("Скобки расставлены неправильно, %d открывающиеся, %d закрывающиеся", strings.Count(str, "("), strings.Count(str, ")"))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите любое предложение или слово на русском:")

	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	fmt.Println("Количество символов в строке:", utf8.RuneCountInString(sentence)) // Задача №1
	fmt.Println("Количество гласных в строке:", vowelLetters(sentence))            // Задача №2
	fmt.Println("Первая буква в слове - заглавная:", capitalizeWords(sentence))    // Задача №3

	reader2 := bufio.NewReader(os.Stdin)

	fmt.Println("Введите формулу со скобками:") // Задача №4

	formula, _ := reader2.ReadString('\n')
	formula = strings.TrimSpace(formula)

	bracketCount(formula)
}
