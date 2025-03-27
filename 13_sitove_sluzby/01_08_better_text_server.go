// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net
// - server posílající klientovi textová data
// - uzavření připojení v bloku defer

package main

import (
	"fmt"
	"net"
)

func main() {
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

		if err != nil {
			fmt.Println("connection refused!")
		} else {
			fmt.Println("connection accepted")
			go func(c net.Conn) {
				// zajistit uzavření připojení
				defer c.Close()

				// poslat klientovi celý textový řádek
				fmt.Fprintf(c, "Holla\n")
				fmt.Println("sent whole text line")
			}(conn)
		}
	}
}
