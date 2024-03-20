package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	//Libros
	var myBook = book.Book{
		Title:  "Moby Dick",
		Author: "Herman Van Dicke",
		Pages:  543,
	}
	book.Print(myBook)

	myBook.SetTitle("Moby Dick (Edición Especial)")
	fmt.Println(myBook.GetTitle())

	mySecondBook := book.NewBook("La Ciudad Y Los Perros", "Mario Vargas Llosa", 843)
	book.Print(mySecondBook)

	myTextBook :=book.NewTextBook("Coquito","Enrique Jimenez",159,"Escuela Nueva", "Primaria")
	book.Print(myTextBook)

	//Animales
	miPerro := animal.Perro{Nombre: "Clifford"}
	miGato := animal.Gato{Nombre: "Nino"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Rusty"},
		&animal.Gato{Nombre: "Kuku"},
		&animal.Gato{Nombre: "Lua"},
		&animal.Perro{Nombre: "Kalua"},
	}

	for _,animal :=range animales{
		animal.Sonido()
	}
}
