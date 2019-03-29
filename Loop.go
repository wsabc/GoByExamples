package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j < 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop infinite")
		break
	}

	for n := 1; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
