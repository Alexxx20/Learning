// Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
// Nome
// Sobrenome
// Sabores favoritos de sorvete
// Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
package main

import "fmt"

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresFavoritos []string
}

func main() {
	pessoa1 := Pessoa{
		nome:             "Talles",
		sobrenome:        "Sousa Braga",
		saboresFavoritos: []string{"Cajá", "Maçã verde", "Baunilia"},
	}
	pessoa2 := Pessoa{
		nome:             "Gyan Carlos",
		sobrenome:        "Matheus Oliveira",
		saboresFavoritos: []string{"Chocomenta", "Chocolate"},
	}
	fmt.Printf("%v %v", pessoa1, pessoa2)
}
