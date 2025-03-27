// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net
// - parsing IPv4 a IPv6 adresy

package main

import (
	"fmt"
	"net"
)

func parseIP(address string) {
	ip := net.ParseIP(address)
	if ip == nil {
		fmt.Printf("%s: ParseIP failure\n", address)
	} else {
		fmt.Printf("%s -> %v\n", address, ip)
	}
}

func main() {
	parseIP("127.0.0.1")
	parseIP("1000::68")
	parseIP("fe80::224:d7ff:fe83:1f28")
	parseIP("fe80:0000:0000:0000:224:d7ff:fe83:1f28")
	parseIP("fe80:0000:0000:0000:0000:0000:0000:0001")
	parseIP("256.0.0.0")
	parseIP("foo.bar.baz")
}
