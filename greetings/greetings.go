package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre Vacío")
	}
	message := fmt.Sprintf(randoFormat(), name)
	return message, nil

}

func randoFormat() string {
	formats := []string{
		"Hola, %v. ¿Cómo estás?",
		"Hola, %v. ¿Cómo te va?",
		"Hola, %v. ¿Cómo andas?",
		"Hola, %v. ¿Qué tal estás?",
		"Hola, %v. ¿Cómo te encuentras?",
		"Hola, %v. ¿Cómo va todo?",
		"Hola, %v. ¿Qué tal tu día?",
		"Hola, %v. ¿Cómo ha estado todo?",
		"Hola, %v. ¿Cómo te sientes?",
		"Hola, %v. ¿Qué tal va todo?"}

	return formats[rand.Intn(len(formats))]

}
