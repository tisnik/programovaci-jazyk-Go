// Technologie GopherJS
//
// - program typu "Hello, world!"
// - použití funkce fmt.Println

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, world!")
	js.Global().Call("alert", "Hello, world!")
}
