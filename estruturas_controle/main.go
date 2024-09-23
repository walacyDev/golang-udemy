package main

import "fmt"

func main() {
	nota1 := 5.0
	nota2 := 8.0

	media := (nota1 + nota2) / 2

	if media >= 7 {
		fmt.Println("Você foi aprovado!")
		fmt.Println("Sua média foi:", media)
	} else if media >= 5 && media <= 6.99 {
		fmt.Println("Você está de exame!")
		fmt.Println("Sua média foi:", media)
	} else {
		fmt.Println("Você foi reprovado!")
		fmt.Println("Sua média foi:", media)
	}

	// ==========================================

	dia := 3

	switch dia {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3:
		fmt.Println("Quarta-feira")
	case 4:
		fmt.Println("Quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	default:
		fmt.Println("Fim de Semana")
	}

}
