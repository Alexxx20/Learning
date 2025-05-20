// Neste problema, deve-se ler o código de uma peça 1, o número de peças 1, o valor unitário de cada peça 1,
// o código de uma peça 2, o número de peças 2 e o valor unitário de cada peça 2. Após, calcule e mostre o valor a ser pago
package main

import "fmt"

func main() {
	var codigo1, codigo2, qtdPeca1, qtdPeca2, valor1, valor2 float64
	fmt.Scan(&codigo1, &qtdPeca1, &valor1)
	fmt.Scan(&codigo2, &qtdPeca2, &valor2)

	valorTotal := (qtdPeca1 * valor1) + (qtdPeca2 * valor2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valorTotal)

}
