package main

import "fmt"

func main() {
	var lista []int //Declarando um slice de inteiro

	fmt.Println("Slice vazio:", lista)

	lista = append(lista, 10, 50, 30)

	fmt.Println("Slice com valores:", lista)
	fmt.Println("Pegando slice na posição 1:", lista[1])

	myArray := [7]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("Meu array:", myArray)

	//Criar um slice a partir de um array existente
	mySlice := myArray[1:5] // Criar um slice incluindo elementos do indice 1 a 4

	fmt.Println("Slice a partir do array:", mySlice)

	mySlice = append(mySlice, 100, 250, 300)

	fmt.Println("mySlice adicionando valores:", mySlice)
	fmt.Println("Pegando slice na posição 3 a 6:", mySlice[0:5])
}
