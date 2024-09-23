package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Conversão de Inteiro paraFloat
	var valor int = 42

	var valorConvertido float64 = float64(valor)

	fmt.Println("Iteiro:", valor)

	fmt.Println("Valor em Float:", valorConvertido)

	// Converter int para string

	var valorString string = strconv.Itoa(valor) // "42"

	fmt.Println(valorString)

	// Converter string para int

	var valor2 string = "230"

	valorInteiro, err := strconv.Atoi(valor2)

	if err != nil {
		fmt.Println("ERRO AO CONVERTER")
	} else {
		fmt.Println(valorInteiro)
	}

	// Converter string para float64
	var pi string = "3.14159"

	piConvertido, error2 := strconv.ParseFloat(pi, 64)

	if error2 != nil {
		fmt.Println("ERRO AO CONVERTER PARA FLOAT")
	} else {
		fmt.Println(piConvertido)
	}
}
