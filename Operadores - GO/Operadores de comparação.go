package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tOperadores de comparação: \n")

	/*Operadores de comparação são usados ​​para comparar dois valores.
	Nota: O valor de retorno de uma comparação é verdadeiro (1) ou falso (0).
	No exemplo a seguir, usamos o operador maior que ( '>' ) para descobrir se 5 é maior que 3:*/

	x := 5
	y := 3

	fmt.Println(x > y) // Retorna 1 (TRUE) porque 5 é maior que 3
}

/*

Uma lista de todos os operadores de comparação:

Operator	Name				        Example

==		|	Equal to			|	x == y
!=		|	Não é igual a		|	x != y
>		|	Maio que			|   x > y
<		|	Menor que			|	x < y
>=		|	Maior ou igual a	|   x >= y
<=		|	Menor ou igual a	|	x <= y

*/
