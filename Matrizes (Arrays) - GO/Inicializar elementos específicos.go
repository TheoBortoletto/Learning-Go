package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tInicializar apenas elementos específicos: \n")

	mtx1 := [5]int{1: 35, 4: 40} /* Apenas a posição 1 e 4 possuem elementos/valores */
	fmt.Println(mtx1, "\n")

	mtx2 := [3]string{0: "Olá", 2: "Mundo!!!"} /* Apenas a posição 0 e 2 possuem elementos/valores */
	fmt.Println(mtx2, "\n")

}
