package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tA declaração 'else if': \n")

	/*

		Use a else ifinstrução para especificar uma nova condição se a primeira condição for false.

			SINTAXE:

		if condition1 {
			code to be executed if condition1 is true
		} else if condition2 {
			code to be executed if condition1 is false and condition2 is true
		} else {
			code to be executed if condition1 and condition2 are both false
		}

	*/

	hora := 22

	if hora < 10 {
		fmt.Println("Bom dia!!! \n")
	} else if hora < 20 {
		fmt.Println("Tenha um bom dia!!! \n")
	} else {
		fmt.Println("Boa Tarde!!!")
	}

}
