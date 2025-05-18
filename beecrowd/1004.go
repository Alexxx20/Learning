// Leia dois valores inteiros. A seguir, calcule o produto entre estes dois valores e atribua esta operação à variável PROD.
// A seguir mostre a variável PROD com mensagem correspondente.
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	prod := a * b
	fmt.Printf("PROD = %v\n", prod)

}
