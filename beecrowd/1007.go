// Leia quatro valores inteiros A, B, C e D.
// A seguir, calcule e mostre a diferença do produto de A e B pelo produto de C e D segundo a fórmula: DIFERENCA = (A * B - C * D).
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	diferenca := (a * b) - (c * d)
	fmt.Printf("DIFERENCA = %v\n", diferenca)
}
