package main

import (
	"fmt"
	"library/book"
)

func main() {
	var myBook = book.Book{
		Title:  "Moby Dick",
		Author: "Herman Van Dicke",
		Pages:  543,
	}
	book.Print(myBook)
	// myBook.PrintInfo()

	myBook.SetTitle("Moby Dick (Edición Especial)")
	fmt.Println(myBook.GetTitle())

	mySecondBook := book.NewBook("La Ciudad Y Los Perros", "Mario Vargas Llosa", 843)
	// mySecondBook.PrintInfo()
	book.Print(mySecondBook)

	myTextBook :=book.NewTextBook("Coquito","Enrique Jimenez",159,"Escuela Nueva", "Primaria")
	// myTextBook.PrintInfo()
	book.Print(myTextBook)
}
