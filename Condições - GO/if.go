package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tA declaração 'if': \n")

	/*Use a instrução 'if' para especificar um bloco de código Go a ser executado se uma condição for true.

	SINTAXE:

	if condição {
		código de execução se for verdadeiro
	}

	*/

	if 30 > 29 {
		fmt.Println("30 é maior que 29. \n")
	}

	x := 20
	y := 18

	if x > y {
		fmt.Println(x, "é maior que", y, "\n")
	}

}
