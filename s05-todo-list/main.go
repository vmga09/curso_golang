package main

import (
	"bufio"
	"fmt"
	"os"
)

// Estructura de una tarea
type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

// Estructura de una lista de tareas
type ListaTareas struct {
	tareas []Tarea
}

// Método para agregar tareas
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Método para marcar como completado una tarea
func (l *ListaTareas) completarTarea(index int) {
	l.tareas[index].completado = true
}

// Método para editar una tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

//Método para eliminar tarea

func (l *ListaTareas) eliminar(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	// Instanciar lista de tareas

	lista := ListaTareas{}
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar una tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agrgada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el índice de la tarea a completar:")
			fmt.Scanln(&index)
			lista.completarTarea(index)
			fmt.Println("Tarea Marcada como completada")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el índice de la tarea para actualizar:")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea Actualizada")
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea a Eliminar:")
			fmt.Scanln(&index)
			lista.eliminar(index)
		case 5:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opción inválida.")

		}

		fmt.Println("lista de Tareas")
		fmt.Println("------------------------------------------------------------------------------")

		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.descripcion, t.completado)

		}
		fmt.Println("------------------------------------------------------------------------------")
	}

}
