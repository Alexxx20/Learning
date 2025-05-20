// Escreva um programa que leia o número de um funcionário, seu número de horas trabalhadas, o valor que recebe por hora e
// calcula o salário desse funcionário. A seguir, mostre o número e o salário do funcionário, com duas casas decimais.
package main

import "fmt"

func main() {
	var num, horas int
	var salario float64
	fmt.Scan(&num)
	fmt.Scan(&horas)
	fmt.Scan(&salario)

	salario = salario * float64(horas)
	fmt.Printf("NUMBER = %v\nSALARY = U$ %.2f\n", num, salario)
}
