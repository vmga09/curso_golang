package main

import "fmt"

func main() {

	colors := map[string]any{
		"rojo":   "#FF0000",
		"numero": 2,
		"lista":  []int{1, 2, 3, 4, 5},
	}

	fmt.Println(colors)

	//Agregar un nuevo elemento al map

	colors["boolean"] = true
	fmt.Println(colors)

	//Verificar la existencia de un valor

	_, ok := colors["planeta"]

	if ok {
		fmt.Println("La clave existe")
	} else {
		fmt.Println("No existe la clave")
	}

	//Otra forma de verificar

	if n, ok := colors["rojo"]; ok {
		fmt.Println(n)
	} else {
		fmt.Println("No exsite el valor")
	}

	//Eliminar un elemento del maps
	delete(colors, "lista")

	fmt.Println(colors)

	//Iterar elementos
	for clave, valor := range colors {
		fmt.Printf("El valor de colors[%s] es \t: %v\n", clave, valor)
	}

}
