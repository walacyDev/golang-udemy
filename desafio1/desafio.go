package main

import "fmt"

func main() {
	fmt.Println("Desafio 1: Calcular a Idade Atual")

	anoAtual := 2024
	anoNascimento := 2004

	idade := anoAtual - anoNascimento

	fmt.Println("A sua idade é:", idade)

	fmt.Println("Desafio : Converter Horas em Minutos")

	horas := 3

	horasEmMinutos := horas * 60

	fmt.Println("3 horas em minutos é igual:", horasEmMinutos)
	fmt.Printf("%d horas tem %d minutos", horas, horasEmMinutos)
}
