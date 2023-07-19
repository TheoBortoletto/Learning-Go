package main

import (
	"fmt"
)

func myFunciton(x int, y int) int {
	return x + y
}

func main() {

	fmt.Println("\n\t\tFuncão 'retur': \n")

	/*Se você deseja que a função retorne um valor,
	você precisa definir o tipo de dados do valor de
	retorno (como int, string, etc) e também usar a
	returnpalavra-chave dentro da função:

	SINTAXE

	func nomeFunção (param1 tipo, param2 tipo) tipo {
		código a ser executado
		return output
	}*/

	fmt.Println(myFunciton(1, 2))

}
