// Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
// Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
// Do quinto ao último item do slice (incluindo o último item!)
// Do segundo ao sétimo item do slice (incluindo o sétimo item!)
// Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
// Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import "fmt"

func main() {
	fatias := make([]int, 10, 10)
	for i, _ := range fatias {
		fatias[i] = i * 2
	}
	fmt.Println(fatias)
	fmt.Printf("Primeiro ao terceiro item: %v\nQuinto ao último: %v\nSegundo ao sétimo: %v\nTerceiro ao penúltimo: %v\nTerceiro ao penúltimo (usando len()): %v", fatias[:3], fatias[4:], fatias[1:7], fatias[3:9], fatias[3:len(fatias)-1])
}
