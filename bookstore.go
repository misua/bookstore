package bookstore

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	Copies     int
	PriceCents int
}

type Catalog map[string]Book

<<<<<<< HEAD
func GetAllBooks(catalog Catalog) GetAllBooks() []Book {
=======
func GetAllBooks(catalog Catalog) []Book {
>>>>>>> d5c5abeb4f5de42492f069111c34550a6e3b139d
	books := []Book{}
	for _, b := range catalog {
		books = append(books, b)
	}
	return books

}
<<<<<<< HEAD
=======

func GetBookDetails(catalog Catalog, ID string) string {
	// for _, b := range catalog {
	// 	if b.ID == ID {
	// 		return fmt.Sprintf("%s, by %s", b.Title, b.Author)
	// 	}
	// }
	b := catalog[ID]
	return fmt.Sprintf("%s, by %s", b.Title, b.Author)
>>>>>>> d5c5abeb4f5de42492f069111c34550a6e3b139d

func GetBookDetails(catalog Catalog, ID string) string {
	for _, b := range catalog {
		if b.ID == ID {
			return fmt.Sprintf("%s, by %s", b.Title, b.Author)
		}
	}
	return ""
}
