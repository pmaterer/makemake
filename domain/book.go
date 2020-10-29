package domain

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

type BookRepository interface {
	AddBook(b Book) error
	DeleteBook(id int) error
	GetAllBooks() ([]Book, error)
	GetBook(id int) (Book, error)
	UpdateBook(b Book) error
}
