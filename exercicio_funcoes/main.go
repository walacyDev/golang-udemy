package main

import "fmt"

func novoUsuário(username string) (string, bool) {
	if username != "" {
		return username, true
	}

	return "", false
}

func mensagem() {
	fmt.Println("BEM VINDO AO SISTEMA")
}

func main() {
	usuario1 := "Matheus"

	resultado, status := novoUsuário(usuario1)

	fmt.Printf("Usuário atual: %s, Status atual: %t \n", resultado, status)

	resultado2, status2 := novoUsuário("")
	fmt.Printf("Usuário atual: %s, status atual: %t \n", resultado2, status2)

	func() {
		fmt.Println("FUNÇÃO ANONIMA FOI CHAMADA")
	}()
}
