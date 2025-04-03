// Reflexe v programovacím jazyku Go
//
// - konverze hodnoty předané v parametru typu any
//   na vybraný typ "rozhraní"

package main

import "fmt"

type Interface1 interface {
	foo()
}

type Interface2 interface {
	bar()
}

// deklarace uživatelského datového typu
type Type1 struct {
	name string
}

// implementace metody předepsané v rozhraní Interface1
func (t Type1) foo() {
}

// deklarace uživatelského datového typu
type Type2 struct {
	name string
}

// implementace metody předepsané v rozhraní Interface2
func (t Type2) bar() {
}

// deklarace uživatelského datového typu
type Type3 struct {
	name string
}

// implementace metody předepsané v rozhraní Interface1
func (t Type3) foo() {
}

// implementace metody předepsané v rozhraní Interface2
func (t Type3) bar() {
}

func test_type(value any) {
	switch v := value.(type) {
	case Interface1:
		fmt.Println("Interface1 value:", v)
		fallthrough
	case Interface2:
		fmt.Println("Interface2 value:", v)
		fallthrough
	default:
		fmt.Println("Unsupported value")
	}
}

func main() {
	x := Type1{"x"}
	test_type(x)

	y := Type2{"y"}
	test_type(y)

	z := Type3{"z"}
	test_type(z)
}
