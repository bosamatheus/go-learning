package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	name  string
	input int
	days  int
	weeks int
}

func (e *Employee) calculateWeeks() {
	e.weeks = e.input / 7
	e.days = e.input % 7
}

type Worker struct {
	id        int
	processed int
}

type Task struct {
	workerID int
	employee Employee
}

func main() {
	//go run main.go 2 João 200 Pedro 100 Maria 7 Otávio 142 José 451 Mariana 5
	args := os.Args[1:]
	poolSize, _ := strconv.Atoi(args[0])
	data := args[1:]

	employees := make(map[string]int)
	for i := 0; i < len(data); i += 2 {
		employees[data[i]], _ = strconv.Atoi(data[i+1])
	}

	jobs := make(chan Task, poolSize)
	results := make(chan Task, poolSize)
	var workers []Worker
	for id := 0; id < poolSize; id++ {
		worker := Worker{id, 0}
		workers = append(workers, worker)
		go run(id, jobs, results)
	}

	for key, value := range employees {
		jobs <- Task{0, Employee{key, value, 0, 0}}
	}
	close(jobs)

	for i := 0; i < len(employees); i++ {
		task := <-results
		e := task.employee
		workers[task.workerID].processed++
		if e.days == 0 {
			fmt.Printf("%s has worked %v weeks in the company\n", e.name, e.weeks)
		} else {
			fmt.Printf("%s has worked %v weeks and %v days in the company\n", e.name, e.weeks, e.days)
		}
	}

	fmt.Printf("\nInfo:\n")
	fmt.Printf("Workers Count: %v\n", len(workers))
	for _, w := range workers {
		fmt.Printf("Worker#%v -> %v elements processed\n", w.id, w.processed)
	}
}

func run(workerID int, jobs <-chan Task, results chan<- Task) {
	for task := range jobs {
		task.workerID = workerID
		task.employee.calculateWeeks()
		results <- task
	}
}
