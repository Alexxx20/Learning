// Crie um programa que:
// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// Demonstre esta outra variável em decimal, binário e hexadecimal

package main

import "fmt"

var x int = 1

func main() {
	fmt.Printf("%d %#b %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d %#b %#x", y, y, y)
}
