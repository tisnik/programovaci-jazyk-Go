// Struktury (záznamy) v programovacím jazyku Go
//
// - funkce, která jako parametr akceptuje záznam
// - funkce, která vytvoří záznam

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

// Funkce, která vytvoří záznam
func NewUser(id uint32, name string, surname string) User {
	user := User{
		id:      id,
		name:    name,
		surname: surname}
	return user
}

// Funkce, která jako parametr akceptuje záznam
func printUser(user User) {
	fmt.Printf("User #%d with name %s and surname %s\n",
		user.id, user.name, user.surname)
}

func main() {
	user1 := NewUser(42, "John", "Doe")

	printUser(user1)
}
