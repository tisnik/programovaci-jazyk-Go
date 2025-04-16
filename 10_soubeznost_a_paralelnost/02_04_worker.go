// Gorutiny a kanály v jazyce Go
//
// - funkce s implementací workera
// - asynchronní spuštění workera
// - předání práce workerovi
// - nečeká se na dokončení workera

package main

import "fmt"

// funkce s implementací workera
func worker(task_channel chan int) {
	fmt.Println("worker started")
	// získání úkolu
	for {
		value := <-task_channel
		fmt.Printf("worker received task with parameter %d\n", value)
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

	// konec
	fmt.Println("main end")
}
