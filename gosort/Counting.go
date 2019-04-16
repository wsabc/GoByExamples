package main

import "fmt"

func main() {
	n := [5]int{22, 28, 29, 6, 14}
	nlen := len(n)
	times := 0
	max := 29
	nn := make([]int, max+1)
	for i := 0; i < max+1; i++ {
		nn[i] = 0
	}
	for i := 0; i < nlen; i++ {
		t := n[i]
		nn[t]++
	}
	nnn := make([]int, nlen)
	in := 0
	for i := 0; i <= max; i++ {
		count := nn[i]
		if count != 0 {
			for j := 0; j < count; j++ {
				nnn[in] = i
				in++
			}
		}
	}
	fmt.Println(nnn)
	fmt.Println("times", times)
}
