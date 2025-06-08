package main

import "fmt"

func main() {
	fmt.Print(fatorial(4))
}

func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fatorial(n-1)
}
