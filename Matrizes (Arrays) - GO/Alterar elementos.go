package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tAlterar elementos de uma matriz:\n")

	valores := [3]int{19.00, 12.00, 20.00}
	var produtos = [3]string{"Leite", "Açúcar", "Presunto"}

	fmt.Println("MATRIZES: \n")
	fmt.Println("Valores:", valores)
	fmt.Println("Produtos:", produtos, "\n")

	fmt.Println("Valor dos produtos: \n")
	fmt.Println(produtos[0], ":", valores[0])
	fmt.Println(produtos[1], ":", valores[1])
	fmt.Println(produtos[2], ":", valores[2], "\n")

	fmt.Println("Valores novos: \n")

	valores[0] = 10 /*Alterando os valores*/
	valores[1] = 15 /*Alterando os valores*/
	valores[2] = 17 /*Alterando os valores*/

	fmt.Println(produtos[0], ":", valores[0])
	fmt.Println(produtos[1], ":", valores[1])
	fmt.Println(produtos[2], ":", valores[2], "\n")

}
