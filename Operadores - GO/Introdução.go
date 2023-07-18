package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tOperadores GO: \n")

	/*Operadores são usados ​​para realizar operações em variáveis ​​e valores.

	O operador '+' soma dois valores, como no exemplo abaixo:*/

	var num = 15 + 25
	fmt.Println(num, "\n")

	/*Embora o +operador seja frequentemente usado
	para somar dois valores, ele também pode ser
	usado para somar uma variável e um valor, ou
	uma variável e outra variável:*/

	var (
		sum1 = 100 + 50
		sum2 = sum1 + 250
		sum3 = sum2 + sum2
	)

	fmt.Println(sum3)

}
