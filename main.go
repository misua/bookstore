package main

import (
	"fmt"
)


var books []Book

func AddBook(catalog []Book, book Book) []Book {
	catalog = append(catalog, book)
	return catalog

}
func main() {
	books = []Book{
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

	// fmt.Printf("%+v\n", books[0].Author)

	// modify data

	// books[0] = Book{Title: "Spark Joy"}
	// fmt.Println(books)

	// books[0].Title = "Tony Stark"
	// fmt.Println(books)

	books = AddBook(books, Book{Title: "Sharks", Author: "Tony Sharks", Copies: 3})

	for _, b := range books {
		fmt.Println(b.Author)
	}
}
