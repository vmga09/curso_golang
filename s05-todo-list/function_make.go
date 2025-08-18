package main

import "fmt"

func main() {

	//Crear un slice del tipo any y longitud 5 y capacidad 10
	// Llenara con elementos nil (nulo)
	meses := make([]any, 5, 10)
	fmt.Println(meses)

	meses[0] = "Enero"
	meses[1] = 2

	fmt.Println(meses)

	sliceUno := []any{1, 2, "tres", 4, "cinco"}
	sliceDos := make([]any, 5)
	// Copia los elementos de sliceUno a sliceDos
	copy(sliceDos, sliceUno)
	fmt.Println(sliceUno)
	fmt.Println(sliceDos)
}
