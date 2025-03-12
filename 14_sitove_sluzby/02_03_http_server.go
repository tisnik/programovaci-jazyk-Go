// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - nejjednodušší HTTP server s jediným endpointem

package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func main() {
	// registrace callback funkce volané při přístupu na endpoint
	http.HandleFunc("/", mainEndpoint)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
