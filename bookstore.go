package bookstore

type Book struct {
	Title      string
	Author     string
	Copies     int
	PriceCents int
}

func GetAllBooks(catalog []Book) []Book {
	return catalog

}
