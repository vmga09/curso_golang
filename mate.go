package main

import (
	"fmt"
	"math"
)

// Funcion matematica Pow

func Pow(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Funcion para calcular la raiz cuadrada de dos numeros
func sqrt() {
	const (
		x = 16
		y = 25
	)
	c := math.Sqrt(x) + math.Sqrt(y)
	fmt.Printf("%.1f\n", c)
}

func main() {

	a := 10
	a++
	fmt.Println("Valor de a:", a)

	b := 20
	b += 5
	fmt.Println("Valor de b:", b)

	c := 30
	c -= 10
	fmt.Println("Valor de c:", c)

	d := 40
	d *= 2
	fmt.Println("Valor de d:", d)

	e := 50
	e /= 5
	fmt.Println("Valor de e:", e)

	//Contantes de la libreria math
	fmt.Println("Valor de pi:", math.Pi)
	fmt.Println("Valor de e:", math.E)

	fmt.Println("2^4 =", Pow(2, 4))

	sqrt()
}

// Removed duplicate main function to resolve redeclaration error
