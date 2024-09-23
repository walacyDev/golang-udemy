package main

import "fmt"

func main() {
	capital := 5000
	taxaJuros := 0.30
	numParcelas := 12

	var jurosFinal float64 = float64(capital) * taxaJuros * float64(numParcelas)

	fmt.Printf("O juros final é: %.2f", jurosFinal)
}
