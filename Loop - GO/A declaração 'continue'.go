package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tA declaração 'continue': \n")

	/*A instrução continue é usada para pular uma ou mais iterações no loop.
	Em seguida, continua com a próxima iteração no loop.*/

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // Se i == 3, pula para a proxima iteração
		}
		fmt.Println(i)
	}
}
