package main

// Los programas quedan bloqueados al no tener un
// canal con un buffer, en estos tenemos que estar 100% seguros
// de que habra alguien que los reciba
import "fmt"

func main() {
	//c := make(chan int) // Canal sin buffer
	c := make(chan int, 3) // Esto es un canal con buffer tienen
	// una cantidad limitada

	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c)

	fmt.Println(<-c)
}
