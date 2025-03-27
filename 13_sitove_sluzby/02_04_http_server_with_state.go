// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - HTTP server se dvěma endpointy
// - nekorektní práce se stavovou proměnnou

package main

import (
	"fmt"
	"io"
	"net/http"
)

var counter int

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func counterEndpoint(writer http.ResponseWriter, request *http.Request) {
	counter++
	fmt.Fprintf(writer, "Counter: %d\n", counter)
}

func main() {
	// registrace callback funkce volané při přístupu na endpointy
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/counter", counterEndpoint)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
