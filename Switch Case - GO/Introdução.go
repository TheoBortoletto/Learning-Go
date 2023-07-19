package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tSwitch Case: \n")

	/*
		Use a instrução switch para selecionar um dos
		muitos blocos de código a serem executados.

		A instrução switch em Go é semelhante às de C, C++, Java,
		JavaScript e PHP. A diferença é que ele executa apenas o
		caso correspondente, portanto não precisa de uma breakinstrução.
	*/

	fmt.Println("\n\tSingle Case: \n")

	/*
					SINTAXE:

					switch expression {
					case x:
					    code block
					case y:
					 code block
					case z:
					...
					default:
						code block
				}

				É assim que funciona:

		- A expressão é avaliada uma vez

		- O valor da expressão switch é
		comparado com os valores de cada case

		- Se houver uma correspondência,
		o bloco de código associado é executado

		- A palavra-chave 'default' é opcional.
		Ele especifica algum código a ser executado
		se não houver case correspondência

	*/

}
