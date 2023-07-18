package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tOperadores aritméticos: \n")

	a := 5
	b := 10
	cont := 20
	decont := 30

	fmt.Println("Adição: \n")
	fmt.Println(a+b, "\n")

	fmt.Println("Subtração: \n")
	fmt.Println(b-a, "\n")

	fmt.Println("Multiplicação: \n")
	fmt.Println(a*b, "\n")

	fmt.Println("Divisão: \n")
	fmt.Println(b/a, "\n")

	fmt.Println("Módulo: \n")
	fmt.Println(b%a, "\n")

	fmt.Println("Incremento: \n")
	cont++ // -> cont = 20 | cont++ = 20 + 1 = 21
	fmt.Println(cont, "\n")

	fmt.Println("Decremento: \n")
	decont-- // -> decont = 30 | decont-- = 30 - 1 = 29
	fmt.Println(decont, "\n")

}
