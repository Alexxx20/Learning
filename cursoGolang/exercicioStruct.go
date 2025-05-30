package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
	cpf   string
}

type funcionario struct {
	pessoa
	salario  float64
	cod      int
	pis      int
	admissao string
}

var funcionario1 funcionario

func main() {
	funcionario1.nome = "Natália"
	funcionario1.idade = 21
	funcionario1.salario = 2020.20
	funcionario1.cod = 1278310
	funcionario1.pis = 189380
	funcionario1.admissao = "12/13/2022"
	fmt.Print(funcionario1.pessoa)
}
