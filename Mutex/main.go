package main

// Mutex (Mutual Exclusion)

// So, a mutex, or a mutual exclusion is a mechanism
// that allows us to prevent concurrent processes
// from entering a critical section of data
// whilst it’s already being executed by a given process.

// Race conditions can cause unexpected issues with your systems
// that are both hard to debug and at times, even harder to fix.

// Writing Go programs that can execute concurrently in a safe manner
// without impacting performance. This is where the mutex comes into play.

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 1000
}

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls [WaitGroup.Add] to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls [WaitGroup.Done] when finished.
// At the same time, [WaitGroup.Wait] can be used to block until all goroutines have finished.

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock() // If the lock is already in use, the calling goroutine blocks until the mutex is available.

	fmt.Printf("Depositing %d to your account balance %d \n", value, balance)

	balance += value

	mutex.Unlock() // It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.

	wg.Done() // Done decrements the [WaitGroup] counter by one
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()

	fmt.Printf("Withdrawing %d from your account balance %d \n", value, balance)

	balance -= value

	mutex.Unlock()

	wg.Done()
}

func main() {
	fmt.Println("Bank Account Details:")
	fmt.Println("Bank Name: DontEverTrust Bank")
	fmt.Println("Bank User: Unknown")
	fmt.Println("Balance: ", balance)

	fmt.Printf("\n")
	fmt.Println("Performing Transactions...")

	var wg sync.WaitGroup
	wg.Add(2)

	go deposit(500, &wg)
	go withdraw(700, &wg)

	wg.Wait() // Wait blocks until the [WaitGroup] counter is zero.

	fmt.Println("Account Balance: ", balance)
	fmt.Println("Transaction Successful.")
}
