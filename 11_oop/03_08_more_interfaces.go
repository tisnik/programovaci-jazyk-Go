// Koncept rozhraní v programovacím jazyku Go
//
// - deklarace dvou rozhraní
// - deklarace uživatelského typu splňujícího obě rozhraní
// - deklarace funkcí akceptujících libovolný typ
//   splňující první resp. druhé rozhraní

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

func f1(i Interface1) {
	fmt.Println("Interface1.f1")
	i.method1()
}

func f2(i Interface2) {
	fmt.Println("Interface2.f2")
	i.method2()
}

func main() {
	t := Type{}

	t.method1()
	t.method2()
	fmt.Println()

	f1(t)
	fmt.Println()

	f2(t)
	fmt.Println()
}
