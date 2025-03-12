// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - HTTP server se dvěma endpointy
// - korektní práce se stavovou proměnnou
// - synchronizace přístupu přes mutex

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func counterEndpoint(writer http.ResponseWriter, request *http.Request) {
	// uzamčít přístup pro další gorutiny
	mutex.Lock()
	counter++
	fmt.Fprintf(writer, "Counter: %d\n", counter)
	// odemčít přístup pro další gorutiny
	mutex.Unlock()
}

func main() {
	// registrace callback funkce volané při přístupu na endpointy
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/counter", counterEndpoint)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
