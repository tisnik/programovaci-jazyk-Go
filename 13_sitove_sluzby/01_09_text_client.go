// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net
// - jednoduchý klient akceptující celý textový řádek

package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Connection refused!")
	} else {
		fmt.Fprintf(conn, "Hello")
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(status, err)
		if err != nil {
			fmt.Println("No response!")
		} else {
			fmt.Println(status)
		}
	}
}
