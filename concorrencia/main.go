package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Ola mundo!!!")
}

func main() {
	go sayHello()
	fmt.Println("APRENDENDO CONCORRENCIA")

	time.Sleep(1 * time.Second)
}
