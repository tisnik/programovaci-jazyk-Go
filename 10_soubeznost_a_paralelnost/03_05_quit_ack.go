// Gorutiny a kanály v jazyce Go
//
// - konstrukce select-case
// - čtení či zápis do kanálu v konstrukci select-case

package main

import (
	"fmt"
	"time"
)

// funkce s implementací workera
func worker(data chan int, quit, ack chan bool) {
	for {
		select {
		case x := <-data:
			fmt.Println("start working on", x)
			time.Sleep(1 * time.Second)
			fmt.Println("finish working on", x)
		case <-quit:
			fmt.Println("finishing...")
			time.Sleep(4 * time.Second)
			fmt.Println("...done")
			ack <- true
			return
		}
	}
}

func main() {
	// kanál pro posílání dat (úloh)
	data := make(chan int)

	// kanál pro ukončení workera
	quit := make(chan bool)

	// kanál pro potvrzení ukončení činnosti
	ack := make(chan bool)

	// spuštění workera
	go worker(data, quit, ack)

	// poslání úloh workerovi
	data <- 0
	data <- 42
	data <- 100

	// vyžadujeme ukončení činnosti workera
	quit <- true

	// čekáme na potvrzení, že worker ukončil činnost
	<-ack

	fmt.Println("All done")
}
