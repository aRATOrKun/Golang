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
	cnt := 0

	for j := range utf8.RuneCountInString(s) {
		cnt += 1
		if strRune[j] == ' ' || j == utf8.RuneCountInString(s)-1 {
			for i := j + 1 - cnt; i < utf8.RuneCountInString(s); i++ {
				if i == j+1-cnt {
					strRune[i] = unicode.ToUpper(strRune[i])
				} else {
					strRune[i] = unicode.ToLower(strRune[i])
				}
			}

			cnt = 0
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

	fmt.Println("Количество символов в строке:", utf8.RuneCountInString(sentence))
	fmt.Println("Количество гласных в строке:", vowelLetters(sentence))
	fmt.Println("Первая буква в слове - заглавная:", capitalizeWords(sentence))

	reader2 := bufio.NewReader(os.Stdin)

	fmt.Println("Введите формулу со скобками:")

	formula, _ := reader2.ReadString('\n')
	formula = strings.TrimSpace(formula)

	bracketCount(formula)
}
