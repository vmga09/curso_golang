package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("La hora es de mañana")
	} else if hora == 12 {
		fmt.Println("Es mediodía")
	} else {
		fmt.Println(" Es de tarde")
	}

	// Otra forma :

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("La hora es de mañana")
	} else if t.Hour() == 12 {
		fmt.Println("Es mediodía")
	} else {
		fmt.Println(" Es de tarde")
	}

	// Otra forma :

	if t := time.Now().Hour(); t < 12 {
		fmt.Println("La hora es de mañana")
	} else if t == 12 {
		fmt.Println("Es mediodía")
	} else {
		fmt.Println(" Es de tarde")
	}

	//os := runtime.GOOS
	os2 := runtime.GOARCH
	//fmt.Println(os)
	fmt.Println(os2)

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run on Windows")
	case "linux":
		fmt.Println("Go run on Linux")
	case "darwin":
		fmt.Println("Go run on macOS")
	default:
		fmt.Printf("Corre en otro os: %s", os)
	}

	switch t := time.Now().Hour(); {
	case t < 12:
		fmt.Println("Mañana\n")
	case t == 12:
		fmt.Printf("Medio Día\n")
	case t < 17:
		fmt.Printf("Tarde\n")
	default:
		fmt.Printf("Noche\n")

	}

}
