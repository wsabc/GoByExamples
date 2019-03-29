package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 4}
	sum := 0
	for _, num := range s {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	for i, num := range s {
		if i == 2 {
			fmt.Println("#2 is: ", num)
		}
	}

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// only the key, not grantee the sequence
	for k := range m {
		fmt.Println("key: ", k)
	}

	// for string, special
	for start, text := range "go" {
		fmt.Println("start byte index", start, text)
	}

}
