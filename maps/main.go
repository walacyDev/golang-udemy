package main

import "fmt"

func main() {
	//Declarar um map e inicializar ele
	var capitais map[string]string = make(map[string]string)

	capitais["França"] = "Paris"
	capitais["Itália"] = "Roma"
	capitais["japão"] = "Tokio"

	fmt.Println(capitais)
	fmt.Println("Capital da Itália", capitais["Itália"])

	capitais["Japão"] = "Cidade esquina"

	fmt.Println(capitais)
}
