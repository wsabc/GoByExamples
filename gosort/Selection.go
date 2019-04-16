package main

import "fmt"

func main() {
	n := [5]int{22, 28, 29, 6, 14}
	nlen := len(n)
	var t int
	times := 0
	for i := 0; i < nlen-1; i++ {
		t = i
		for j := i + 1; j < nlen; j++ {
			if n[j] < n[t] {
				times++
				t = j
			}
		}
		if t != i {
			n[i], n[t] = n[t], n[i]
		}
	}
	fmt.Println(n)
	fmt.Println("times", times)
}
