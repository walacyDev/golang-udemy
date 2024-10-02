package main

import "fmt"

type Animal interface {
	EmitirSom()
	Andar()
}

func (c Cachorro) EmitirSom() {
	fmt.Println("Au au AU!")
}

func (c Cachorro) Andar() {
	fmt.Println("O cachorro:", c.Nome)
	fmt.Println("O cachorro está andando...")
}

func (g Gato) EmitirSom() {
	fmt.Println("Miau Miau!!")
}

func (g Gato) Andar() {
	fmt.Println("O gato:", g.Nome)
	fmt.Println("O gato está andando...")
}

type Cachorro struct {
	Nome string
}

type Gato struct {
	Nome string
}

func FazerEmitirSom(a Animal) {
	a.EmitirSom()
}

func FazerAnimalAndar(a Animal) {
	a.Andar()
}

func main() {
	cachorro1 := Cachorro{Nome: "Lola"}

	gato1 := Gato{Nome: "Puma"}

	FazerEmitirSom(cachorro1)
	FazerAnimalAndar(cachorro1)

	FazerEmitirSom(gato1)
	FazerAnimalAndar(gato1)

}
