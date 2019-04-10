package main

import (
	"fmt"
	"strings"
)

func main() {
	// go does not support generics
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "apple"))
	fmt.Println(Any(strs, func(i string) bool {
		return strings.HasPrefix(i, "p")
	}))

	fmt.Println(All(strs, func(i string) bool {
		return strings.HasPrefix(i, "p")
	}))

	fmt.Println(Filter(strs, func(i string) bool {
		return strings.Contains(i, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}

// below likes Java stream functions
func Index(vs []string, t string) int {
	for i, s := range vs {
		// note: using == to compare string
		if s == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) != -1
}

func Any(vs []string, f func(string) bool) bool {
	for _, s := range vs {
		if f(s) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, s := range vs {
		if !f(s) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	results := make([]string, 0)
	for _, s := range vs {
		if f(s) {
			results = append(results, s)
		}
	}
	return results
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, s := range vs {
		vsm[i] = f(s)
	}
	return vsm
}
