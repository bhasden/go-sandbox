package cache_test

import (
	"sandbox/cache"
	"sort"

	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCachePutThingsIn(t *testing.T) {
	set := cache.New[string]()
	set.Add("b")
	set.Add("d")

	setTwo := cache.New[string]()
	setTwo.Add("c")
	setTwo.Add("e")

	unionSet := set.Union(*setTwo)
	got := unionSet.Members()
	sort.Slice(got, func(i, j int) bool {
		return got[i] < got[j]
	})
	want := []string{"b", "c", "d", "e"}
	if !cmp.Equal(got, want) {
		t.Fatalf(cmp.Diff(want, got))
	}
}

func TestConcurrency(t *testing.T) {
	s := cache.New[int]()

	go func() {
		for i := 0; i < 1000; i++ {
			s.Members()
		}
	}()

	for i := 0; i < 1000; i++ {
		s.Add(i)
	}

}
