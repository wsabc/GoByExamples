package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(s, " : ", i)
	}
}

func main() {
	f("direct")

	// lightweight thread
	go f("goroutine")

	go func(s string) {
		fmt.Printf(s)
	}("anonymous goroutine")

	fmt.Scanln()
	fmt.Println("done")

	// channels are goroutine pipes
	q := make(chan string)

	// block before both sender and receiver are ready
	// by default it's unbuffered
	go func() { q <- "ping" }()

	msg := <-q
	fmt.Println(msg)

	// make buffered channels
	bq := make(chan string, 2)
	bq <- "one"
	bq <- "two"

	fmt.Println(<-bq)
	fmt.Println(<-bq)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed string")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// combine go routines and channels with select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("loop: ", i)
		select {
		case msg := <-c1:
			fmt.Println("received:", msg)
		case msg := <-c2:
			fmt.Println("received:", msg)
		}
	}

	// timeout
	t1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		t1 <- "result 1"
	}()

	select {
	case msg := <-t1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	// nio
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received messages:", msg)
	default:
		fmt.Println("no messages received")
	}

	msghi := "hi"
	select {
	case messages <- msghi:
		fmt.Println("messages sent", msghi)
	default:
		// no receiver
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received messages:", msg)
	case sig := <-signals:
		fmt.Println("received signals:", sig)
	default:
		fmt.Println("no activity")
	}

	jobs := make(chan int, 5)
	jobdone := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job ", j)
			} else {
				fmt.Println("all done")
				jobdone <- true
				return
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		fmt.Println("send job ", j)
		jobs <- j
	}

	close(jobs)
	fmt.Println("sent all jobs")
	// wait the goroutine
	<-jobdone

	qq := make(chan string, 2)
	qq <- "one"
	qq <- "two"
	close(qq)
	// still can iterate channel even after it's closed
	for e := range qq {
		fmt.Println("range chan:", e)
	}
}

// channel as function parameters with directions
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
