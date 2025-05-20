package main

import (
	"fmt"
)

var x int
var fo, f uint64

func main() {
	fo, f = 0, 1
	fmt.Scan(&x)
	fmt.Printf("1 - %v\n2 - %v\n", fo, f)
	for i := 0; i < x; i++ {
		fo, f = f, fo+f
		fmt.Printf("%v - %v\n", i+3, f)
	}
}
