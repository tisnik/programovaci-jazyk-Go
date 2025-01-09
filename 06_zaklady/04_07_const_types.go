// Základy programovacího jazyka Go
//
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty
// - pokus o přiřazení do proměnných typu MyString

package main

import "fmt"

const (
	zprava1        = "foo"
	zprava2 string = "bar"
)

type MyString string

func main() {
	var v1, v2 MyString
	v1 = zprava1
	v2 = zprava2
	fmt.Println(v1, v2)
}
