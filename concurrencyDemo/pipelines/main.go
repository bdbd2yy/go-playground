package main

import (
	"fmt"
	"sync"
)

func main() {
	// fan-out, fan-in
	// fan-out: 1->N, distribute the work amongst a group of workers
	// fan-in: N->1
	in := gen(2, 3)

    done := make(chan struct{})
	c1 := sq(done, in)
	c2 := sq(done, in)

    // once done is closed, the receive opration on this channel will immediately receive the zero-value of the channel
    defer close(done)
	for n := range merge(done, c1, c2) {
		fmt.Println(n)
	}

    // done <- struct{}{}
    // done <- struct{}{}
}

// you can directly assign a bidirectional channel to a receive-only or send-only channel
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
        defer close(out)
        for n := range in {
            select {
            case out <- n * n:
            case <- done:
                return
            }
        }
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int, 1)

	var wg sync.WaitGroup
	output := func(c <-chan int) {
        // break will only exit the select rather than the for loop, so we need a label
        for num := range c {
            select {
            case out <- num:
            case <-done:
                return
            }
        }
        wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
