package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Definimos la estructura Job
type Job struct {
	Name   string        // Nombre del trabajo
	Delay  time.Duration // Tiempo de espera antes de ejecutar el trabajo
	Number int           // Valor que se utilizará para calcular la secuencia de Fibonacci
}

// Definimos la estructura Worker
type Worker struct {
	Id         int           // Identificador del trabajador
	JobQueue   chan Job      // Cola de trabajos que pertenecen al trabajador
	WorkerPool chan chan Job // Grupo de trabajadores que comparten la cola de trabajos
	QuitChan   chan bool     // Canal de comunicación para que el trabajador detenga su trabajo
}

// Definimos la estructura Dispatcher
type Dispatcher struct {
	WorkerPool chan chan Job // Grupo de trabajadores que comparten la cola de trabajos
	MaxWorkers int           // Número máximo de trabajadores que pueden ser activos al mismo tiempo
	JobQueue   chan Job      // Cola de trabajos
}

// Función que crea y devuelve un nuevo trabajador
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

// Función que inicia un trabajador
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				// Se imprime un mensaje que indica que el trabajador está iniciando
				fmt.Printf("Worker with id %d Started\n", w.Id)
				// Se calcula el valor de Fibonacci
				fib := Fibonacci(job.Number)
				// Se espera por el tiempo especificado en Delay
				time.Sleep(job.Delay)
				// Se imprime un mensaje que indica que el trabajador ha terminado y cuál es el resultado
				fmt.Printf("Worker with id %d Finished with result %d\n", w.Id, fib)
			case <-w.QuitChan:
				// Si el trabajador recibe un mensaje a través del canal QuitChan, se imprime un mensaje y el trabajador detiene su trabajo
				fmt.Printf("Worker with id %d Stopped\n", w.Id)
			}
		}
	}()
}

// Función que detiene el trabajo de un trabajador
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Función que calcula el valor de Fibonacci de un número
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" { // GET, PUT, DELETE
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()
	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
