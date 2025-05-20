// Faça um programa que leia o nome de um vendedor, o seu salário fixo e o total de vendas efetuadas por ele no mês (em dinheiro). Sabendo que este vendedor
// ganha 15% de comissão sobre suas vendas efetuadas, informar o total a receber no final do mês, com duas casas decimais.
package main

import "fmt"

func main() {
	var nome string
	var salario, vendas float64
	fmt.Scan(&nome)
	fmt.Scan(&salario)
	fmt.Scan(&vendas)

	total := salario + (vendas * 0.15)
	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
