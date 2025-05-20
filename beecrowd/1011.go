// Faça um programa que calcule e mostre o volume de uma esfera sendo fornecido o valor de seu raio (R).
// A fórmula para calcular o volume é: (4/3) * pi * R3. Considere (atribua) para pi o valor 3.14159.
package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14159

func main() {
	var raio float64
	fmt.Scan(&raio)
	volume := (4.0 / 3.0) * pi * math.Pow(raio, 3)
	fmt.Printf("VOLUME = %.3f\n", volume)
}
