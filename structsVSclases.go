package main

import "fmt"

type Employee struct {
	// struct es el equivalente a una clase y sus propiedades
	id   int
	name string
}

//Receiver function
// El struct va a poseer un metodo llamado SetId
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetId() int {
	return e.id
}

func main() {
	e := Employee{}
	e.id = 1
	e.name = "Name"
	fmt.Printf("%v", e)
	e.SetId(1075306743)
	e.SetName("orlando")
	fmt.Printf("%v", e)
	fmt.Println(e.GetId())
}
