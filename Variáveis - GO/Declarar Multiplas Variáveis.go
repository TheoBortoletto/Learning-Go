package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n\t Declaração de variáveis múltiplas, na mesma linha: \n")

	// Em 'Go', é possível declarar mútipllas variáveis na mesma linhas:

	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("\nSe o tipo de variável não for especificada, você pode declarar diferentes tipos de variáveis na mesma linha:\n")

	var w, x = 6, "Olá"
	y, z := 7, "Mundo"

	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("\nDeclaração de variável em um bloco: \n")

	var (
		i int
		j int    = 1
		k string = "hello"
	)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
