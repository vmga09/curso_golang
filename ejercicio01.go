package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Ejercicio 01 de calculo de triangulo rectangulo")
	var primero, segundo float64
	fmt.Print("Ingrese primer lado: ")
	fmt.Scanln(&primero)
	fmt.Print("Ingrese segundo lado: ")
	fmt.Scanln(&segundo)

	var hipotenusa = math.Sqrt(math.Pow(primero, 2) + math.Pow(segundo, 2))
	var area = (primero * segundo) / 2
	var perimetro = primero + segundo + hipotenusa

	fmt.Println("Resultados:")

	fmt.Printf("\tHipotenusa: %.2f\n\tArea: %.2f\n\tPerimetro: %.2f\n", hipotenusa, area, perimetro)

}
