package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	for _, i := range []int{2, 7} {
		r, e := f1(i)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r, "worked.")
		}
		// r, e are explain each other?
	}

	for _, i := range []int{3, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(3)
	// ae is the casted object, ok means if cast worked
	// note the format e.(*anError)
	ae, ok := e.(*anError)
	if ok {
		println(ae.message)
	}
}

func f1(arg int) (int, error) {
	if arg == 2 {
		return -1, errors.New("invalid arguments")
	}
	return -1 * arg, nil
}

type anError struct {
	code    int
	message string
}

// what is the interface of error?
func (r anError) Error() string {
	return strconv.Itoa(r.code) + ": " + r.message
}

func f2(arg int) (int, error) {
	if arg == 3 {
		return -1, &anError{3, "not 3"}
	}
	return arg * 3, nil
}
