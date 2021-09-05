package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

<<<<<<< HEAD
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
=======
// func TestBook(t *testing.T) {
// 	t.Parallel()
// 	_ = bookstore.Book{
// 		ID:"1".
// 		Title:  "Nicholas Chuckleby",
// 		Author: "Charles Dickens",
// 		Copies: 1,
// 	}
// }

var catalog = bookstore.Catalog{
	"1": {
>>>>>>> d5c5abeb4f5de42492f069111c34550a6e3b139d
		ID:     "1",
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 1,
	},

	"2": {
		ID:         "2",
		Title:      "Delightfully Uneventful Trip",
		Author:     "Agatha Christie",
		Copies:     2,
		PriceCents: 0,
	},
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
<<<<<<< HEAD

	want := []bookstore.Book{

		{
			ID:     "1",
			Title:  "Nicholas Chuckleby",
			Author: "Charles Dickens",
			Copies: 1,
=======
	want := []bookstore.Book{
		{
			ID:         "1",
			Title:      "Nicholas Chuckleby",
			Author:     "Charles Dickens",
			Copies:     1,
			PriceCents: 0,
>>>>>>> d5c5abeb4f5de42492f069111c34550a6e3b139d
		},
		{
			ID:     "2",
			Title:  "Delightfully Uneventful Trip",
			Author: "Agatha Christie",
			Copies: 2,
		},
	}
<<<<<<< HEAD
	got := catalog.GetAllBooks(catalog)
=======
	got := bookstore.GetAllBooks(catalog)
>>>>>>> d5c5abeb4f5de42492f069111c34550a6e3b139d
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
<<<<<<< HEAD

=======
>>>>>>> d5c5abeb4f5de42492f069111c34550a6e3b139d
}
