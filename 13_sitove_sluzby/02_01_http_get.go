// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - použití HTTP metody GET

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	// přečtení celého těla odpovědi
	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Response body can't be read")
	} else {
		// převod sekvence bajtů na řetězec
		body := string(rawBody)
		// tisk zpracovaného těla odpovědi serveru
		fmt.Println(body)
	}
}
