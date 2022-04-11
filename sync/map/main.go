package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("ping", "pong")
	if v, ok := m.Load("ping"); ok {
		fmt.Println(v)
	}
	actual, loaded := m.LoadOrStore("ping", "ppong")
	fmt.Printf("After LoadOrStore, actual: %v, loaded: %v\n", actual, loaded)

	v, loaded := m.LoadAndDelete("ping")
	fmt.Printf("After LoadAndDelete, v: %v, loaded: %v\n", v, loaded)

	m.Store(1, 1)
	m.Store(2, 2)
	m.Store(3, 3)

	m.Range(func(key, value interface{}) bool {
		fmt.Println("key: ", key, " value: ", value)
		return true
	})

	// pong
	// After LoadOrStore, actual: pong, loaded: true
	// After LoadAndDelete, v: pong, loaded: true
	// key:  1  value:  1
	// key:  2  value:  2
	// key:  3  value:  3
}
