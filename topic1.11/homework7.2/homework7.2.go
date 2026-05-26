package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(intCh chan int) {

	for range 5 {
		intCh <- rand.Intn(10) + 1
	}
	close(intCh)
}

func reciever(intCh chan int) {

	for {
		value, ok := <-intCh
		if !ok {
			break
		}
		fmt.Println("Получено число:", value)
	}
	fmt.Println("Канал закрыт")
}

func main() {
	intCh := make(chan int)
	go generator(intCh)
	go reciever(intCh)
	time.Sleep(1 * time.Second)
}
