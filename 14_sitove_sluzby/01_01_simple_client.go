// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net
// - síťový klient, který přečte ze serveru sekvenci bajtů

package main

import (
	"fmt"
	"net"
)

func main() {
	// navázání připojení
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Connection refused!")
		return
	}

	// alokace bufferu o velikosti jednoho bajtu
	// alternativní (lepší) způsob: b := make([]byte, 0, 4096)
	var b [1]byte

	// pozor - nedokážeme omezit počet přečtených bajtů!
	n, err := conn.Read(b[:])

	// kontrola, zda byl bajt přenesen
	if err != nil {
		fmt.Println("No response!")
		return
	}
	if n == 1 {
		// přečetl se jediný bajt - v pořádku
		fmt.Printf("Received %d byte: %v\n", n, b)
	} else {
		// problematická část - alokovaný buffer
		// nemusí mít dostatečnou kapacitu
		fmt.Printf("Received %d bytes: %v\n", n, b)
	}
}
