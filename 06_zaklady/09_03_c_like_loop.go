// Základy programovacího jazyka Go
//
// - programová smyčka odvozená od jazyka C
// - jako počitadlo je použita běžná lokální proměnná

package main

import "fmt"

func main() {
	// proměnná, která bude použita jako počitadlo smyčky
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println("Counting:", i)
	}
	fmt.Println()

	// proměnná i je viditelná i po dokončení smyčky
	fmt.Println("i =", i)
}
