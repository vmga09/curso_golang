package main

import (
	"log"
	"os"
)

func main() {
	// Uso del paquete LOG
	// Especifica un archivo donde guardar el log
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("Este es un registro")
	log.Print("Este es otro registro")
	str := "POWER"
	log.Printf("Esto es %s", str)
	log.SetPrefix("main() -> ")
	log.Print("Aqui otro registro")
	log.Fatal("Esto es el fin")

}
