// Implementando un waitgroup
// En vez de realizar canales para la sincronizacion
// usaremos un waitgroup

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup //Sincronizar nuestras rutinas

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomehting(i, &wg)
	}

	wg.Wait() // Bloquea el programa hasta que llegue a cero

}

func doSomehting(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	// simulara un proceso muy largo que estemos llevando a cabo
	fmt.Printf("Started %d \n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
