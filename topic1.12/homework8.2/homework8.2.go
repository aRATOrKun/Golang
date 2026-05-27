package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func combineFiles(directory []os.DirEntry, file *os.File) {
	for _, i := range directory {
		if filepath.Ext(i.Name()) == ".txt" && i.Name() != "combined.txt" {
			src, _ := os.Open(string(i.Name()))
			io.Copy(file, src)
			file.WriteString("\n")
			src.Close()
		}
	}
}
func main() {
	directory, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.OpenFile("combined.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	combineFiles(directory, file)
}
