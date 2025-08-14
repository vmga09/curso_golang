package main

import "fmt"

func main() {

	//Bucle infinito

	// i := 0
	// for {
	// 	fmt.Println("HOLA FOR :", i)
	// 	i++
	// }

	for i := 0; i < 10; i++ {
		fmt.Println("HOLA :", i)
	}

	texto := "Esto es un texto"

	for i, char := range texto {
		fmt.Printf("Índice: %d, Carácter: %c, Valor: %d\n", i, char, char)
		//fmt.Println(texto[i])
	}

	for _, char := range texto {
		fmt.Printf("Carácter: %c\n", char)
	}

	runes := []rune(texto)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("Carácter %d: %c\n", i, runes[i])
	}

	fmt.Println(runes)

	// Obtener carácter en posición específica
	char := runes[2]        // Obtiene 't'
	charStr := string(char) // Convertir a string

	fmt.Printf("Carácter: %c\n", char)       // Imprime: t
	fmt.Printf("Como string: %s\n", charStr) // Imprime: t

	fmt.Println(char)    // Imprime: numero
	fmt.Println(charStr) // Imprime : t

	for i := 0; i < 10; i++ {

		if i > 5 {
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println("HOLA :", i)
	}

}
