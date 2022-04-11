package main

import (
	"fmt"
	"time"
)

var (
	intC = make(chan int, 5)
)

func main() {
	go Produce()
	go Consume()
	time.Sleep(time.Second * 5)
	close(intC)
}

func Produce() {
	for i := 0; i < 10; i++ {
		intC <- i
		fmt.Printf("Producing %d\n", i)
	}
}

func Consume() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Consuming %d\n", <-intC)
	}
}
