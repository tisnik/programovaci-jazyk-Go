// Technologie GopherJS
//
// - program typu "Hello, world!"
// - použití funkce println

package main

import (
	"syscall/js"
)

func main() {
	println("Hello, world!")
	js.Global().Call("alert", "Hello, world!")
}
