// První kroky v jazyku Go
//
// - import symbolů ze dvou balíčků
// - každý balíček je naimportován přes samostatný příkaz "import"

package main

// import jediného balíčku
import (
	"fmt"
	"os"
)

// import jediného balíčku

func main() {
	// zde můžeme používat funkce z obou naimportovaných balíčků
	fmt.Println("Hello, world!")
	os.Exit(0)
}
