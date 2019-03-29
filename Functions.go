package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusMore(1, 2, 3)
	fmt.Println("1+2+3=", res)

	a, b := swap(1, 2)
	fmt.Println("first: ", a)
	fmt.Println("second: ", b)

	_, c := swap(1, 2)
	fmt.Println("c: ", c)

	i := sum(1, 2)
	fmt.Println("sum with 1, 2: ", i)

	j := sum(1, 2, 3)
	fmt.Println("sum with 1, 2, 3: ", j)

	s := []int{1, 2, 3, 4}
	// note: note the ... end with the slice name
	d := sum(s...)
	fmt.Println("sum with slice:", d)

	nxtInt := intSeq()
	fmt.Println(nxtInt())
	fmt.Println(nxtInt())
	fmt.Println(nxtInt())

	newInt := intSeq()
	fmt.Println(newInt())

	fmt.Println(fact(7))
}

func plus(a int, b int) int {
	return a + b
}

// omit the type name for the like-typed parameters up to the final
// parameter that declares the type
func plusMore(a, b, c int) int {
	return a + b + c
}

func swap(a, b int) (int, int) {
	return b, a
}

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(i int) int {
	if i == 0 || i == 1 {
		return 1
	} else {
		return i * fact(i-1)
	}
}
