/* Tipos de Variáveis em 'GO':

- int
- float32
- string
- bool (Armazena valores com dois estados: verdadeiro ou falso)
*/

/* Em 'Go', existem duas maneiras de declarar uma variável:

1°: Palavra-chave "var"

Use a palavra-chave "var", em seguida o nome da variável

Exemplo:

var num int = 2

2°: Com o sinal ":=" :

Use o sinal ":=", em seguida o valor da variável

Exemplo:

num := 2

NOTA: Neste caso, o tipo da variável é inferido a
partir do valor (significa que o compilador decide o
tipo da variável com base no valor).

NOTA: Não é possível declarar uma variável usando :=, sem atribuir um valor a ela.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n\tDeclaração de variável com valor inicial:\n")

	/* Se o valor de uma variável for conhecido desde o início,
	você pode declarar a variável e atribuir um valor a ela em uma linha: */

	var student1 string = "John" // Tipo da váriavel é String
	var student2 = "Jane"
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// Observação: os tipos de variáveis ​​de student2e xsão inferidos de seus valores.

	fmt.Println("\n\tDeclaração de variável sem valor inicial: \n")

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	/* As variáveis acima não recebem valor inicial,
	portanto a = " "; b = 0; c = false.*/

	fmt.Println("\n\tAtribuição de valor após a declaração: \n")

	/**/
}
