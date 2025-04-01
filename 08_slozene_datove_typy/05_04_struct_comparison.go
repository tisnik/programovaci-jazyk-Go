// Struktury (záznamy) v programovacím jazyku Go
//
// - inicializace prvků datové struktury
// - vylepšený zápis inicializace s vyjmenováním prvků
// - porovnání dvou struktur

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

	// deklarace druhé proměnné typu záznam
	var user2 User
	user2.id = 1
	user2.name = "Pepek"
	user2.surname = "Vyskoč"

	// deklarace třetí proměnné typu záznam
	user3 := User{
		id:      1,
		name:    "Josef",
		surname: "Vyskočil"}

	// záznamy (struktury) je možné porovnat operátorem ==
	fmt.Println(user1 == user2)
	fmt.Println(user1 == user3)
}
