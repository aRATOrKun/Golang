package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var arr [10]int
	slice := make([]int, len(arr))

	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	copy(slice, arr[:])

	sort.Ints(slice)

	fmt.Println(arr)
	fmt.Println(slice)
}
