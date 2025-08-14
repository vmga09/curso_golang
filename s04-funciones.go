package main

import (
	"fmt"
	"strings"
)

func main() {

	hello("Victor")

	nombre := upperName("Guillermo")
	fmt.Println(nombre)

	suma, mult := calculadora(5, 6)

	fmt.Printf("La suma es: %v y la mult es: %v\n", suma, mult)
	sum, mul := calc(7, 9)
	fmt.Printf("La suma %v , La mult %v\n", sum, mul)
}

// Funcion con argumentos
func hello(name string) {
	fmt.Printf("Esto es una llamada desde la funcion, mi nombre es: %s \n ", strings.ToUpper(name))
}

// Funcion con argumentos
func upperName(name string) string {
	return strings.ToUpper(name)
}

func calculadora(a, b int) (int, int) {
	suma := a + b
	mult := a * b
	return suma, mult
}

func calc(a, b int) (suma, mult int) {
	suma = a + b
	mult = a * b
	return
}
