package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.NewTicker(200 * time.Millisecond)
	for req := range requests {
		<-limiter.C
		fmt.Println("process request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i < 4; i++ {
		// for bursty requests
		burstyLimiter <- time.Now()
	}
	go func() {
		// for rate limit requests
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for br := range burstyRequests {
		<-burstyLimiter
		fmt.Println("pr:", br, time.Now())
	}
}
