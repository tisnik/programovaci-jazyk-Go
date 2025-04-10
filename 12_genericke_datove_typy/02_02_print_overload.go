// Generické datové typy v Go
//
// - Go neumožnuje přetěžování funkcí
// - tento příklad nelze přeložit!

package main

import "fmt"

// přetížená funkce
func printValue(value string) {
	fmt.Println(value)
}

// přetížená funkce
func printValue(value int) {
	fmt.Println(value)
}

func main() {
	printValue("Učíme se jazyk Go!")
}
