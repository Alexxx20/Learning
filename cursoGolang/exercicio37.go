// Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
package main

import "fmt"

func main() {
	defer fmt.Println("Tchau!")
	fmt.Println("Oi")
}
