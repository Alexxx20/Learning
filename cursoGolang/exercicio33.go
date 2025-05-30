// Crie um novo tipo: veículo
// O tipo subjacente deve ser struct
// Deve conter os campos: portas, cor
// Crie dois novos tipos: caminhonete e sedan
// Os tipos subjacentes devem ser struct
// Ambos devem conter "veículo" como struct embutido
// O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
// O tipo sedan deve conter um campo bool chamado "modeloLuxo"
// Usando os structs veículo, caminhonete e sedan:
// Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
// Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
// Demonstre estes valores.
// Demonstre um único campo de cada um dos dois.
package main

import "fmt"

type Veiculo struct {
	portas int
	cor    string
}
type Caminhonete struct {
	Veiculo
	tracaoNasQuatroRodas bool
}
type Sedan struct {
	Veiculo
	modeloLuxo bool
}

func main() {
	saveiro := Caminhonete{Veiculo{2, "vermelha"}, true}
	hb20 := Sedan{Veiculo{4, "preta"}, false}

	fmt.Printf("Saveiro: %v\nHB20: %v\n", saveiro, hb20)
	fmt.Printf("%v, %v", saveiro.Veiculo, hb20.cor)
}
