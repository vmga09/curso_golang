package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	fmt.Printf("Bienvenido al juego\n")
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10
	min, max := 0, 100

	for intentos < maxIntentos {

		fmt.Printf("Por favor ingrese un numero ( Te quedan %v intentos) : ", maxIntentos-intentos)
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Printf("Felicitaciones, has acertado el numero es: %v\n", numIngresado)
			jugarNuevamente()
			return
		} else if numIngresado > numAleatorio {

			fmt.Printf("Lo siento no has acertado, El número esta entre %v y %v\n", numIngresado, min)
			max = numIngresado
		} else {

			fmt.Printf("Lo siento no has acertado, El número esta entre %v y %v\n", numIngresado, max)
			min = numIngresado
		}
		intentos++
	}
	fmt.Printf("Lo siento, no acertaste el número era: %v", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var opcion string
	fmt.Println("¿Desea jugar nuevamente? (s/n)")
	fmt.Scanln(&opcion)

	switch strings.ToLower(opcion) {
	case "s":
		jugar()
	case "n":
		fmt.Printf("Gracias por jugar, hasta pronto\n")
	default:
		fmt.Println("Ingrese una opción válida ")
		jugarNuevamente()
	}

}
