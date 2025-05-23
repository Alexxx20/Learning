// Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
// "Nome", "Sobrenome", "Hobby favorito"
// Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
package main

import "fmt"

func main() {
	pessoas := [][]string{{"Alex", "Lucena", "Jogar"}, {"Natália", "Silvestre", "Música"}, {"John", "Sousa", "Dormir"}}
	for _, v := range pessoas {
		fmt.Printf("Nome: %v %v\nHobby: %v\n", v[0], v[1], v[2])
	}
}
