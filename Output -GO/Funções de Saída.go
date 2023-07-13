package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\tFunções de Saída: \n")

	/*A Linguagrm 'GO' tem três funções para produzir texto

	- Print()
	- Println()
	- Printf()
	*/

	fmt.Println("A função 'Print()':\n")

	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)
}
