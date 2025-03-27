// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - HTTP server posílající sekvenci bajtů

package main

import (
	"fmt"
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	response := "Hello world!\n"
	// posílat budeme sekvenci bajtů
	writer.Write([]byte(response))
}

func main() {
	// registrace callback funkce volané při přístupu na endpoint
	http.HandleFunc("/", mainEndpoint)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
