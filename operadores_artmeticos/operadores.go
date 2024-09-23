package main

import "fmt"

func main() {
	valor1 := 10
	valor2 := 7

	resultado := valor1 + valor2

	fmt.Println("Resultado da Soma:", resultado)

	// Subtração ( - )

	subtracao := valor1 - valor2
	fmt.Println("Valor da subtração:", subtracao)

	// Multiplicação ( * )

	resultadoMultiplicacao := valor1 * valor2

	fmt.Println("Valor da multiplicação:", resultadoMultiplicacao)

	// Divisão ( / )

	nota1 := 7.00
	nota2 := 6.50

	resultadoNota := (nota1 + nota2) / 2

	fmt.Println("Resultado nota:", resultadoNota)

	x := 10
	y := 3

	resto := x % y

	fmt.Println("Resto da divisão:", resto)
}
