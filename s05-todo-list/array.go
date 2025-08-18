package main

import "fmt"

func main() {

	//Definir un array de largo 5
	var a [5]int
	a[0] = 10
	a[1] = 20

	fmt.Println(a)

	//Definir array inicializado
	var b = [5]int{10, 20, 30, 40, 50}
	b[0] = 100
	fmt.Println(b)

	//Definir array con un largo indeterminado
	var c = [...]int{10, 20, 30}
	c[1] = 44
	fmt.Println(c)

	//	c[3] = 40 // Removed: index 3 is out of bounds for array 'c'

	//Definir un array bidimencional

	var matriz = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(matriz)

}
