// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - vytištění hlavičky HTTP odpovědi

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// provedení HTTP dotazu typu GET
	response, err := http.Get("http://httpbin.org/uuid")
	if err != nil {
		fmt.Println("Connection refused")
	}
	defer response.Body.Close()

	// tisk informace o odpovědi HTTP serveru
	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Content length: %d\n", response.ContentLength)

	// tisk všech rozeznaných hlaviček
	for name, headers := range response.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}
