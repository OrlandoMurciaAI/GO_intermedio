// Repasando go

// map son estructuras llave valor
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Declaracion de variables
	var x int
	x = 8  // Version explicita
	y := 7 //Version simplificada
	fmt.Println(x)
	fmt.Println(y)

	//Manejo de errores
	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v \n", err)
	} else {
		fmt.Println(myValue)
	}

	// MAPS
	m := make(map[string]int) // llave de tipo string con valores enteros
	m["Key"] = 6
	fmt.Println(m["Key"])

	// Slice
	s := []int{1, 2, 3} // es un vector
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	// Adicionando valores a un vector
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g)
	// El apuntador de g es h es una referencia
	//para acceder a su valor se utiliza el asterisco
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

// Simulacion de una Go Routine
// Una go routine se crea pero necesita de los canales para funcionar
// La go routine va a ejecutarse hasta que envie la informacion al canal
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
