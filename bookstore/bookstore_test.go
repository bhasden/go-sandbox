package bookstore_test

import (
	"sandbox/bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCatalogGetAllBooks(t *testing.T) {
	want := []bookstore.Book{
		{Title: "Harry Potter and the Philosopher's Scone"},
		{Title: "Chocolat 2: Fifty Shades of Chocolat"},
		{Title: "The Catcher in the Pie"},
	}
	c := bookstore.Catalog(want)
	got := c.GetAllBooks()
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
