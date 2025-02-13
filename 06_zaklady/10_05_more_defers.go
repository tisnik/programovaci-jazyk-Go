package main

import "fmt"

func onFinish(i int) {
	fmt.Printf("Defer #%2d\n", i)
}

func main() {
	defer onFinish(1)
	defer onFinish(2)
	defer onFinish(3)
	defer onFinish(4)
	fmt.Println("Finishing main() function")
}
