package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tA declaração 'break':")

	/*A instrução break é usada para interromper/encerrar a execução do loop.*/

	for i := 0; i < 5; i++ {
		if i == 3 {
			break // -> ENCERRA O CÓDIGO SE FOR IGUAL A 3
		}
		fmt.Println(i)
	}
}
