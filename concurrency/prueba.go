package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan []int, 5)

	// Lanzar una goroutine que envíe valores al canal
	go func() {
		for i := 0; i < 5; i++ {
			x := []int{i, i * 2, i * 3}
			c <- x
			time.Sleep(time.Second) // Esperar un segundo entre cada envío de valor
		}
		close(c) // Cerrar el canal cuando se hayan enviado todos los valores
	}()

	// Recorrer el canal y mostrar los valores
	for v := range c {
		fmt.Println("Valor recibido:", v)
	}

	fmt.Println("Fin del programa")
}
