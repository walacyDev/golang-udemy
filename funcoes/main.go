package main

import (
	"errors"
	"fmt"
)

func boasvindas(nome string) {
	fmt.Printf("Bem vindo %s ao sistema \n", nome)
}

func soma(numero1 int, numero2 int) int {
	resultado := numero1 + numero2

	return resultado
	//fmt.Println("A soma dos numeros é:", resultado)
}

type Usuario struct {
	Nome  string
	Senha string
}

func Autentica(user Usuario, senha string) (string, error) {

	if user.Senha != senha {
		return "", errors.New("Senha invalida")
	}

	return user.Nome, nil

	//fmt.Println("Usuário logado:", user)
}

func main() {

	boasvindas("Sujeito programador")

	resultadoSoma := soma(10, 50)

	fmt.Println("Resultado da soma é:", resultadoSoma)

	soma2 := soma(100, 70)
	fmt.Println("Resultado da somae:", soma2)

	usuario1 := Usuario{Nome: "Matheus Fraga", Senha: "123123"}
	usuario2 := Usuario{Nome: "Ana", Senha: "123456"}

	nome, err := Autentica(usuario1, "123123")

	if err != nil {
		fmt.Println("Erro na autenticação:", err)
	} else {
		fmt.Println("Usuário logado:", nome)
	}

	nome, err = Autentica(usuario2, "123123")

	if err != nil {
		fmt.Println("Erro na autenticação:", err)
	} else {
		fmt.Println("Usuário logado:", nome)
	}

}

/* 	boasvindas("Lucas Silva")

boasvindas("Ana Caroline") */
