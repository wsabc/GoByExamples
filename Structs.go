package main

import (
	"fmt"
)

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	// must have filed name when less than defined parameters
	fmt.Println(Person{name: "Ann"})
	// print the address?
	fmt.Println(&Person{"Billy", 40})

	s := Person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println("sage: ", sp.sage())
	fmt.Println("sage: ", s.sage2())
}

type Person struct {
	name string
	age  int
}

// receiver is pointer or value, both are ok
// note: functions have no receivers, yet methods have
func (r *Person) sage() int {
	return r.age - len(r.name)
}

func (r Person) sage2() int {
	return r.age - len(r.name)
}
