package main

import "fmt"

func main() {
	n := [5]int{22, 28, 29, 6, 14}
	nlen := len(n)
	times := 0
	for i := 0; i < nlen-1; i++ {
		for j := i + 1; j > 0; j-- {
			if n[j] < n[j-1] {
				times++
				n[j], n[j-1] = n[j-1], n[j]
			} else {
				break
			}
		}
	}
	fmt.Println(n)
	fmt.Println("times", times)
}
