package main

import (
	"fmt"
	"sort"
)

func main() {
	// built in sorting
	s := []string{"c", "b", "a"}
	// note: sort in-place, not return a new one
	sort.Strings(s)
	fmt.Println("sorted", s)

	i := []int{3, 7, 2, 1, 8}
	sort.Ints(i)
	fmt.Println("sorted int", i)

	println(sort.IntsAreSorted(i))

	// custom sorting
	fruits := []string{"apple", "kiwi", "banana"}
	// casting
	sort.Sort(byLen(fruits))
	fmt.Println("custom sorted", fruits)
}

type byLen []string

func (s byLen) Len() int {
	return len(s)
}

func (s byLen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
