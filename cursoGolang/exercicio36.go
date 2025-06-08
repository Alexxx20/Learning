// Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// Passe um valor do tipo slice of int como argumento para a função;
// Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// Passe um valor do tipo slice of int como argumento para a função.
package main

func main() {
	lista := []int{1, 2, 3, 4, 5}
	println(somatorio(lista...), somatorioaSlice(lista))
}

func somatorio(x ...int) int {
	soma := 0
	for i := 0; i < len(x); i++ {
		soma += x[i]
	}
	return soma
}

func somatorioaSlice(x []int) int {
	soma := 0
	for _, i := range x {
		soma += i
	}
	return soma
}
