// Crie uma variável de tipo string utilizando uma raw string literal.
// Demonstre-a.

package main

import "fmt"

var texto = `\t \n \k \ç \p            A`

func main() {
	fmt.Println(texto)
}
