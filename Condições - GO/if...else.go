package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tA declaração 'if...else': \n")

	/*Use a elseinstrução para especificar um bloco de código a ser executado se a condição for false.

	SINTAXE:

	if condição {
		código a ser executado se a condição for verdadeira
	} else {
		código a ser executad se a condição for falsa
	}

	O 'ELSE' EM UMA LINHA DIFERENTE GERARÁ ERRO

	*/

	hora := 20

	if hora < 18 {
		fmt.Println("Bom dia!!! \n")
	} else {
		fmt.Println("Boa Tarde!!! \n")
	}

	temp := 24

	if temp > 15 {
		fmt.Println("Está quente lá fora. \n")
	} else {
		fmt.Println("Não está quente lá fora. \n")
	}
}
