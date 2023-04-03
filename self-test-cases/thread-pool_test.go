package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/panjf2000/ants"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func TestThreadingPool(t *testing.T) {
	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a method,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}
}

// func print1(i int) {
// 	fmt.Printf("print1: %v.\n", i)
// }

// func print2(i int) {
// 	fmt.Printf("print2: %v.\n", i)
// }

func TestThreadingPoolV2(t *testing.T) {
	var wg sync.WaitGroup
	workSize := 10
	p, _ := ants.NewPool(workSize, ants.WithPreAlloc(true))
	defer p.Release()
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		ants.Submit(func() {
			fmt.Printf("print1: %v.\n", i)
			wg.Done()
		})
	}
	for i := 101; i <= 200; i++ {
		wg.Add(1)
		ants.Submit(func() {
			fmt.Printf("print2: %v.\n", i)
			wg.Done()
		})
	}
	wg.Wait()
}
