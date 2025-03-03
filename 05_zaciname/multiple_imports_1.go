// První kroky v jazyku Go
//
// - import symbolů (funkcí atd.) ze dvou balíčků
// - každý balíček je naimportován přes samostatný příkaz "import"

package main

// import jediného balíčku
import "fmt"

// import jediného balíčku
import "os"

func main() {
	// zde můžeme používat funkce z obou naimportovaných balíčků
	fmt.Println("Hello, world!")
	os.Exit(0)
}
