package main

import (
	"fmt"
	"log"

	"github.com/vmga09/greetings"
)

func main() {
	log.SetPrefix("greetings -> ")
	log.SetFlags(0)
	message, error := greetings.Hello("LALALALAL")
	if error != nil {
		//log.Fatal(error)
		fmt.Println(error)
		return
	}
	fmt.Println(message)
}
