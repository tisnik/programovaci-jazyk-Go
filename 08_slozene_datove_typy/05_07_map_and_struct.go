package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	m1 := make(map[string]User)
	fmt.Println(m1)

	m1["prvni"] = User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	m1["druhy"] = User{
		id:      2,
		name:    "Josef",
		surname: "Vyskočil"}

	fmt.Println(m1)
}
