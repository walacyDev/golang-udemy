package main

import "fmt"

func main() {
	var tarefas []string

	//Adicionar itens no slice

	tarefas = append(tarefas, "Estudar GO")
	tarefas = append(tarefas, "Seguir o canal do sujeito")
	tarefas = append(tarefas, "Ler um livro")
	tarefas = append(tarefas, "Comprar coca")

	fmt.Println("Tarefas:", tarefas)
	fmt.Println("Tamanho do slice atual:", len(tarefas))

	//Slicing / Cortar

	tarefas = tarefas[1:]

	fmt.Println("Remover a primeira tarefa:", tarefas)

	tarefas = tarefas[:len(tarefas)-1]
	fmt.Println("Removendo ultimo item:", tarefas)

	tarefas = append(tarefas[:1], tarefas[2:]...)
	fmt.Println("Removendo o item 2 da slice")
}
