package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = []bookstore.Book{
	{
		ID:     "1",
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 1,
	},
	{
		ID:     "2",
		Title:  "Delightfully Uneventful Trip",
		Author: "Agatha Christie",
		Copies: 2,
	},
}

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		ID:     "1",
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 1,
	}
}

func TestAllBooks(t *testing.T) {
	t.Parallel()

	want := []bookstore.Book{

		{
			ID:     "1",
			Title:  "Nicholas Chuckleby",
			Author: "Charles Dickens",
			Copies: 1,
		},
		{
			ID:     "2",
			Title:  "Delightfully Uneventful Trip",
			Author: "Agatha Christie",
			Copies: 2,
		},
	}
	got := catalog.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetBookDetails(t *testing.T) {
	want := "Nicholas Chuckleby, by Charles Dickens"
	got := bookstore.GetBookDetails(catalog, "1")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
