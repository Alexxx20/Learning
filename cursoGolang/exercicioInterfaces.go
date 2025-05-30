package main

import "fmt"

type PessoaI struct {
	nome      string
	sobrenome string
	idade     int
}

type ProfissaoI struct {
	PessoaI
	salario float64
}

type Dentista struct {
	ProfissaoI
	CRO int
}

type Arquiteto struct {
	ProfissaoI
	CAU int
}

func (a Arquiteto) info() {
	fmt.Printf("%v %v %v %v %v\n", a.nome, a.sobrenome, a.idade, a.salario, a.CAU)
}
func (d Dentista) info() {
	fmt.Printf("%v %v %v %v %v\n", d.nome, d.sobrenome, d.idade, d.salario, d.CRO)
}

type apresentacao interface {
	info()
}

func apresentacaoI(a apresentacao) {
	a.info()
}

func main() {
	id_1 := Arquiteto{
		ProfissaoI: ProfissaoI{
			PessoaI: PessoaI{
				nome:      "Lucas",
				sobrenome: "Neto",
				idade:     21,
			},
			salario: 3000,
		},
		CAU: 2319283,
	}
	id_2 := Dentista{
		ProfissaoI: ProfissaoI{
			PessoaI: PessoaI{
				nome:      "Natália",
				sobrenome: "Silvestre",
				idade:     22,
			},
			salario: 3200,
		},
		CRO: 231922,
	}

	// id_1.info()
	id_2.info()
	apresentacaoI(id_1)

}
