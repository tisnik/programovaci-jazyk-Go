// Struktury (záznamy) v programovacím jazyku Go
//
// - deklarace uživatelské datové struktury
// - zápis hodnot prvků do datové struktury

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	// proměnná s vynulovanou strukturou
	var user1 User

	// tisk obsahu celé struktury
	fmt.Println(user1)

	// zápis hodnot prvků do datové struktury
	user1.id = 1
	user1.name = "Pepek"
	user1.surname = "Vyskoč"

	// tisk obsahu celé struktury
	fmt.Println(user1)
}
