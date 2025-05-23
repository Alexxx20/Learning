// Usando uma literal composta:
// Crie uma slice de tipo int
// Atribua 10 valores a ela
// Utilize range para demonstrar todos estes valores.
// E utilize format printing para demonstrar seu tipo.

package main

import "fmt"

func main() {
	fatia := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		fatia[i] = i + 1
	}
	for _, v := range fatia {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("Supported values type:%T", fatia)
}
