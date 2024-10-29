package main

import (
	"fmt"
)

type Tarea struct {
	ID          int
	Nombre      string
	Descripción string
	Prioridad   int
	Completada  bool
}

func (t *Tarea) MarcarComoCompletada() {
	t.Completada = true
}

func (t *Tarea) MostrarTarea() {
	fmt.Printf("ID: %d\nNombre: %s\nDescripción: %s\nPrioridad: %d\nCompletada: %t\n", t.ID, t.Nombre, t.Descripción, t.Prioridad, t.Completada)
}

type GestorTareas struct {
	Tareas         []Tarea
	ContadorTareas int
}

func (gt *GestorTareas) AgregarTarea(nombre, descripción string, prioridad int) {
	gt.ContadorTareas++
	tarea := Tarea{ID: gt.ContadorTareas, Nombre: nombre, Descripción: descripción, Prioridad: prioridad, Completada: false}
	gt.Tareas = append(gt.Tareas, tarea)
}

func (gt *GestorTareas) EliminarTarea(id int) {
	for i, tarea := range gt.Tareas {
		if tarea.ID == id {
			gt.Tareas = append(gt.Tareas[:i], gt.Tareas[i+1:]...)
			break
		}
	}
}

func (gt *GestorTareas) ObtenerTareasPendientes() []Tarea {
	var tareasPendientes []Tarea
	for _, tarea := range gt.Tareas {
		if !tarea.Completada {
			tareasPendientes = append(tareasPendientes, tarea)
		}
	}
	return tareasPendientes
}

func main() {
	gestor := GestorTareas{}
	gestor.AgregarTarea("Tarea 1", "Descripción 1", 1)
	gestor.AgregarTarea("Tarea 2", "Descripción 2", 2)

	gestor.Tareas[0].MarcarComoCompletada()

	for _, tarea := range gestor.ObtenerTareasPendientes() {
		tarea.MostrarTarea()
	}
}
