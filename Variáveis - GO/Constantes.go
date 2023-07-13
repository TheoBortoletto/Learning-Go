package main

import (
	"fmt"
)

const PI = 3.14 // <- CONSTANTE
const A int = 1 // <- CONSTANTE COM TIPO
const B = 3     // <- CONSTANTE SEM TIPO

func main() {
	fmt.Println("\n\t Constantes: \n")

	/*Se uma variável deve ter valor fixo e que não pode ser alterado
	pode-se usar a palavra-chave 'const'

	Ela declara a variável como constante, o que significa imutável e somente leitura

	Exemplo: const x int = 9

	OBS: O valor de uam constante deve ser atrubuído ao declará-la*/

	fmt.Println(PI)

	fmt.Println("\n\t Tipos de constante: \n")

	/*Constantes com tipo e constantes sem tipo*/

	fmt.Println("Constante com tipo: ", A, "\n")
	fmt.Println("Constante sem tipo: ", B, "\n")
}
