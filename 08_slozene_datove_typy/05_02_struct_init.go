// Struktury (záznamy) v programovacím jazyku Go
//
// - inicializace prvků datové struktury
// - zápis inicializace odvozený od jazyka C

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	// deklarace proměnné s její inicializací
	user1 := User{
		1,
		"Pepek",
		"Vyskoč"}

	// tisk obsahu celé struktury
	fmt.Println(user1)

	// zápis hodnot prvků do datové struktury
	user1.id = 2
	user1.name = "Josef"
	user1.surname = "Vyskočil"

	// tisk obsahu celé struktury
	fmt.Println(user1)
}
