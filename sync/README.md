# Sync Reading Notes
> Package `sync` provides basic synchronization primitives such as mutual exclusion locks.
Other than the `Once` and `WaitGroup` types, 
most are intended for use by low-level library routines.
Higher-level synchronization is better done via `channels` and communication.
> 
- Translation
    
    `sync` 包提供了基础的同步原语，例如互斥锁，
    
    除了 `Once` 和 `WaitGroup` 类型，
    
    大多数都是用来提供低级库历程服务，
    
    更高级的同步最好使用 `channel`和通信来实现。
    

---

**Index：**

---

## type `Locker`

A `Locker` represents an object that can be locked and unlocked.

```go
type Locker interface {
	Lock()
	Unlock()
}
```

## type [`Cond`](https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/sync/cond.go;l=21)

`Cond` represents a conditional variable which has an associate `Locker`.

```go
type Cond struct {

	// L is held while observing or changing the condition
	L Locker
	// contains filtered or unexported fields
}
```

### func `NewCond`

```go
func NewCond(l Locker) *Cond
```

### func `(*Cond) Broadcast`

```go
//Broadcast wakes all goroutines waiting on c.
func (c *Cond) Broadcast()
```

### func `(*Cond) Signal`

```go
//Signal wakes one goroutine waiting on c, if there is any.
func (c *Cond) Signal()
```

### func `(*Cond) Wait`

```go
//Wait() wait locks c.L before returning.
//So the caller should Wait in a loop.
c.L.Lock()
for !condition() {
	c.Wait() //Wait() return with c.L.Locked.
}
...
c.L.Unlock()
```

### Make pro-con

- low - level sync with `Cond`
    
    ```go
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
    ```
    
- Higher level synchronization is better via `channels` and communication.
    
    ```go
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
    ```
    

---

## type `Map`

Like a Go `map[interface{}]interface{}` but is safe for concurrent use.

## type `Mutex`

`Mutex` is an object which has realized `Locker` interface.

### func of `Mutex`

```go
func (m *Mutex) Lock()
func (m *Mutex) Unlock()
func (m *Mutex) TryLock() bool 
```

## type `Once`

`Once` is an object that will perform exactly one action.

### func `(*Once) Do`

```go
func (o *Once) Do(f func())
```

`Do` calls the function f if and only if `Do` is being called for the first time for this instance of `Once`.

`Do` is intended for initialization that must be run exactly once. such as:

```go
config.once.Do(func() {config.init(filename)} )
```

- example
    
    ```go
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
    	done := make(chan bool)
    	for i := 0; i < 10; i++ {
    		go func() {
    			once.Do(onceBody)
    			done <- true
    		}()
    	}
    	for i := 0; i < 10; i++ {
    		<-done
    	}
    
    	// Output:
    	// Only once.
    }
    ```
    

---

## type `WaitGroup`

A `WaitGroup` waits for a collection of goroutines to finish.
The main goroutine calls `Add` to set the number of goroutines to wait for.
The each of the goroutines runs and calls `Done` when finished.
At the same time, `Wait` can be used to block until all goroutines have finished.

### func of `WaitGroup`

```go
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
```

- example
    
    ```go
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
    ```