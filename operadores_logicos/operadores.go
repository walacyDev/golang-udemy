package main

import "fmt"

func main() {
	// and (&&)
	// or (||)
	// negação (!)

	estoque := false
	vendaLiberada := false
	freteGratis := false

	fmt.Println("AND:", estoque && vendaLiberada)
	fmt.Println("OR:", estoque || vendaLiberada)

	fmt.Println("Negação:", !freteGratis)

	fmt.Println(5 != 10)
}
