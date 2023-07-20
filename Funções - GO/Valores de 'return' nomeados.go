package main

import (
	"fmt"
)

func myFunction(x int, y int) (result int) {
	result = x + y
	return //*= Também pode ser escrito 'return result'
}

func main() {

	fmt.Println("\n\t\tValores de 'return' nomeados:\n")

	/*Em Go, você pode nomear os valores de retorno de uma função.

	Aqui, nomeamos o valor de retorno como 'result' (de tipo int)
	e retornamos o valor com um retorno simples (significa que
	usamos a instrução 'return' sem especificar o nome da variável):*/

	fmt.Println(myFunction(1, 2))

}
