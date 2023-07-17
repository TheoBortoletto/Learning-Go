package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tEncontrar o comprimento de uma  matriz: \n")

	carros := [4]string{"Volvo", "BMW", "Ford", "Fiat"}
	numero := [...]int{1, 2, 3, 4, 5}

	fmt.Println("Comprimentoda matriz 'carros':", len(carros))
	fmt.Println("Comprimento da matriz 'numeros':", len(numero), "\n")

}
