package main

/*
Decorators essentially allow you to wrap existing functionality
and append or prepend your own custom functionality on top.
*/

// A Simple Decorator

import (
	"fmt"
	"time"
)

func myfunc() {
	fmt.Printf("Type %T \n", myfunc)
	time.Sleep(1 * time.Second)
}

func decoratorfunc(d func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	d()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func init() {
	fmt.Println("Decorating function...")

	decoratorfunc(myfunc)

	fmt.Println("Decorated function.")
}

func main() {
	fmt.Println("Hello")
}
