package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	// Utilizando la palabra reservada select
	// nos permite tener diferentes casos trabajando con diferentes canales
	// Esto evita imprimir de manera secuencial
	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}

	// el programa se qeuda bloqeuadno esperando el resultado del
	// canal 1
	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
