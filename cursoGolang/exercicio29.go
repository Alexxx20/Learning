// Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
package main

import (
	"fmt"
)

func main() {
	pessoas := map[string][]string{"Lucena Alex": []string{"Caçar", "Jogar"}, "Silvestre Natália": []string{"Séries", "Ler"}, "Sousa João": []string{"Pernoitar", "Beber"}}
	pessoas["Silva Alfredo"] = []string{"Piscina", "Praia", "Caçar tatu"}
	for k, v := range pessoas {
		fmt.Printf("Nome: %v\nHobbies: %v\n", k, v)
	}

}
