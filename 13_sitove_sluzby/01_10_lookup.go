// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net
// - překlad doménového jména na IP adresy

package main

import (
	"fmt"
	"net"
)

func performLookup(address string) {
	// pokus o překlad doménového jména na IP adresy
	ip, err := net.LookupIP(address)
	if err != nil {
		fmt.Printf("%s: Lookup failure\n", address)
	} else {
		fmt.Printf("%s: %v\n", address, ip)
	}
}

func main() {
	performLookup("8.8.8.8")
	performLookup("localhost")
	performLookup("root.cz")
	performLookup("seznam.cz")
	performLookup("unknown.foo.bar")
}
