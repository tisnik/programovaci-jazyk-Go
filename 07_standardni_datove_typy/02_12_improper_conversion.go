// Striktní kontroly prováděné překladačem při konverzi mezi datovými typy.

package main

import "fmt"

func main() {
	var a uint8 = 255

	// pokus o implicitní konverzi mezi typy uint8 a uint16
	var b uint16 = a

	fmt.Println(a)
	fmt.Println(b)
}
