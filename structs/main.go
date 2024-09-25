package main

import (
	"fmt"
)

type Pessoa struct {
	Nome          string
	idade         int
	Nacionalidade string
}

func main() {

	matheus := Pessoa{Nome: "Matheus Fraga", idade: 30, Nacionalidade: "Brasil"}
	pessoa2 := Pessoa{Nome: "Joaozinho", idade: 28, Nacionalidade: "Brasil"}

	fmt.Println(matheus)
	fmt.Println(pessoa2.idade)

	pessoa2.idade = 29
	pessoa2.Nome = "Joao da Silva"
	fmt.Println("Pessoa 2 dados:", pessoa2)

	pessoas := map[int]Pessoa{
		1: {Nome: "Sujeito", idade: 32, Nacionalidade: "Brasil"},
		2: {Nome: "Fulano", idade: 18, Nacionalidade: "Canadá"},
	}

	fmt.Println(pessoas)
}
