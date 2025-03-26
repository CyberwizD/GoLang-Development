package main

import (
	"fmt"
)

func compute(value int) {
	for i := range value {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("GoLang Goroutines Concurrency")

	go compute(10)
	go compute(20)

	go func() {
		fmt.Println("Go Routine")
	}()

	var input string
	fmt.Scanln(&input)
}
