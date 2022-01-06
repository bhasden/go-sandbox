package main

import (
	"fmt"

	"sandbox/cache"
)

func main() {
	myIntBunch := cache.Bunch[int]{}
	myIntBunch = append(myIntBunch, 1)
	myIntBunch = append(myIntBunch, 3)
	myIntBunch = append(myIntBunch, 2)
	fmt.Println("Printing the max: ")
	fmt.Println(cache.Max(myIntBunch))

	myStringBunch := cache.Bunch[string]{}
	myStringBunch = append(myStringBunch, "a")
	myStringBunch = append(myStringBunch, "b")
	myStringBunch = append(myStringBunch, "c")
	fmt.Println("Printing the max: ")
	fmt.Println(cache.Max(myStringBunch))

	fmt.Printf("Does myIntBunch contain 5? %t\n", myIntBunch.Contains(5))
	fmt.Printf("Does myIntBunch contain 1? %t\n", myIntBunch.Contains(1))

	// this is broken, you can't use a type set as a type
	// myInterfaceBunch := cache.Bunch[constraints.Integer]{}
	// myInterfaceBunch = append(myInterfaceBunch, 1)

	// this is broken b/c error is not constraints.Ordered
	// errorBunch := cache.Bunch[error]{}
	// errorBunch = append(errorBunch, errors.New("new error"))
}
