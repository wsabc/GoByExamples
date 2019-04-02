package main

import (
	"fmt"
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

}
