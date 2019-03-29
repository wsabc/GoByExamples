package main

import "fmt"

func zeroval(ival int) {
	// copy the ival, won't change the ival
	ival = 0
}

func zeroptr(iptr *int) {
	// set to the memory address, will change outside variable
	*iptr = 0
}

func main() {
	i := 1

	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// & - get memory address
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// print the address
	fmt.Println("address:", &i)
}
