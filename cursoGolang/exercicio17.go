// Crie um programa que demonstre o funcionamento da declaração if.
package main

import "fmt"

func main() {
	var idade int
	println("Qual a sua idade?")
	fmt.Scan(&idade)

	if idade < 16 {
		println("Você é muito jovem para votar")
	}
}
