package _4_tips

var booker *Book = &Book{}

type Book struct {
	Name string
}

func NewBookPointer(name string) *Book {
	return &Book{
		Name: name,
	}
}
