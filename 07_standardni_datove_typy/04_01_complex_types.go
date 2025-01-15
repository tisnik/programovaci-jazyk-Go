// Základní datové typy v jazyce Go
//
// - deklarace proměnných typu komplexní číslo

package main

import "fmt"

func main() {
	var a complex64 = -1.5 + 0i
	var b complex64 = 1.5 + 1000i
	var c complex64 = 1e30 + 1e30i
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// pouze imaginární složka
	var d complex64 = 1i
	fmt.Println(d)

	var e complex128 = -1.5 + 0i
	var f complex128 = 1.5 + 1000i
	var g complex128 = 1e300 + 1e300i

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// lze použít i další formy zápisu reálné a imaginární složky
	// již je známe z popisu hodnot typu float32 a float64

	var i complex128 = 1000_0000e2 + 1000i
	fmt.Println(i)

	// pouze reálná složka
	var j complex128 = 0xFF_FF
	fmt.Println(j)

	// pouze imaginární složka
	var k complex128 = 0xFF_FFp1i
	fmt.Println(k)
}
