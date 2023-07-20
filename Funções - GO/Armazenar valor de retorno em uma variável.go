package main

import (
	"fmt"
)

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func main() {

	fmt.Println("\n\t\tArmazenar valor de retorno em uma variável: \n")

	total := myFunction(1, 2)
	fmt.Println(total)
}
