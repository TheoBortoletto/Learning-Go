package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tLoop 'for': \n")

	/*SINTAXE:

	for instrução1; instrução2; instrução3 {
	   	código a ser executado em cada iteração
	}

	instrução1 Inicializa o valor do contador de loop.

	instrução2 Avaliado para cada iteração do loop.
	Se for avaliado como TRUE, o loop continua.
	Se for avaliado como FALSE, o loop termina.

	instrução3 Aumenta o valor do contador de loop.
	*/

	fmt.Println("Exemplo 1: \n")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*Exemplo 1 explicado

	i:=0; - Inicializa o contador de loop (i) e define o valor inicial para 0
	i < 5; - Continua o loop enquanto i for menor que 5
	i++ - Aumente o valor do contador do loop em 1 para cada iteração
	*/

	fmt.Println("\nExemmplo 2:\n")

	for x := 0; x <= 100; x += 10 {
		fmt.Println(x)
	}

	/*Exemplo 2 explicado
	i:=0; - Inicializa o contador de loop (i) e define o valor inicial para 0
	i <= 100; - Continua o loop enquanto i for menor ou igual a 100
	i+=10 - Aumenta o valor do contador do loop em 10 para cada iteração
	*/
}
