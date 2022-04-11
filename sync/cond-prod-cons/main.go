package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int = 0
	Full  int = 5
	Empty int = 0
	mux   sync.Mutex
	cond  = sync.NewCond(&mux)
)

func main() {
	go produce()
	go produce()
	go consume()
	time.Sleep(time.Second * 10)
}

func produce() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		cond.L.Lock()
		for count == Full {
			cond.Wait()
		}
		count++
		fmt.Println("Producing... count: ", count)
		cond.L.Unlock()
		cond.Broadcast()
	}
}

func consume() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		cond.L.Lock()
		for count == Empty {
			cond.Wait()
		}
		count--
		fmt.Println("Consuming... count: ", count)
		cond.L.Unlock()
		cond.Broadcast()
	}
}
