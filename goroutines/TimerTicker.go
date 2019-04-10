package main

import (
	"fmt"
	"time"
)

func main() {
	// timer and ticker are both use the channel to sync
	timer1 := time.NewTimer(2 * time.Second)
	// block until receive expire notification
	msg := <-timer1.C
	fmt.Println("timer1 expired", msg)

	timer2 := time.NewTimer(time.Second)
	// TODO when the timer would be expired? 1s after now?
	go func() {
		// not get a chance to run?
		<-timer2.C
		fmt.Println("timer2 expired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped")
	}

	// timer focus on future task
	// ticker focus on the interval
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
