package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tif aninhado: \n")

	/*
		Você pode ter ifinstruções dentro ifde instruções, isso é chamado de if aninhado.

		EXEMPLO:

		if condition1 {
			code to be executed if condition1 is true
		if condition2 {
		 	code to be executed if both condition1 and condition2 are true
			  }
			}
	*/

	num := 20

	if num >= 10 {
		fmt.Println("num é maior que 10.")
		if num > 15 {
			fmt.Println("num também é maior que 15.")
		}
	} else {
		fmt.Println("num é menor que 10.")
	}
}
