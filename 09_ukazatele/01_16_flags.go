// Ukazatele v jazyku Go
//
// - praktické použití ukazatelů
// - přečtení argumentů předaných na příkazovém řádku

package main

import (
	"flag"
	"fmt"
)

func main() {
	// proměnné, které budou nastaveny podle parametrů
	// zadaných na příkazovém řádku
	var width int
	var height int
	var aa bool
	var input string
	var output string

	// definice vazby parametrů na proměnné (přes ukazatele)
	flag.IntVar(&width, "width", 0, "image width")
	flag.IntVar(&height, "height", 0, "image height")
	flag.BoolVar(&aa, "aa", false, "enable antialiasing")
	flag.StringVar(&input, "input", "in.png", "input file name")
	flag.StringVar(&output, "output", "out.png", "output file name")

	// načtení parametrů a naplnění proměnných
	flag.Parse()

	// výpis naplněných proměnných
	fmt.Printf("width: %d\n", width)
	fmt.Printf("height: %d\n", height)
	fmt.Printf("antialiasing: %t\n", aa)
	fmt.Printf("input file name: %s\n", input)
	fmt.Printf("output file name: %s\n", output)
}
