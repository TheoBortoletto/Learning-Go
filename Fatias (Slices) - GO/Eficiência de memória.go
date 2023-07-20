package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tEficiência de Memória: \n")

	/*
		Ao usar fatias, Go carrega todos os elementos subjacentes na memória.

		Se a matriz for grande e você precisar apenas de alguns elementos, é melhor
		copiar esses elementos usando a função 'copy'().

		A função 'copy()' cria uma nova matriz subjacente apenas com os
		elementos necessários para a fatia. Isso reduzirá a memória usada para o programa.

		SINTAXE:

		- copy(dest, src)

		A função 'copy()' recebe duas fatias dest e src e copia os
		dados de src para dest . Ele retorna o número de elementos copiados.
	*/

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// FATIA ORIGINAL:
	fmt.Printf("numeros = %v\n", numeros)
	fmt.Printf("Comprimento = %d\n", len(numeros))
	fmt.Printf("Capacidade = %d\n", cap(numeros))
	fmt.Println()

	// CRIANDO UMA CÓPIA COM OS NÚMEROS NECESSÁRIOS:
	numNecessarios := numeros[:len(numeros)-10]
	numerosCopia := make([]int, len(numNecessarios))
	copy(numerosCopia, numNecessarios)

	fmt.Printf("numerosCopia = %v\n", numerosCopia)
	fmt.Printf("Comprimento = %d\n", len(numerosCopia))
	fmt.Printf("Capacidade = %d\n", cap(numerosCopia))

}
