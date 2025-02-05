// Základy programovacího jazyka Go
//
// - vnořené programové smyčky
// - příkaz break pro výskok z vnořených smyček

package main

import "fmt"

func main() {
	// návěští
Exit:
	// vnější počítaná smyčka pro hodnoty i od 0 do 9
	for i := range 10 {
		// vnitřní počítaná smyčka pro hodnoty i od 0 do 9
		for j := range 10 {
			result := (i + 1) * (j + 1)
			fmt.Printf("%3d ", result)
			if result == 42 {
				fmt.Println("\nodpověď nalezena!\n")
				// ukončení obou smyček
				break Exit
			}
		}
		// odřádkování
		fmt.Println()
	}
}
