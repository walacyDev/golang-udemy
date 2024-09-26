package main

import (
	"fmt"
)

/* type Pessoa struct {
	Nome          string
	idade         int
	Nacionalidade string
} */

type Usuario struct {
	Nome      string
	Idade     int
	status    bool
	permissao string
}

func main() {

	usuarios := map[int]Usuario{
		1: {Nome: "Sujeito dev", Idade: 28, status: true, permissao: "admin"},
		2: {Nome: "Ana", Idade: 18, status: false, permissao: "user"},
	}

	//atribuindo de forma Hard Code
	usuarios[3] = Usuario{Nome: "Jose", Idade: 90, status: false, permissao: "user"}

	//Atribuindo de forma dinamica
	usuarios[len(usuarios)+1] = Usuario{Nome: "Henrique", Idade: 278, status: true, permissao: "admin"}

	i := 1
	for i <= len(usuarios) {
		fmt.Printf("Nome: %s, Idade %d, permissão atual: %s \n", usuarios[i].Nome, usuarios[i].Idade, usuarios[i].permissao)
		i++
	}

	fmt.Println("===========================================")

	usuarios[len(usuarios)+1] = Usuario{Nome: "Maria", Idade: 22, status: false, permissao: "admin"}

	usuarios[1] = Usuario{Nome: "Sujeito Programador", Idade: 28, status: true, permissao: "admin"}

	fmt.Println(usuarios)

	/* 	matheus := Pessoa{Nome: "Matheus Fraga", idade: 30, Nacionalidade: "Brasil"}
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

	   	fmt.Println(pessoas) */

}
