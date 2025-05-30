package main

import "fmt"

func main() {
	t, z := somenteImpares(soma, 1, 2, 3, 4, 5, 6, 7, 9, 10, 199, 12)
	fmt.Println(t, z)
}

func soma(x ...int) int {
	somatorio := 0
	for _, i := range x {
		somatorio += i
	}
	return somatorio
}

func somenteImpares(f func(x ...int) int, y ...int) (int, []int) {
	var impares []int
	for _, i := range y {
		if i%2 != 0 {
			impares = append(impares, i)
		}
	}
	total := f(impares...)

	return total, impares
}
