// Základy programovacího jazyka Go
//
// - nejjednodušší způsob použití konstrukce defer
// - funkce onFinish bude zavolána při opouštění funkce main

package main

import "fmt"

func onFinish() {
	fmt.Println("Finished")
}

func main() {
	// odložené volání funkce onFinish
	defer onFinish()

	// simulace nějaké činnosti
	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}

	// poslední příkaz funkce main
	fmt.Println("Finishing main() function")
}
