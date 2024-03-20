package book

import (
	"fmt"
)

type TextBook struct {
	Book
	Editorial string
	Level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		Editorial: editorial,
		Level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s \nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n", b.Title, b.Author, b.Pages, b.Editorial, b.Level)
}
func (textbook *TextBook) GetEditorial() string {
	return textbook.Editorial
}

func (textbook *TextBook) GetLevel() string {
	return textbook.Level
}

func (textbook *TextBook) SetEditorial(Editorial string) *TextBook {
	textbook.Editorial = Editorial
	return textbook
}

func (textbook *TextBook) SetLevel(Level string) *TextBook {
	textbook.Level = Level
	return textbook
}
