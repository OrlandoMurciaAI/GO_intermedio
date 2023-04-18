package main

import "fmt"

// los valores no son nulos a la hora de instanciarlos
// GO de por si ya le da valores
type Employee struct {
	id       int
	name     string
	vacation bool
}

// la funcion que recomienda el profe para simular los constructores
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := Employee{}
	fmt.Printf("%v \n", e)

	e2 := Employee{
		id:       1,
		name:     "Orlando",
		vacation: true,
	}

	fmt.Printf("%v \n", e2)

	e3 := new(Employee) // las propiedades quedan por defecto pero me devuelve es un apuntador
	// esto me crea es una referencia de e3
	e3.id = 2
	e3.name = "Daniel"
	fmt.Printf("%v \n", *e3)

	e4 := NewEmployee(1, "Monica", false)
	fmt.Printf("%v \n", *e4)
}
