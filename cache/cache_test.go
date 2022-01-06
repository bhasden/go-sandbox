package cache_test

import (
	"fmt"
	"sandbox/cache"
	"sort"

	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_cache_putThingsIn(t *testing.T) {
	set := cache.New[string]()
	set.Add("b")
	set.Add("d")

	setTwo := cache.New[string]()
	set.Add("c")
	set.Add("e")

	println("Union Members: ")
	unionSet := set.Union(setTwo)
	unionMembers := unionSet.Members()
	sort.Slice(unionMembers, func(i, j int) bool {
		return unionMembers[i] < unionMembers[j]
	})
	if !cmp.Equal(unionMembers, []string{"b", "c", "d", "e"}) {
		t.Fatalf("didn't get what we wanted")
	}
	fmt.Println(unionMembers)
}
