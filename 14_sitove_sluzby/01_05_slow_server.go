// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net
// - server odpovídající klientovi opožděně
// - omezení na místní připojení
// - po navázání připojení se odešle jeden bajt a připojení se ukončí

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// čítač počtu připojení
	cnt := byte(0)

	// server bude naslouchat na portu 1234
	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Can't open the port!")
		return
	}
	defer l.Close()

	// postupné zpracování požadavků od klientů
	for {
		// potvrzení připojení
		conn, err := l.Accept()

		// simulace déletrvající operace
		fmt.Print("working....")
		time.Sleep(2 * time.Second)
		fmt.Println("done")

		if err != nil {
			fmt.Println("connection refused!")
		} else {
			fmt.Println("connection accepted")
			// zajistit uzavření připojení
			defer conn.Close()

			// poslat můžeme jen jeden bajt
			// - klient totiž nemá větší buffer
			var b = []byte{cnt}
			conn.Write(b)
			fmt.Printf("sent byte [%d]\n", cnt)
			// změna obsahu posílaného bajtu
			cnt++
		}
		fmt.Println("connection closed")
	}
}
