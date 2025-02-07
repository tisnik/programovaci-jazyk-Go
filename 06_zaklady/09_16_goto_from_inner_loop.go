// Základy programovacího jazyka Go
//
// - vnořené programové smyčky
// - výskok z vnořené smyčky pomocí goto

package main

import "fmt"

func main() {
	// vnější počítaná smyčka pro hodnoty i od 0 do 9
	for i := range 10 {
		// vnitřní počítaná smyčka pro hodnoty i od 0 do 9
		for j := range 10 {
			result := (i + 1) * (j + 1)
			fmt.Printf("%3d ", result)
			if result == 42 {
				fmt.Println("\nodpověď nalezena!\n")
				// výskok z obou smyček
				goto Exit
			}
		}
		// odřádkování
		fmt.Println()
	}
	// návěští
Exit:
}
