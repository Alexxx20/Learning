// Utilizando a solução do exercício anterior:
// Em package-level scope, atribua os seguintes valores às variáveis:
// para "x" atribua 42
// para "y" atribua "James Bond"
// para "z" atribua true
// Na função main:
// Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
// Demonstre a variável "s".

package main

import "fmt"

var x int
var y string
var z bool

func main(){
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}