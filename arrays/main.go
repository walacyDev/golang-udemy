package main

import "fmt"

func main() {

	var numeros [4]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 15
	numeros[3] = 05

	fmt.Println(numeros)

	var valores = [3]int{3, 10, 90}

	fmt.Println(valores)

	permissoes := [3]string{"usuário", "admin", "editor"}

	fmt.Println(permissoes[0])
}
