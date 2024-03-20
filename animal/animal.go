package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau guau")
}

type Gato struct{
	Nombre string
}

func (p *Gato) Sonido() {
	fmt.Println(p.Nombre + " hace miau miau")
}

//funcion para imprimir sonido
func HacerSonido(animal Animal)  {
	animal.Sonido()
}