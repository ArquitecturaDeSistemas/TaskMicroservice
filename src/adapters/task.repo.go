package adapters

import (
	"errors"
	"fmt"
	"log"

	"github.com/ArquitecturaDeSistemas/taskmicroservice/src/database"
	model "github.com/ArquitecturaDeSistemas/taskmicroservice/src/dominio"
	"github.com/ArquitecturaDeSistemas/taskmicroservice/src/ports"

	"gorm.io/gorm"
)

/**
* Es un adaptador de salida

 */

type taskRepository struct {
	db *database.DB
}

func NewTaskRepository(db *database.DB) ports.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// Obtener Trabajo obtiene un trabajo por su ID.
func (ur *taskRepository) Tarea(id string) (*model.Tarea, error) {
	if id == "" {
		return nil, errors.New("El ID de la tarea es requerido")
	}

	var tareaGORM model.TareaGORM
	//result := ur.db.GetConn().First(&usuarioGORM, id)
	result := ur.db.GetConn().First(&tareaGORM, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Printf("Error al obtener tarea con ID %s: %v", id, result.Error)
		return nil, result.Error
	}

	return tareaGORM.ToGQL()
}

// Usuarios obtiene todos los usuarios de la base de datos.
func (ur *taskRepository) Tareas() ([]*model.Tarea, error) {
	var tareasGORM []model.TareaGORM
	result := ur.db.GetConn().Find(&tareasGORM)

	if result.Error != nil {
		log.Printf("Error al obtener las tareas: %v", result.Error)
		return nil, result.Error
	}

	var tareas []*model.Tarea
	for _, usuarioGORM := range tareasGORM {
		tarea, _ := usuarioGORM.ToGQL()
		tareas = append(tareas, tarea)
	}
	return tareas, nil
}

func (ur *taskRepository) CrearTarea(input model.CrearTareaInput) (*model.Tarea, error) {

	tareaGORM :=
		&model.TareaGORM{
			Titulo:       input.Titulo,
			Descripcion:  input.Descripcion,
			FechaInicio:  input.FechaInicio,
			FechaTermino: input.FechaTermino,
			UserId:       input.UserId,
		}
	result := ur.db.GetConn().Create(&tareaGORM)
	if result.Error != nil {
		log.Printf("Error al crear la tarea: %v", result.Error)
		return nil, result.Error
	}

	response, err := tareaGORM.ToGQL()
	return response, err
}

func (ur *taskRepository) ActualizarTarea(id string, input *model.ActualizarTareaInput) (*model.Tarea, error) {
	var tareaGORM model.TareaGORM
	if id == "" {
		return nil, errors.New("El ID de usuario es requerido")
	}

	result := ur.db.GetConn().First(&tareaGORM, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Tarea con ID %s no encontrado", id)
		}
		return nil, result.Error
	}

	// Solo actualiza los campos proporcionados
	if input.Titulo != nil {
		tareaGORM.Titulo = *input.Titulo
	}
	if input.Descripcion != nil {
		tareaGORM.Descripcion = *input.Descripcion
	}
	if input.FechaInicio != nil {
		tareaGORM.FechaInicio = *input.FechaInicio
	}
	if input.FechaTermino != nil {
		tareaGORM.FechaTermino = *input.FechaTermino
	}
	if input.UserId != nil {
		tareaGORM.UserId = *input.UserId
	}

	result = ur.db.GetConn().Save(&tareaGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Tarea actualizada: %v", tareaGORM)
	return tareaGORM.ToGQL()
}

// EliminarUsuario elimina un usuario de la base de datos por su ID.
func (ur *taskRepository) EliminarTarea(id string) (*model.RespuestaEliminacion, error) {
	// Intenta buscar el usuario por su ID
	var tareaGORM model.TareaGORM
	result := ur.db.GetConn().First(&tareaGORM, id)

	if result.Error != nil {
		// Manejo de errores
		if result.Error == gorm.ErrRecordNotFound {
			// El tarea no se encontró en la base de datos
			response := &model.RespuestaEliminacion{
				Mensaje: "El tarea no existe",
			}
			return response, result.Error

		}
		log.Printf("Error al buscar el tarea con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al buscar el tarea",
		}
		return response, result.Error
	}

	// Elimina el tarea de la base de datos
	result = ur.db.GetConn().Delete(&tareaGORM, id)

	if result.Error != nil {
		log.Printf("Error al eliminar el tarea con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al eliminar tarea",
		}
		return response, result.Error
	}

	// Éxito al eliminar el tarea
	response := &model.RespuestaEliminacion{
		Mensaje: "Tarea eliminada con éxito",
	}
	return response, result.Error

}
