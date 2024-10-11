package main

import (
	"fmt"
)

func soma(slice []int, channel chan int) {
	sum := 0

	for _, value := range slice {
		sum += value
	}

	//Enviar a soma para o channel
	channel <- sum
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	channel := make(chan int)

	// dividir o slice (array) em duas partes
	mid := len(slice) / 2
	part1 := slice[:mid]
	part2 := slice[mid:]

	go soma(part1, channel)
	go soma(part2, channel)

	//Receber os resultados dos channel
	soma1 := <-channel
	soma2 := <-channel

	totalSoma := soma1 + soma2

	fmt.Printf("Total da soma dos numeros: %d", totalSoma)

}
