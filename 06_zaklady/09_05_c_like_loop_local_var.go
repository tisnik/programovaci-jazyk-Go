// Základy programovacího jazyka Go
//
// - programová smyčka odvozená od jazyka C
// - počitadlem je proměnná viditelná jen v bloku smyčky
// - POZOR: tento příklad nelze přeložit

package main

import "fmt"

func main() {
	// proměnná i je viditelná jen v bloku smyčky
	for i := 0; i < 10; i++ {
		fmt.Println("Counting:", i)
	}
	fmt.Println()

	// pokus o přístup k počitadlu smyčky
	fmt.Println("i =", i)
}
