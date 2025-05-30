// Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// Demonstre os valores do map utilizando range.
// Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
package main

import "fmt"

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresFavoritos []string
}

var pessoas_cadastradas = make(map[string]Pessoa)

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
	pessoas_cadastradas[pessoa1.sobrenome] = pessoa1
	pessoas_cadastradas[pessoa2.sobrenome] = pessoa2

	for e, i := range pessoas_cadastradas {
		fmt.Printf("%v %v\n", i.nome, e)
		for _, j := range i.saboresFavoritos {
			fmt.Printf("%v\n", j)
		}
	}
}
