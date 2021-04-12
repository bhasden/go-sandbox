// Example calendar_test.go
package calendar_test

import (
	"sandbox/calendar"
	"testing"
)

// We do some redundant var handling to ensure no compiler optimizations a la https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
var result int

// func TestCalendar(t *testing.T) {
// 	r := calendar.CountSundays()
// 	if r != 171 {
// 		t.Errorf("Wrong sunday count, expect %d, got %d", 171, r)
// 	}
// }

func BenchmarkCalendar(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = calendar.CountSundays()
		if r != 171 {
			b.FailNow()
		}
	}
	result = r
	if result != 171 {
		b.FailNow()
	}
}
