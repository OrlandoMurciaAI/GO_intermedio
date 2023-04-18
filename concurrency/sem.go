package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 2) // El canal estuvo actuando como semaforo
	// ya que la cantidad limita la cantidad de go routines que se ejecutan
	// al mismo tiempo
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomehting(i, &wg, c)
	}
	wg.Wait()
}

func doSomehting(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done() // resta 1 al contador
	fmt.Printf("Id %d started \n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d ended \n", i)
	<-c
}
