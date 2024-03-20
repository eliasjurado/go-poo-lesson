package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s \nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}

func (book *Book) GetTitle() string {
	return book.Title
}

func (book *Book) GetAuthor() string {
	return book.Author
}

func (book *Book) GetPages() int {
	return book.Pages
}

func (book *Book) SetTitle(Title string) *Book {
	book.Title = Title
	return book
}

func (book *Book) SetAuthor(Author string) *Book {
	book.Author = Author
	return book
}

func (book *Book) SetPages(Pages int) *Book {
	book.Pages = Pages
	return book
}
