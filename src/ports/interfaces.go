package ports

import (
	model "github.com/ArquitecturaDeSistemas/taskmicroservice/src/dominio"
)

// puerto de salida
type TaskRepository interface {
	CrearTarea(input model.CrearTareaInput) (*model.Tarea, error)
	Tarea(id string) (*model.Tarea, error)
	ActualizarTarea(id string, input *model.ActualizarTareaInput) (*model.Tarea, error)
	EliminarTarea(id string) (*model.RespuestaEliminacion, error)
	Tareas() ([]*model.Tarea, error)
}
