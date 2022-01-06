package cache

import (
	"constraints"
)

type Bunch[E constraints.Ordered] []E

// method
func (b Bunch[E]) Contains(anE E) bool {
	for _, value := range b {
		if value == anE {
			return true
		}
	}
	return false
}

// function
func Max[E constraints.Ordered](myBunch Bunch[E]) E {
	var max E
	for _, value := range myBunch {
		if value > max {
			max = value
		}
	}
	return max
}
