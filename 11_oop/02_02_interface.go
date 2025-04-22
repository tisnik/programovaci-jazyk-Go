// Koncept rozhraní v programovacím jazyku Go
//
// - rozhraní se dvěma předepsanými metodami

type Shape interface {
	area() float64
	perimeter() float64
}
