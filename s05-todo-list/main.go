package main

import "fmt"

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

	fmt.Println("lista de Tareas")
	fmt.Println("------------------------------------------------------------------------------")

	for i, t := range lista.tareas {
		fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.descripcion, t.completado)

	}
	fmt.Println("------------------------------------------------------------------------------")

}
