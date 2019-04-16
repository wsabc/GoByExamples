package main

import "fmt"

func main() {
	n := [5]int{22, 28, 29, 6, 14}
	nlen := len(n)
	times := 0
	for i := 0; i < nlen-1; i++ {
		for j := i + 1; j < nlen; j++ {
			if n[i] > n[j] {
				times++
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	fmt.Println(n)
	fmt.Println("times", times)
}
