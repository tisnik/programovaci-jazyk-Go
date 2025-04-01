// Struktury (záznamy) v programovacím jazyku Go
//
// - pole, jehož prvky jsou typu struktura (záznam)

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	// pole záznamů
	var users [3]User

	// tisk obsahu celého pole
	fmt.Println(users)

	// první prvek, který zapíšeme do pole
	user1 := User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	// druhý prvek, který zapíšeme do pole
	var user2 User
	user2.id = 1
	user2.name = "Pepek"
	user2.surname = "Vyskoč"

	// třetí prvek, který zapíšeme do pole
	user3 := User{
		id:      1,
		name:    "Josef",
		surname: "Vyskočil"}

	// změna obsahu prvků pole
	users[0] = user1
	users[1] = user2
	users[2] = user3

	// tisk obsahu celého pole
	fmt.Println(users)
}
