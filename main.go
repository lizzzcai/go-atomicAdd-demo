package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var l sync.Mutex
var wg sync.WaitGroup

// normal add
func add() {
	x++
	wg.Done()
}

// mutex add
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// atomic add
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func demo(f func()) {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
}

func demo1(f func()) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(x)
	fmt.Println(elapsed.Nanoseconds())
}

func main() {
	// f := add // normal add
	// f := mutexAdd // mutex add
	f := atomicAdd // atomic add
	demo1(f)
}
