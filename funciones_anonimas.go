package main

import (
	"fmt"
	"time"
)

// Funcion anonima es aquella que solo usaremos una vez
//pueden recibir parametros como una funcion normal

func main() {
	// funcion anonima que recibe un dato
	z := func(n int) int {
		return n * 2
	}(5)
	fmt.Printf("%d \n", z)

	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	// Funcion anonima concurrente (GO ROUTINE)
	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c

}
