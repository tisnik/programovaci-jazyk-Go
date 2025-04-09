// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - základní použití šablon

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// IndexPageDynContent holds all required information for a main page template
type IndexPageDynContent struct {
	Title  string
	Header string
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	// načtení šablony
	t, err := template.ParseFiles("index_template.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
		return
	}

	// data posílaná do šablonovacího systému
	dynData := IndexPageDynContent{Title: "Test", Header: "Welcome!"}

	// aplikace dat na šablonu
	err = t.Execute(writer, dynData)
	if err != nil {
		fmt.Println("Error executing template")
	}
}

func main() {
	// registrace callback funkce volané při přístupu na endpoint
	http.HandleFunc("/", mainEndpoint)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
