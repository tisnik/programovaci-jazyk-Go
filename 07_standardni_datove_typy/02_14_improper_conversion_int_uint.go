// Základní datové typy v jazyce Go
//
// - striktní kontroly prováděné překladačem při konverzi
//   mezi datovými typy

package main

import "fmt"

func main() {
	var a int8 = 100

	// pokus o implicitní konverzi mezi typem int8 a uint8
	var b uint8 = a

	fmt.Println(a)
	fmt.Println(b)

	var c uint8 = 100

	// pokus o implicitní konverzi mezi typem uint8 a int8
	var d int8 = c

	fmt.Println(c)
	fmt.Println(d)
}
