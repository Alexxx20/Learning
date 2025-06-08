// Crie uma função que retorne um int
// Crie outra função que retorne um int e uma string
// Chame as duas funções
// Demonstre seus resultados
package main

import "fmt"

func main() {
	bonus, nome := bonus("Alex", 21)
	fmt.Printf("%v %v %v", soma(10), nome, bonus)
}

func soma(x ...int) int {
	i := len(x) - 1
	soma := 0
	for i >= 0 {
		soma += x[i]
		i--
	}
	return soma
}

func bonus(nome string, idade int) (int, string) {
	bonus := idade * len(nome)
	return bonus, nome
}
