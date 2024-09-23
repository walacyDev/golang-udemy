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
}
