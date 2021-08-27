package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 1,
	}
}

func TestAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{
			Title:  "Nicholas Chuckleby",
			Author: "Charles Dickens",
			Copies: 1,
		},
		{
			Title:  "Delightfully Uneventful Trip",
			Author: "Agatha Christie",
			Copies: 2,
		},
	}
	want := catalog
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetBpplDetails(t *testing.T){
	bookstore.GetBookDetails()
}