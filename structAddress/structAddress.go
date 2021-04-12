package main

import (
	"fmt"
)

type theStructLiteral struct {
	theSlice []int
}

// memory of slices
// pointers to slices
// pointers to structs with slices (pointers to slices?)
// structs with slices (pointers to slices?)
func main() {
	// What happens to a struct when it's slice literal is re-allocated?
	// Is it ever different than if a slice pointer is re-allocated?

	// I would think that sometimes the slice would have to be moved to a larger block of contiguous memory.
	// If that is true then the location of the slice moves, does the location of the poin
	firstStruct := theStructLiteral{theSlice: make([]int, 1)}
	secondStruct := firstStruct
	// for i := 0; i < 1000000000; i++ {
	// 	if i % 10000000 == 0 {
	// 		fmt.Printf("Struct Address: %p\n", &firstStruct)
	// 		fmt.Printf("Slice Address:  %p\n", &(firstStruct.theSlice)) // will be different every time
	// 	}
	// 	firstStruct.theSlice = append(firstStruct.theSlice, i)
	// }
	firstStruct.theSlice[0] = 7
	// fmt.Println(firstStruct.theSlice[0])
	// fmt.Println(secondStruct.theSlice[0])
	fmt.Printf("%p\n", firstStruct.theSlice)
	fmt.Printf("%p\n", secondStruct.theSlice)
	fmt.Printf("%p\n", &(firstStruct.theSlice))
	fmt.Printf("%p\n", &(secondStruct.theSlice))
	// isEqual := cmp.Equal(secondStruct.theSlice, firstStruct.theSlice, cmp.Comparer(func(x []int, y []int) bool {
	// 	return &x == &y
	// }))
	// fmt.Printf("%v\n", isEqual)
	// fmt.Printf("%v\n", &secondStruct.theSlice == &firstStruct.theSlice)
	fmt.Println("Hello world")
}
