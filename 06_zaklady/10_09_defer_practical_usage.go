// Základy programovacího jazyka Go
//
// - praktické použití konstrukce defer
// - realizace kopie obsahu souboru
// - delší varianta vypisující prováděné operace

package main

import (
	"fmt"
	"io"
	"os"
)

// pomocná funkce pro uzavření souboru
func closeFile(file *os.File) {
	fmt.Printf("Closing file '%s'\n", file.Name())
	file.Close()
}

func copyFile(srcName, dstName string) (written int64, err error) {
	// otevření zdrojového souboru pro čtení
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("Cannot open file '%s' for reading\n", srcName)
		return
	}
	fmt.Printf("File '%s' opened for reading\n", srcName)

	// zajistit uzavření souboru za všech okolností
	defer closeFile(src)

	// otevření cílového souboru pro zápis
	dst, err := os.Create(dstName)
	if err != nil {
		fmt.Printf("Cannot create destination file '%s'\n", dstName)
		return
	}
	fmt.Printf("File '%s' opened for writing\n", dstName)

	// zajistit uzavření souboru za všech okolností
	defer closeFile(dst)

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
