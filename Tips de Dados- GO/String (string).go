package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tTipo de dados flutuantes:\n")

	/*O stringtipo de dados é usado para armazenar uma
	sequência de caracteres (texto). Os valores de string
	devem estar entre aspas duplas:*/

	var txt1 string = "Olá!!"
	var txt2 string
	txt3 := "Mundo 1"

	fmt.Printf("Tipo: %T, valor: %v \n", txt1, txt1)
	fmt.Printf("Tipo: %T, valor: %v \n", txt2, txt2)
	fmt.Printf("Tipo: %T, valor: %v \n\n", txt3, txt3)

}
