// Crie e use um struct anônimo.
// Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
package main

import "fmt"

func main() {
	comanda := struct {
		mesa            int
		pedidos         []string
		quemPagouQuanto map[string]float64
	}{
		mesa:            4,
		pedidos:         []string{"Baião de 2", "x4 Espeto de boi", "x10 Skol"},
		quemPagouQuanto: map[string]float64{"Alex": 50.5, "Natália": 20.5},
	}
	fmt.Println(comanda)
}
