package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tFunções de Saída: \n")

	/*A Linguagrm 'GO' tem três funções para produzir texto

	- Print()
	- Println()
	- Printf()
	*/

	fmt.Println("\tA função 'Print()':\n")

	/*A Print()função imprime seus
	argumentos com seu formato padrão
	sem coloca-los em uma nova linha*/

	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)

	fmt.Println("\n\n\tA função 'Println()':\n")

	/*A Println()função é semelhante a
	Print()com a diferença de que um
	espaço em branco é adicionado entre
	os argumentos e uma nova linha é
	adicionada no final:*/

	var x, y string = "Hello", "World"

	fmt.Println(x, y)

	fmt.Println("\n\t A função Printf():\n ")

	/*A Printf()função primeiro formata seu argumento
	com base no verbo de formatação fornecido e, em
	seguida, os imprime.

		Aqui vamos usar dois verbos de formatação:

	 - %v é usado para imprimir o valor dos argumentos
	 - %T é usado para imprimir o tipo dos argumentos*/

	var a string = "Olááá"
	var b int = 20

	fmt.Printf("a tem valor: %v e tipo: %T \n", a, a)
	fmt.Printf("b tem valor: %v e tipo: %T \n", b, b)
}
