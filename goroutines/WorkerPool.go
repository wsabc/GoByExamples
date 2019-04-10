package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w < 3; w++ {
		go worker(w, jobs, results)
	}

	// producer
	for p := 1; p <= 5; p++ {
		jobs <- p
	}
	close(jobs)

	// consumer
	for c := 1; c <= 5; c++ {
		fmt.Println(<-results)
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job * 2
	}
}
