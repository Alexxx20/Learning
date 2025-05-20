// Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import "fmt"

func main() {
	for i := 11; i < 100; i++ {
		fmt.Println(i % 4)
	}
}
