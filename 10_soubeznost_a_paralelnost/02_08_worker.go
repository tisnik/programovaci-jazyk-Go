// Gorutiny a kanály v jazyce Go
//
// - funkce s implementací workera
// - asynchronní spuštění tří workerů
// - předání práce workerovi
// - uzavření kanálu indikuje, že se mají workeři ukončit
// - aktivní čekání na dokončení workerů

package main

import (
	"fmt"
	"time"
)

// funkce s implementací workera
func worker(id int, task_channel chan int, worker_done chan bool) {
	fmt.Printf("worker %d started\n", id)
	// získání úkolu a test na uzavření kanálu
	for {
		value, more := <-task_channel
		if more {
			fmt.Printf(
				"worker %d received task with parameter %d\n",
				id, value)
			time.Sleep(2 * time.Second)
			fmt.Printf(
				"worker %d finished task with parameter %d\n",
				id, value)
		} else {
			fmt.Printf("finishing worker %d\n", id)
			// potvrzení ukončení činnosti workera
			worker_done <- true
			fmt.Printf("worker %d finished\n", id)
			return
		}
	}
}

const WORKERS = 3

func main() {
	// kanál pro předávání úkolů
	task_channel := make(chan int)

	// kanál pro potvrzení ukončení činnosti workera
	worker_done := make(chan bool)

	fmt.Println("main begin")

	// asynchronní spuštění tří workerů
	for i := range WORKERS {
		go worker(i+1, task_channel, worker_done)
	}

	// předání deseti úkolů workerům
	for i := 1; i <= 10; i++ {
		fmt.Printf("sending task with parameter %d\n", i)
		task_channel <- i
	}
	// uzavření kanálu
	close(task_channel)

	fmt.Println("waiting for workers...")

	// čekání na všechny tři workery
	for range WORKERS {
		code, status := <-worker_done
		fmt.Printf("received code: %t and status: %t\n", code, status)
	}

	// konec
	fmt.Println("main end")
}
