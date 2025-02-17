// Základy programovacího jazyka Go
package main

import "fmt"

func onFinish(message string) {
	fmt.Println(message)
}

func main() {
	// odložené volání funkce onFinish s předáním parametrů
	defer onFinish("Finished")

	// simulace nějaké činnosti
	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
