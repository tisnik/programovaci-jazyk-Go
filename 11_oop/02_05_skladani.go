// Koncept rozhraní v programovacím jazyku Go
//
// - skládání rozhraní
// - hierarchie typu "diamant"

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

type Interface3 interface {
	// toto rozhraní ve skutečnosti předepisuje
	// metody method1 a method3
	Interface1
	method3()
}

type Interface4 interface {
	// toto rozhraní ve skutečnosti předepisuje
	// metody method1, method2 a method3
	Interface2
	Interface3
}
