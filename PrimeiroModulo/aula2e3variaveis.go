package main

import "fmt"

//Variáveis do Tipo Var

func main() {
	var nome string = "Lucas Silva"
	var idade = 28
	var texto string

	fmt.Print("Variaveis")
	fmt.Println(nome)
	fmt.Println(idade)

	idade = 29
	texto = "Meu primeiro projeto em GO"

	nome = "Sujeito programador"

	fmt.Println("Agora a idade é", idade)
	fmt.Println(texto)
	println(nome)

	//Variáveis do Tipo Const

	const pi = 3.14159
	const url = "https://sujeitoprogramador.com"

	fmt.Println(pi)
	fmt.Println(url)

}
