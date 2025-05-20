// Escreva um programa que leia três valores com ponto flutuante de dupla precisão: A, B e C. Em seguida, calcule e mostre:
// a) a área do triângulo retângulo que tem A por base e C por altura.
// b) a área do círculo de raio C. (pi = 3.14159)
// c) a área do trapézio que tem A e B por bases e C por altura.
// d) a área do quadrado que tem lado B.
// e) a área do retângulo que tem lados A e B.

package main

import (
	"fmt"
	"math"
)

const pi = 3.14159

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	tri := areaTriangulo(a, c)
	cir := areaCirculo(c)
	tra := areaTrapezio(a, b, c)
	qua := areaQuadrado(b)
	ret := areaRetangulo(a, b)

	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n", tri, cir, tra, qua, ret)
}

func areaTriangulo(b float64, h float64) float64 {
	return (b * h) / 2
}

func areaCirculo(r float64) float64 {
	return pi * math.Pow(r, 2)
}

func areaTrapezio(bigB float64, b float64, h float64) float64 {
	return ((bigB + b) * h) / 2
}

func areaQuadrado(l float64) float64 {
	return l * l
}

func areaRetangulo(l1 float64, l2 float64) float64 {
	return l1 * l2
}
