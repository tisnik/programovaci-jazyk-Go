// Základy programovacího jazyka Go
//
// - manipulace s návratovýi hodnotami přímo v konstrukci defer

package main

import "fmt"

func function1() (i int) {
	// přímé nastavení návratové hodnoty
	i = 1
	// holý příkaz return
	return
}

func function2() (i int) {
	// odložené nastavení návratové hodnoty
	defer func() { i = 2 }()
	// explicitní návrat s hodnotou 1
	return 1
}

func function3() (i int) {
	// odložená modifikace návratové hodnoty
	defer func() { i += 2 }()
	// explicitní návrat s hodnotou 1
	return 1
}

func main() {
	fmt.Printf("Return value of function1: %d\n", function1())
	fmt.Printf("Return value of function2: %d\n", function2())
	fmt.Printf("Return value of function3: %d\n", function3())
}
