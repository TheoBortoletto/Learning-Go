package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tTipos de dados float: \n")

	/*O tipo de dados float tem duas palavras-chave:

	Type:		 Size:			Range:
	float32		32 bits		-3.4e+38 to 3.4e+38.
	float64		64 bits		-1.7e+308 to +1.7e+308.

	Dica: O tipo padrão para float é float64. Se você não especificar
	um tipo, o tipo será float64.*/

	fmt.Println("\tA palavra-chave 'float32': \n")

	var x float32 = 132.78
	var y float32 = 3.4e+38

	fmt.Printf("Tipo: %T, valor: %v \n", x, x)
	fmt.Printf("Tipo: %T, valor: %v \n\n", y, y)

	fmt.Println("\tA palavra-chave 'float64':\n")

	/*O float64 pode armazenar um conjunto maior de números do que float32.*/

	var num float64 = 1.7e+308

	fmt.Printf("Tipo: %T, valor: %v \n", num, num)

}
