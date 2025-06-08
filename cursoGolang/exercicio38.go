// Crie um tipo struct "pessoa" que contenha os campos:
// nome
// sobrenome
// idade
// Crie um método para "pessoa" que demonstre o nome completo e a idade;
// Crie um valor de tipo "pessoa";
// Utilize o método criado para demonstrar esse valor.
package main

import "fmt"

type PessoaV struct {
	nome      string
	sobrenome string
	idade     int
}

func (p PessoaV) infor() {
	nomeC := p.nome + " " + p.sobrenome
	idade := p.idade
	fmt.Printf("Nome completo: %v\nIdade: %v", nomeC, idade)
}

var alex PessoaV

func main() {
	alex.nome = "Alex"
	alex.sobrenome = "Lucena"
	alex.idade = 21

	alex.infor()
}
