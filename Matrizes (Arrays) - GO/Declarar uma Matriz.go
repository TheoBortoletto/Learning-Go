package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tMatrizes 'GO'\n")

	/*Arrays são usados ​​para armazenar vários
	valores do mesmo tipo em uma única variável,
	em vez de declarar variáveis ​​separadas para
	cada valor.*/

	fmt.Println("\tDeclarar uma matriz:\n")

	/*Existem duas maneiras de declarar uma matriz: usando 'var' ou com o sinal ':='*/

	var mtx1 = [5]int{0, 1, 2, 3, 4}   // -> tamanho da matriz definido
	var mtx2 = [...]int{1, 2, 3, 4, 5} // -> tamanho da matriz indefinido

	/*  OU  */

	mtx3 := [4]int{2, 7, 3, 9}
	mtx4 := [...]int{2, 4, 9, 1, 98, 33}

	fmt.Println(mtx1)
	fmt.Println(mtx2)
	fmt.Println(mtx3)
	fmt.Println(mtx4)

	fmt.Println("\n\tMatriz de strings: \n")

	var name = [3]string{"Alan", "Beatriz", "Théo"}
	/*-> 3 é o tamanho da matriz, mas a posição 0 é a primeira*/

	fmt.Println("Matriz de string: ", name)

}
