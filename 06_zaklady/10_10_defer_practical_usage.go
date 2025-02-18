// Základy programovacího jazyka Go
//
// - praktické použití konstrukce defer
// - realizace kopie obsahu souboru
// - kratší varianta bez výpisů prováděných operací

package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcName, dstName string) (written int64, err error) {
	// otevření zdrojového souboru pro čtení
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// zajistit uzavření souboru za všech okolností
	defer src.Close()

	// otevření cílového souboru pro zápis
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// zajistit uzavření souboru za všech okolností
	defer dst.Close()

	// vlastní kopie
	// (vrátí se počet zkopírovaných bajtů a chybová struktura)
	return io.Copy(dst, src)
}

// otestování, zda kopie obsahu souboru pracuje korektně
func testCopyFile(srcName, dstName string) {
	copied, err := copyFile(srcName, dstName)
	if err != nil {
		fmt.Printf("copyFile('%s', '%s') failed!!!\n", srcName, dstName)
	} else {
		fmt.Printf("Copied %d bytes\n", copied)
	}
	fmt.Println()
}

func main() {
	testCopyFile("10_09_defer_practical_usage.go", "new.go")
	testCopyFile("not_exists", "new.go")
	testCopyFile("10_09_defer_practical_usage.go", "")
	testCopyFile("10_09_defer_practical_usage.go", "/dev/full")
	testCopyFile("/dev/null", "new2.go")
}
