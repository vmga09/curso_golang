package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Funcion con un error personalizado
func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero") // Tambien se puede usar fmt.Errorf
	}
	return dividendo / divisor, nil
}

func main() {

	str := "123"

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("El error es :", err)
		return
	}
	fmt.Println(num)

	resultado, error := divide(10, 3)
	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	fmt.Println(resultado)

}
