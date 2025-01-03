// Striktní kontroly prováděné překladačem při konverzi mezi datovými typy.

package main

import "fmt"

func main() {
	var a uint16 = 255

	// pokus o implicitní konverzi mezi typy uint16 a uint8
	var b uint8 = a

	fmt.Println(a)
	fmt.Println(b)
}
