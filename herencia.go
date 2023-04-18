package main

import "fmt"

// Composicion sobre herencia.
// El polimorfismo no funciona tal cual toca aplicar
// algo llamado las interfaces

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	// Se realiza una estructura por medio de la composicion
	// de las otras
	Person //Campo anonimo
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time Employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary time Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d \n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Orlando"
	ftEmployee.age = 2
	ftEmployee.id = 1075306743
	fmt.Printf("%v \n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
