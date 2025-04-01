// Struktury (záznamy) v programovacím jazyku Go
//
// - inicializace prvků datové struktury
// - vylepšený zápis inicializace s vyjmenováním prvků

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
	// s vyjmenováním nastavovaných prvků
	user1 := User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	// tisk obsahu celé struktury
	fmt.Println(user1)

	// zápis hodnot prvků do datové struktury
	user1.id = 2
	user1.name = "Josef"
	user1.surname = "Vyskočil"

	// tisk obsahu celé struktury
	fmt.Println(user1)
}
