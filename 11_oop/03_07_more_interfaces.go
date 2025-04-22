// Koncept rozhraní v programovacím jazyku Go
//
// - deklarace dvou rozhraní
// - deklarace uživatelského typu splňujícího obě rozhraní

package main

import "fmt"

type Interface1 interface {
	method1()
}

type Interface2 interface {
	method2()
}

type Type struct{}

func (Type) method1() {
	fmt.Println("Type.method1")
}

func (Type) method2() {
	fmt.Println("Type.method2")
}

func main() {
	t := Type{}

	t.method1()
	t.method2()
	fmt.Println()
}
