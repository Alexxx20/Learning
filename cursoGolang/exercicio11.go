// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// Demonstre estes valores

package main

import "fmt"

const (
	_ = iota + 2025
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v %v %v %v", a, b, c, d)
}
