package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {
	var rOps uint64
	var wOps uint64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case r := <-reads:
				r.resp <- state[r.key]
			case w := <-writes:
				state[w.key] = w.val
				w.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{rand.Intn(5), make(chan int)}
				reads <- read
				// wont block
				<-read.resp
				atomic.AddUint64(&rOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{rand.Intn(5), rand.Intn(100), make(chan bool)}
				writes <- write
				// wont block
				<-write.resp
				atomic.AddUint64(&wOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let goroutines run 1 second
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&rOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&wOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}
