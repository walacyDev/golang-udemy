package main

import "fmt"

func somar(num *int) {
	*num = *num + 1
}

func subtrair(num *int) {
	*num = *num - 1
}

func main() {
	numero := 10
	fmt.Println("Valor inicial da variavel numero:", numero)
	fmt.Println("Endereço de memoria da variavel numero", &numero)

	somar(&numero)
	somar(&numero)

	fmt.Println("valor atual da variavel numero:", numero)
	subtrair(&numero)

	/* 	var numero int = 60
	   	var p *int = &numero

	   	fmt.Println(numero)
	   	fmt.Println("Endereço na memoria da var numero:", &numero)

	   	fmt.Println("Valor do ponteiro P:", p)
	   	fmt.Println("Valor apontado por p:", *p) */
	// Imprimir o valro no endereço armazenado em P (dereferência)
}
