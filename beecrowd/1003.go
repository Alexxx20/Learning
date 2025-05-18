// Leia dois valores inteiros, no caso para variáveis A e B. A seguir, calcule a soma entre elas e atribua à variável SOMA.
// A seguir escrever o valor desta variável.

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	soma := a + b

	fmt.Printf("SOMA = %v\n", soma)
}
