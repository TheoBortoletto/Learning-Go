package main

import "fmt"

type Pessoa struct {
	nome    string
	idade   int
	salario int
	cargo   string
}

func main() {

	fmt.Println("\n\t\tAcessar membros da estrutura: \n")

	/*Para acessar qualquer membro de uma estrutura,
	use o operador ponto (.) entre o nome da variável
	da estrutura e o membro da estrutura:*/

	var pessoa1 Pessoa
	var pessoa2 Pessoa

	// PESSOA 1:
	pessoa1.nome = "Théo Bortoletto"
	pessoa1.idade = 20
	pessoa1.salario = 1300
	pessoa1.cargo = "Estagiário de TI"

	// PESSOA 2:
	pessoa2.nome = "Eduarda Medeiros"
	pessoa2.idade = 22
	pessoa2.salario = 2500
	pessoa2.cargo = "Analista de dados"

	// ACESSANDO E IMPRIMINDO PESSOA 1 info:
	fmt.Println("Nome: ", pessoa1.nome)
	fmt.Println("Idade: ", pessoa1.idade)
	fmt.Println("Salário: ", pessoa1.salario)
	fmt.Println("Cargo: ", pessoa1.cargo, "\n")

	// ACESSANDO E IMPRIMINDO PESSOA 2 info:
	fmt.Println("Nome: ", pessoa2.nome)
	fmt.Println("Idade: ", pessoa2.idade)
	fmt.Println("Salário: ", pessoa2.salario)
	fmt.Println("Cargo: ", pessoa2.cargo)

}
