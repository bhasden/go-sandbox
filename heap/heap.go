package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxCandies(arr []int, k int) int {
	temp := IntHeap(arr)
	h := &temp
	heap.Init(h)
	candyCount := 0
	for i := 0; i < k; i++ {
		// fmt.Printf("Contents of heap: %v\n", h)
		if biggestBag, ok := heap.Pop(h).(int); ok {
			fmt.Printf("Biggest bag: %d\n", biggestBag)
			candyCount += biggestBag
			heap.Push(h, biggestBag / 2)
		} else {
			panic("not an int!?")
		}
	}
	// Write your code here
	return candyCount;
}

func main() {
	// Call maxCandies() with test cases here
	fmt.Printf("CandyCount: %s\n", maxCandies([]int{2,1,7,4,2}, 3))
	// fmt.Printf("CandyCount: %s\n", maxCandies([]int{4,3,2,1,2,3,4,5,6,7,8}, 10))
}