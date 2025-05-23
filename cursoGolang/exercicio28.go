// Crie um map com key tipo string e value tipo []string.
// Key deve conter nomes no formato sobrenome_nome
// Value deve conter os hobbies favoritos da pessoa
// Demonstre todos esses valores e seus índices.

package main

import (
	"fmt"
)

func main() {
	pessoas := map[string][]string{"Lucena Alex": []string{"Caçar", "Jogar"}, "Silvestre Natália": []string{"Séries", "Ler"}, "Sousa João": []string{"Pernoitar", "Beber"}}
	for k, v := range pessoas {
		fmt.Printf("Nome: %v\nHobbies: %v\n", k, v)
	}

}
