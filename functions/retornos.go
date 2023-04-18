package main

// Creacion de funciones variadicas y retornos con nombres
// en una funcion variadica, el parametro de entrada se vuelve
// un slice esto es util cuando no se esta seguro la cantidad
// de valores que va a recibir  nuestra funcion

// Las funciones con retorno nos permite especificar el nombre
// de las variables de salida por lo que en su retorno
// ya go sabe que cuales tiene que mostrar
import "fmt"

func getValues(x int) (double int, triple int, quad int) {
	// Retornos con valor
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}

}

func sum(values ...int) int {
	//values se vuelve un slice de enteros
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	printNames("Orlando", "Laura")
	fmt.Println(getValues(2))
}
