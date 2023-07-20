package main

import "fmt"

type Pessoa struct {
	nome    string
	idade   int
	cargo   string
	salario int
}

func main() {

	/*Você também pode passar uma estrutura como um argumento de função, assim:*/

	fmt.Println("\n\t\tPassar sctruct como argumento de função: \n")

	var pessoa1, pessoa2 Pessoa

	pessoa1.nome = "Théo Bortoletto"
	pessoa1.idade = 20
	pessoa1.cargo = "Estagiário de TI"
	pessoa1.salario = 1300

	pessoa2.nome = "Eduarda Medeiros"
	pessoa2.idade = 22
	pessoa2.cargo = "Analista de dados"
	pessoa2.salario = 2500

	printPessoa(pessoa1)
	fmt.Println()
	printPessoa(pessoa2)

}

func printPessoa(pessoa Pessoa) {
	fmt.Println("Nome: ", pessoa.nome)
	fmt.Println("Idade: ", pessoa.idade)
	fmt.Println("Cargo: ", pessoa.cargo)
	fmt.Println("Salário: ", pessoa.salario)
}
