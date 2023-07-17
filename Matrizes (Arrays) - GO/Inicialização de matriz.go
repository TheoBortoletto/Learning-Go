package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tInicialização de matriz: \n")

	mtx1 := [5]int{}                   // Matriz não inicializada
	mtx2 := [5]int{10, 20}             // Matriz parcialmente inicializada
	mtx3 := [5]int{10, 20, 30, 40, 50} // Matrix completa

	fmt.Println("Matriz não inicializada:", mtx1)
	fmt.Println("Matriz parcialmente inicializada:", mtx2)
	fmt.Println("Matrix completa:", mtx3)
}
