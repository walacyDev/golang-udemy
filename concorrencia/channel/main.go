package main

import "fmt"

func sayHello(chname chan string) {
	chname <- "Bem vindo ao sistema!"
}

func main() {
	ch := make(chan int)
	chname := make(chan string)

	go sayHello(chname)
	go func() {
		ch <- 123
	}()

	fmt.Print("USANDO CHANNELS \n")

	valor := <-ch
	fmt.Print("Valor numero: ", valor)
	fmt.Print("\n")

	msg := <-chname
	fmt.Print(msg)
}
