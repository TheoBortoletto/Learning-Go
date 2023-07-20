package main

import (
	"fmt"
)

/*Go aceita funções de recursão. Uma função é
recursiva se chama a si mesma e atinge uma condição
 de parada.

No exemplo a seguir, 'testcount()' é uma função que
chama a si mesma. Usamos a xvariável como os dados,
que incrementam com 1 ( x + 1) toda vez que recursamos.
A recursão termina quando a x variável for igual a 11 ( x == 11 ). */

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func main() {

	fmt.Println("\n\t\tFunções de recurssão: \n")

	testcount(1)

}
