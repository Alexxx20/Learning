// Crie um programa que utilize a declaração switch, sem switch statement especificado.
package main

import "fmt"

func main() {
	var x int
	fmt.Println("Digite um número")
	fmt.Scan(&x)
	switch {
	case x%2 == 0 && x != 0:
		fmt.Printf("%v é par", x)
	case x%2 != 0 && x != 0:
		fmt.Printf("%v é ímpar", x)
	default:
		fmt.Printf("%v é nulo", x)
	}
}
