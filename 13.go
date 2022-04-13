package main

import "fmt"

func main() {
	a := 3
	b := 5

	a = a + b
	b = a - b
	a = a - b
	//a, b = b, a
	fmt.Println("A =", a, "B =", b)
}
