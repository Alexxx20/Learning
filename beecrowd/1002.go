// A fórmula para calcular a área de uma circunferência é: area = π . raio2. Considerando para este problema que π = 3.14159:

// Efetue o cálculo da área, elevando o valor de raio ao quadrado e multiplicando por π.
package main

import (
	"fmt"
	"math"
)

const n float64 = 3.14159

func main() {
	var a float64
	fmt.Scan(&a)
	area := n * math.Pow(a, 2)
	fmt.Printf("A=%.4f\n", area)

}
