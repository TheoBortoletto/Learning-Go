package main

import (
	"fmt"
)

/*Se (por algum motivo) não quisermos usar algum dos
valores retornados, podemos adicionar um sublinhado ( _ ),
para omitir esse valor.

Aqui, queremos omitir o primeiro valor retornado
( result- que é armazenado na variável a ):

EX1:*/

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {

	fmt.Println("\n\t\tOmitir valores de uma função: \n")

	fmt.Println("Exemplo 1:")
	_, b := myFunction(5, "Hello")
	fmt.Println(b)
}
