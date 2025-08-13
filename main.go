package main

import (
	"fmt"
	"strconv"
)

// Variable global
var var01, var02, var03 = "uno", "dos", 5

//Definir constantes

const pi float32 = 3.14
const Plataforma = "linux"
const Verdad = true

func circulo() float32 {
	return pi * 3
}

//

func main() {
	fmt.Println("Hello, World!dd ")
	fmt.Println("hola")

	// Declaracion de una variable
	var firstName, lastName string
	var age int

	firstName = "Victor"
	lastName = "Gonzalez"
	age = 48

	// Otra fornma
	var (
		primerNombre, segundoNombre string
		newAge                      int
	)

	primerNombre = "Juan"
	segundoNombre = "Gonzalez"
	newAge = 40

	fmt.Println(firstName, lastName, age)
	fmt.Println(primerNombre, segundoNombre, newAge)

	// Otra forma de declarar variables
	myNombre := "victor"
	myApellido := "Gonzalez"
	myEdad := 49

	fmt.Println(myNombre, myApellido, myEdad)

	fmt.Println(var01, var02, var03)

	fmt.Println(circulo())

	//Otros tipos de datos
	// Tipos de datos byte
	var abyte byte = '@'
	println(abyte)

	// Convesion de tipo
	var integer16 int16 = 50
	var integer32 int32 = 100
	// Cambia de base 32 a 16
	fmt.Println(integer16 + int16(integer32))

	f := "100"
	// strconv.Atoi convierte de string a entero
	j, _ := strconv.Atoi(f)
	println(j)
	// Pasar de string a float
	g := "3.14"
	h, _ := strconv.ParseFloat(g, 32)
	println(h)

	// Entero a string
	t := 56
	u := strconv.Itoa(t)
	println(u)
	//contatenar
	v := u + u
	println(v)

	// Paquete fmt
	// Format
	fmt.Printf("Hola mi nombre es %s %s y tengo %d años \n", myNombre, myApellido, myEdad)
	// Concatenación en otra variable
	greeting := fmt.Sprintf("Hola mi nombre es %s %s y tengo %d años \n", myNombre, myApellido, myEdad)
	println(greeting)
	// Ingresar valores (similar input en python)

	//var primero string
	//var segundo int

	// fmt.Print("Ingrese primer valor: ")
	// fmt.Scanln(&primero)
	// fmt.Print("Ingrese segundo valor: ")
	// fmt.Scanln(&segundo)

	// fmt.Printf("Valor primero: %s , valor segundo: %d \n", primero, segundo)

	diez := 10
	once := 4

	fmt.Println((float32(diez) / float32(once)))
	fmt.Println(diez % once)
}
