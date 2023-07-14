package main

import (
	"fmt"
)

func main() {

	fmt.Println("Tipos de dados: \n")

	/*Go tem três tipos básicos de dados:

	bool : representa um valor booleano e é true ou false
	numérico : representa tipos inteiros, valores de ponto flutuante e tipos complexos
	string : representa um valor de string*/

	var a bool = true     //Booleano
	var b int = 5         //Inteiro
	var c float32 = 3.14  //Número flutuante
	var d string = "Olá!" //String

	fmt.Println("Booleano: ", a)
	fmt.Println("Inteiro: ", b)
	fmt.Println("Float: ", c)
	fmt.Print("String: ", d, "\n")

}
