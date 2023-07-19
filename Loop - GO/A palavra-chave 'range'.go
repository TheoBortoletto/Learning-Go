package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tA palavra-chave 'range': \n")

	/*A palavra-chave 'range' é usada para iterar mais
		facilmente em uma matriz, fatia ou mapa. Ele retorna o
		índice e o valor.

		A palavra-chave 'range' é usada assim:

		SINTAXE:

		for index, value := array|slice|map {
	    code to be executed for each iteration
	}

	*/

	fmt.Println("Exemplo 1: \n")

	/*Este exemplo usa 'range' para iterar sobre uma matriz e
	imprimir os índices e os valores em cada um
	( 'idx' armazena o índice, 'val' armazena o valor):*/

	frutas := [3]string{"Maçã", "Laranja", "Banana"}
	for idx, val := range frutas {
		fmt.Printf("%v\t%v\n\n", idx, val)
	}

	fmt.Println("Exemplo 2: \n")

	/*Dica: Para mostrar apenas o valor ou o índice,
	você pode omitir a outra saída usando um sublinhado ( _ ).*/

	fruits := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

}
