package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operadores booleanos en Go")

	fmt.Println("Verdadero:", true)
	fmt.Println("Falso:", false)

	// Operadores lógicos
	a := true
	b := false

	fmt.Println("a AND b:", a && b) // AND
	fmt.Println("a OR b:", a || b)  // OR
	fmt.Println("NOT a:", !a)       // NOT

	// Comparaciones
	x := 5
	y := 10

	fmt.Println("x == y:", x == y) // Igualdad
	fmt.Println("x != y:", x != y) // Desigualdad
	fmt.Println("x < y:", x < y)   // Menor que
	fmt.Println("x > y:", x > y)   // Mayor que
	fmt.Println("x <= y:", x <= y) // Menor o igual que
	fmt.Println("x >= y:", x >= y) // Mayor o igual que

	// Combinación de operadores lógicos
	resultado := (x < y) && (a || b)
	fmt.Println("Resultado de combinación:", resultado)

	// Comparación de cadenas
	// str1 := "Go"
	// str2 := "GoLang"

	const str1, str2 = "Go", "Golang"

	fmt.Println("str1 == str2:", str1 == str2) // FALSE
	fmt.Println("str1 != str2:", str1 != str2) // TRUE
	fmt.Println("str1 < str2:", str1 < str2)   // TRUE
	fmt.Println("str1 > str2:", str1 > str2)   // FALSE
	fmt.Println("str1 <= str2:", str1 <= str2) // TRUE
	fmt.Println("str1 >= str2:", str1 >= str2) // FALSE

	const X, Y = true, false
	fmt.Printf("El valor de X es: %t y el valor de Y es: %t\n", X, Y)
	const Z = X && Y
	println("X && Y :", Z)
	const W = X || Y
	println("X || Y :", W)

	//Condicion NOT
	const V = !X
	fmt.Printf("El valor de V es : %t\n", V)

	//Operaciones

	const A, B, C = 5, 10, 15

	const RESULT = ((A+B)*C)/(A*B) > C && A != B

	fmt.Printf("El valor de RESULT es %v\n ", RESULT)

}
