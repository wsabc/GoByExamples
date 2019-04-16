package main

import "fmt"

func main() {
	a := []int{22, 28, 29, 6, 14}
	n := len(a)
	times := 0
	quick(a, 0, n)
	fmt.Println(a)
	fmt.Println("times", times)
}

func quick(a []int, start, end int) {
	if start >= end {
		return
	}
	for i := start; i < end; i++ {
		pivot := i
		store := pivot + 1
		for j := store; j < end; j++ {
			if a[j] < a[pivot] {
				a[store], a[j] = a[j], a[store]
				store++
			}
		}
		a[pivot], a[store-1] = a[store-1], a[pivot]
		quick(a, start, store-1)
		quick(a, store, end)
	}
}
