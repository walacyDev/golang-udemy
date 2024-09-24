package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	numero := 0

	for numero < 5 {
		fmt.Println("PASSANDO NO FOR: ", numero)
		numero++
	}

	for {
		fmt.Println("LOOP INFINITO")
		break
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("VALOR DO I É IGUAL A 3")
		}

		if i == 8 {
			fmt.Println("PARANDO NOSSO FOR")
			break
		}

		fmt.Println("valor do i", i)
	}
}
