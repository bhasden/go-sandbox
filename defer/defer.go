package main

import "fmt"

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func f1() int { // Anonymous return value
	var r int = 6
	defer func() {
		r *= 7
	}()
	return r
}

func f2() (r int) { // Named Return Value
	defer func() {
		r *= 7
	}()
	return 6
}

func f3() (r int) { // Named Return Value
	defer func(r int) {
		r *= 7
	}(r)
	return 6
}
