package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func errorSearch(line string, cnt *int) {
	re := regexp.MustCompile(`\berror\b|\bERROR\b`)
	if re.MatchString(line) {
		*cnt++
	}
}

func main() {
	file, err := os.Open("server.log")
	if err != nil {
		fmt.Println("Не получилось открыть файл:", err)
		return
	}
	defer file.Close()

	var cnt int

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		errorSearch(line, &cnt)
	}

	fmt.Print(cnt)
}
