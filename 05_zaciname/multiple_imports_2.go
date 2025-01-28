// První kroky v jazyku Go
//
// - import symbolů ze dvou balíčků
// - oba balíčky jsou naimportovány jediným příkazem "import"

package main

// import dvou balíčků v jediném bloku "import"
import (
	"fmt"
	"os"
)

func main() {
	// zde můžeme používat funkce z obou naimportovaných balíčků
	fmt.Println("Hello, world!")
	os.Exit(0)
}
