// Generické datové typy v Go
//
// - různé varianty funkce pro porovnání
//   dvou hodnot
// - každá varianta této funkce má odlišné jméno

package main

import "fmt"

func compareInts(x int, y int) bool {
	return x < y
}

func compareFloats(x float64, y float64) bool {
	return x < y
}

func compareStrings(x string, y string) bool {
	return x < y
}

func main() {
	fmt.Println(compareInts(1, 2))
	fmt.Println(compareFloats(1.5, 2.6))
	fmt.Println(compareStrings("foo", "bar"))
}
