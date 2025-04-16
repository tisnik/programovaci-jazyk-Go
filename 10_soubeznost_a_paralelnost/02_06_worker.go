// Gorutiny a kanály v jazyce Go
//
// - funkce s implementací workera
// - asynchronní spuštění tří workerů
// - předání práce workerovi
// - uzavření kanálu indikuje, že se mají workeři ukončit

package main

import "fmt"

// funkce s implementací workera
func worker(id int, task_channel chan int) {
	fmt.Printf("worker %d started\n", id)
	// získání úkolu a test na uzavření kanálu
	for {
		value, more := <-task_channel
		if more {
			fmt.Printf(
				"worker %d received task with parameter %d\n",
				id, value)
		} else {
			fmt.Printf("worker %d finished\n")
			return
		}
	}
}

func main() {
	// kanál pro předávání úkolů
	task_channel := make(chan int)

	fmt.Println("main begin")

	// asynchronní spuštění tří workerů
	go worker(1, task_channel)
	go worker(2, task_channel)
	go worker(3, task_channel)

	// předání deseti úkolů workerům
	for i := 1; i <= 10; i++ {
		fmt.Printf("sending task with parameter %d\n", i)
		task_channel <- i
	}
	// uzavření kanálu
	close(task_channel)

	fmt.Println("waiting for worker...")

	// konec
	fmt.Println("main end")
}
