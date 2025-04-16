// Gorutiny a kanály v jazyce Go
//
// - funkce s implementací workera
// - asynchronní spuštění workera
// - předání práce workerovi
// - uzavření kanálu indikuje, že se má worker ukončit

package main

import "fmt"

// funkce s implementací workera
func worker(task_channel chan int) {
	fmt.Println("worker started")
	// získání úkolu a test na uzavření kanálu
	for {
		value, more := <-task_channel
		if more {
			fmt.Printf(
				"worker received task with parameter %d\n",
				value)
		} else {
			fmt.Println("worker finished")
			return
		}
	}
}

func main() {
	// kanál pro předávání úkolů
	task_channel := make(chan int)

	fmt.Println("main begin")

	// asynchronní spuštění workera
	go worker(task_channel)

	// předání deseti úkolů workerovi
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
