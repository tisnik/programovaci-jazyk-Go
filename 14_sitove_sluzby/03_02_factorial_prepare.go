// Vývoj síťových aplikací v programovacím jazyku Go
//
// - webová aplikace pro výpočet faktoriálu (příprava)

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type FactorialPageDynContent struct {
	Result int
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("factorial.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
		return
	}

	dynData := FactorialPageDynContent{1}
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
