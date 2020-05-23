// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Fix the race condition in this program.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// numbers maintains a set of random numbers.
var numbers []int

// rwMutex is used to define a critical section of code.
var rwMutex sync.RWMutex

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Number of goroutines to use.
	const grs = 3

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create three goroutines to generate random numbers.
	for i := 0; i < grs; i++ {
		go func() {
			random(10)
			wg.Done()
		}()
	}

	// Wait for all the goroutines to finish.
	wg.Wait()

	rwMutex.RLock()
	{
		// Display the set of random numbers.
		for i, number := range numbers {
			fmt.Println(i, number)
		}
	}
	rwMutex.RUnlock()
}

// random generates random numbers and stores them into a slice.
func random(amount int) {
	rwMutex.Lock()
	{
		// Generate as many random numbers as specified.
		for i := 0; i < amount; i++ {
			n := rand.Intn(100)
			numbers = append(numbers, n)
		}
	}
	rwMutex.Unlock()
}

// ==================
// WARNING: DATA RACE
// Read at 0x00000125a1d0 by goroutine 8:
//   main.random()
//       /Users/ppannu/code/learn/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0x92
//   main.main.func1()
//       /Users/ppannu/code/learn/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37

// Previous write at 0x00000125a1d0 by goroutine 7:
//   main.random()
//       /Users/ppannu/code/learn/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0xf6
//   main.main.func1()
//       /Users/ppannu/code/learn/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37

// Goroutine 8 (running) created at:
//   main.main()
//       /Users/ppannu/code/learn/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xbc

// Goroutine 7 (finished) created at:
//   main.main()
//       /Users/ppannu/code/learn/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xbc
// ==================
