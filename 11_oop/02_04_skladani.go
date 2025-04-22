// Koncept rozhraní v programovacím jazyku Go
//
// - skládání rozhraní

package main

type Interface1 interface {
	method1()
}

type Interface2 interface {
	// toto rozhraní ve skutečnosti předepisuje
	// metody method1 a method2
	Interface1
	method2()
}
