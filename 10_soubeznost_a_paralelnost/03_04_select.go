// Gorutiny a kanály v jazyce Go
//
// - konstrukce select-case
// - explicitní ukončení workera
// - čekání na potvrzení workera

package main

import "fmt"

// funkce s implementací workera
func worker(data chan int, quit chan bool) {
	for {
		select {
		case x := <-data:
			fmt.Println("Received data", x)
		case <-quit:
			// potvrzujeme ukončení činnosti
			quit <- true
			return
		}
	}
}

func main() {
	// kanál pro posílání dat (úloh)
	data := make(chan int)

	// kanál pro ukončení workera a potvrzení ukončení činnosti
	quit := make(chan bool)

	// spuštění workera
	go worker(data, quit)

	// poslání úloh workerovi
	data <- 0
	data <- 42
	data <- -100

	// vyžadujeme ukončení činnosti workera
	quit <- true

	// čekáme na potvrzení, že worker ukončil činnost
	<-quit

	fmt.Println("All done")
}
