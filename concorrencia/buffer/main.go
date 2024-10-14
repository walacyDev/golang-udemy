package main

import (
	"fmt"
	"math/rand"
	"time"
)

func prepereOrder(orderID int, completeOrders chan string) {
	fmt.Printf("Preparando o pedido #%d... \n", orderID)

	//simulação de tempo e preparação (1 a 3 sec)
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	completeOrders <- fmt.Sprintf("Pedido #%d concluido!!", orderID)
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	orderCount := 4
	completeOrders := make(chan string, orderCount)

	for i := 1; i <= orderCount; i++ {
		go prepereOrder(i, completeOrders)
	}

	for i := 1; i <= orderCount; i++ {
		fmt.Println(<-completeOrders)
	}

	//Fechando canal após uso
	close(completeOrders)
}
