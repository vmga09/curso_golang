package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	//Una forma de instancear una estrutura
	var p Persona
	p.nombre = "Victor"
	p.edad = 48
	p.correo = "vmgonzalez@outlook.com"

	a := Persona{nombre: "Juan", edad: 44, correo: "jc.gonzalez@gmail.com"}
	fmt.Println(a)
}
