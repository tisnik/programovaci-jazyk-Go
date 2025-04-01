// Struktury (záznamy) v programovacím jazyku Go
//
// - funkce, která jako parametr akceptuje záznam

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

// Funkce, která jako parametr akceptuje záznam
func printUser(user User) {
	fmt.Printf("User #%d with name %s and surname %s\n",
		user.id, user.name, user.surname)
}

func main() {
	// deklarace proměnné s její inicializací
	// s vyjmenováním nastavovaných prvků
	user1 := User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	printUser(user1)
}
