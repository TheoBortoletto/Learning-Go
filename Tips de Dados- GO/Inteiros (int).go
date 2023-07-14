package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tTipos de dados inteiros: \n")

	/*O tipo de dados inteiro tem duas categorias:

	Inteiros com sinal - podem armazenar valores positivos e negativos
	Inteiros não assinados - só podem armazenar valores não negativos*/

	fmt.Println("\tInteiros assinados: \n")

	/*Inteiros assinados, declarados com uma das int
	palavras-chave, podem armazenar valores positivos e negativos:*/

	var num1 int = 500
	var num2 int = -4500

	fmt.Printf("Tipo: %T, valor: %v \n", num1, num1)
	fmt.Printf("Tipo: %T, valor: %v \n\n", num2, num2)

	fmt.Println("\tInteiros não assinados: \n")

	/*Inteiros sem sinal, declarados com uma das
	uint palavras-chave, só podem armazenar valores não negativos:*/

	var x unit = 500
	var y unit = 400

	fmt.Printf("Tipo: %T, valor: %v \n", x, x)
	fmt.Printf("Tipo: %T, valor: %v \n\n", y, y)

}
