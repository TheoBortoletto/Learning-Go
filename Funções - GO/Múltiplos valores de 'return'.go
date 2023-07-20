package main

import (
	"fmt"
)

/*Aqui, myFunction() retorna um inteiro ( result ) e uma string ( txt1 ):
EX1:*/

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

/*Aqui, armazenamos os dois valores de retorno em duas variáveis ​​( a , b ):
EX2:*/

func myFunction2(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {

	fmt.Println("\n\t\tMúltiplos valores de 'return':\n")

	fmt.Println("\nExemplo 1:\n")
	fmt.Println(myFunction(5, "Hello"))

	fmt.Println("\nExemplo 2:\n")
	a, b := myFunction2(5, "Hello")
	fmt.Println(a, b)

}
