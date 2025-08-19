package main

import "fmt"

func divide(dividendo, divisor int) {

	// Esto de defer + recover es similar al try catch
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	// validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

//	func validateZero(divisor int) {
//		if divisor == 0 {
//			panic("No se puede dividir por 0")
//		}
//	}
func main() {

	// Ver PANIC

	divide(100, 20)
	divide(100, 40)
	divide(300, 0)
	divide(40, 2)

}
