package main

import "fmt"

func main() {
	// Definir la lista de tareas
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	// Definir el número de trabajadores que se usarán para realizar las tareas
	nWorkers := 3
	// Crear dos canales: uno para enviar trabajos (jobs) y otro para recibir resultados (results)
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	// Iniciar los trabajadores como goroutines
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	// Enviar las tareas a través del canal de trabajos (jobs)
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs) // Cerrar el canal de trabajos cuando se hayan enviado todas las tareas

	// Recibir los resultados a través del canal de resultados (results)
	for r := 0; r < len(tasks); r++ {
		<-results // Esperar a que se reciba un resultado del canal
	}
}

// Definir la función de los trabajadores
func Worker(id int, jobs <-chan int, results chan<- int) {
	// Iterar sobre los trabajos recibidos a través del canal de trabajos (jobs)
	for job := range jobs {
		// Calcular el número de Fibonacci del trabajo actual
		fib := Fibonacci(job)
		// Mostrar información sobre el trabajo y el número de Fibonacci calculado
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		// Enviar el resultado a través del canal de resultados (results)
		results <- fib
	}
}

// Definir la función para calcular el número de Fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
