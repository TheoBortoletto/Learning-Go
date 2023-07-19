package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tMulti-case: \n")

	/*É possível ter vários valores para cada case na instrução switch:

		SINTAXE:

		switch expression {
	case x,y:
		code block if expression is evaluated to x or y
	case v,w:
	    code block if expression is evaluated to v or w
	case z:
	...
	default:
	    code block if expression is not found in any cases
	}
	*/

	dia := 5

	switch dia {
	case 1, 3, 5:
		fmt.Println("Dia da semana ímpar.")
	case 2, 4:
		fmt.Println("Dia da semana par.")
	case 6, 7:
		fmt.Println("Final de semana.")
	default:
		fmt.Println("Número do dia da semana inválido.")

	}
}
