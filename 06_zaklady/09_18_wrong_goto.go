// Základy programovacího jazyka Go
//
// - neplatné použití goto pro skok mezi bloky

package main

func main() {
	{
		// první blok
		// skok do dalšího bloku
		goto Target
	}
	{
		// druhý blok
	Target:
	}
}
