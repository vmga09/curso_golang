package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slice

	//var a[]int
	diasSemana := []string{"Lunes", "Martes", "Miercoles"}
	fmt.Println(cap(diasSemana))
	diasSemana = append(diasSemana, "Jueves")
	fmt.Println(diasSemana)

	fmt.Println(diasSemana[2:]) // Miercoles Jueves
	fmt.Println(diasSemana[:2]) // Lunes Martes

	//Crear un slice desde otro
	diass := diasSemana[:3]
	fmt.Println(diass) // Lunes, Martes y Miercoles

	//Cantidad de elementos, capacidad
	fmt.Println(len(diass)) // 3
	fmt.Println(cap(diass)) // 6

	// Agregar un elemento al final del slice
	diasSemana = append(diasSemana, "Viernes", "Sabado", "Domingo")
	fmt.Println(len(diasSemana))
	fmt.Println(cap(diasSemana))
	fmt.Println(diasSemana)

	//Eliminar un elemento
	diasSemana = append(diasSemana[0:2], diasSemana[3:]...)
	fmt.Println(diasSemana)

	//Eliminar el primer elemento de un slice
	diasSemana = diasSemana[1:]
	fmt.Println(diasSemana)
	//Eliminar el ultimo elemento de un slice
	diasSemana = diasSemana[:len(diasSemana)-1]
	fmt.Println(diasSemana)
	// Agregar un elemento al inicio de un slice
	diasSemana = append([]string{"Lunes"}, diasSemana...)
	fmt.Println(diasSemana)

	//Buscar un elemento dentro de un slice
	fmt.Println(slices.Contains(diasSemana, "x"))

	index, status := findSlice(diasSemana, "Domingo")

	fmt.Println(index, status)

	// for _,x := range diasSemana {
	// 	if x == "Lunes" {
	// 		fmt.Println("Si se encuentra")
	// 		break
	// 	}
	// }

}

func findSlice(slice []string, valor string) (int, bool) {
	for index, element := range slice {
		if element == valor {
			return index, true
		}
	}
	return -1, false
}
