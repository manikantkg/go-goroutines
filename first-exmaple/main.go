package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Hello>>")
	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	wg.Add(len(words))
	for i, x := range words {
		go Words(fmt.Sprintf("%d %s\n", i, x), &wg)

	}
	wg.Wait()
}

func Words(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
