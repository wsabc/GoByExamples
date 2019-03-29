package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a [5]int
	fmt.Println("empty: ", a)

	a[4] = 4
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("length: ", len(a))

	var d = [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", d)

	var twod [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twod)

	// note: slices much more often than arrays
	var s = make([]string, 5)
	fmt.Println("emp slice: ", s)

	s[0] = "s1"
	s[1] = "s2"
	s[2] = "s3"
	fmt.Println("set slice: ", s)
	fmt.Println("get slice: ", s[1])
	fmt.Println("len slice: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append slice: ", s)

	ss := make([]string, len(s))
	copy(ss, s)
	fmt.Println("copy slice: ", ss)

	l := ss[2:5]
	fmt.Println("slice 2:5 ", l)

	l2 := ss[2:]
	fmt.Println("slice 2: ", l2)

	l3 := ss[:5]
	fmt.Println("slice :5 ", l3)

	// note the difference between initialize an array
	t := []string{"A", "B"}
	fmt.Println("dcl: ", t)

	var D2 = make([][]string, 3)
	for n := 0; n < 3; n++ {
		inL := n + 1
		D2[n] = make([]string, inL)
		for m := 0; m < inL; m++ {
			// note: how to cast a int to string
			D2[n][m] = strconv.Itoa(n * m)
		}
	}
	fmt.Println("2d:", D2)

	// maps

	mp := make(map[string]int)
	fmt.Println("emp: ", mp)

	mp["good"] = 1
	mp["bad"] = 2
	fmt.Println("set: ", mp)

	fmt.Println("get: ", mp["good"])
	fmt.Println("len: ", len(mp))

	delete(mp, "bad")
	fmt.Println("deleted: ", mp)

	// not present may also return 0 or ""
	// pres can indicate if the key present
	_, pres := mp["good"]
	fmt.Println("prs: ", pres)

	np := map[string]int{"g": 1, "b": 2}
	fmt.Println("dcl:", np)

}
