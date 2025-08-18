package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) saludo() {
	fmt.Println(" Hola, mi nombre es", p.nombre)
}

func main() {
	// Punteros
	var x int = 10
	var p *int = &x
	fmt.Println(x) // Variable
	fmt.Println(p) // Poscicion de memoria

	editar(&x)
	fmt.Println(x)

	usuario_01 := Persona{"Victor", 38, "vmga.cl@gmail.com"}

	usuario_01.saludo()

}

func editar(x *int) {
	*x = 20
}
