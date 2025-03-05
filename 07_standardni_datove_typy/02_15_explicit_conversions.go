// Základní datové typy v jazyce Go
//
// - explicitní konverze mezi hodnotami různých datových typů

package main

import "fmt"

func main() {
	// příprava proměnných, jejichž hodnoty se budou později
	// konvertovat na jiné datové typy
	var a int8 = -10
	var signedInt int32 = -100000
	var unsignedInt uint32 = 100000
	var e float32 = 1e4
	var f float64 = 1.5e30

	// konverze hodnot
	var x int32 = int32(a)

	// zde dochází k zaokrouhlení
	var y int32 = int32(e)

	// zde dochází ke ztrátě přesnosti
	// a může dojít i k přetečení výsledku
	var z float32 = float32(f)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// konverze hodnot, dochází zde ke zmenšení
	// rozahu a k přetečení výsledků
	var b2 uint8 = uint8(signedInt)
	var b3 uint8 = uint8(unsignedInt)

	fmt.Println(b2)
	fmt.Println(b3)
}
