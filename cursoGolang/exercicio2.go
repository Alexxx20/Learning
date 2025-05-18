// Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:

// Identificador "x" deverá ter tipo int
// Identificador "y" deverá ter tipo string
// Identificador "z" deverá ter tipo bool

// Na função main:

// Demonstre os valores de cada identificador
// O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
// R: valor zero

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v %v %v\n", x, y, z);
}
