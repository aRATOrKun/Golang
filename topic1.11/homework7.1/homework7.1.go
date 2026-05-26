package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {

	fmt.Printf("Горутина %d начала работу \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина %d завершила работу \n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go work(1, &wg)
	go work(2, &wg)
	go work(3, &wg)

	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}
