// Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string
// com identificador "esporteFavorito".

package main

import (
	"fmt"
	"strings"
)

func main() {
	var esporteFavorito string
	println("Qual seu esporte favorito?")
	fmt.Scan(&esporteFavorito)

	switch strings.ToLower(esporteFavorito) {
	case "futebol":
		println("Bora jogar bola")
	case "vôlei":
		println("Monta a rede")
	case "basquete":
		println("Sei nem para onde vai aí")
	default:
		println("Tô não")
	}
}
