package main

import "fmt"

func main() {
	var doces = []string{"brigadeiro", "mousse", "pudim", "doce de leite", "mariola"}
	for i := 0; i < len(doces); i++ {
		fmt.Println(doces[i])
	}
	for idx, valor := range doces {
		fmt.Println(idx, valor)
	}
}
