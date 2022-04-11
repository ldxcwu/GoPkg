package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once.")
	}
	//A channel to synchronizing the main thread and goroutines.
	// done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(onceBody)
			// done <- true
			wg.Done()
			fmt.Println("Goroutine ", i)
		}(i)
	}
	// for i := 0; i < 10; i++ {
	// <-done
	// }
	wg.Wait()
	fmt.Println("Main Finished.")

	// Output:
	// Only once.
}
