package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/ArquitecturaDeSistemas/taskmicroservice/database"
	model "github.com/ArquitecturaDeSistemas/taskmicroservice/dominio"
	"github.com/ArquitecturaDeSistemas/taskmicroservice/ports"

	"gorm.io/gorm"
)

/**
* Es un adaptador de salida

 */

type taskRepository struct {
	db             *database.DB
	activeSessions map[string]string
}

func NewTaskRepository(db *database.DB) ports.TaskRepository {
	return &taskRepository{
		db:             db,
		activeSessions: make(map[string]string),
	}
}

func ToJSON(obj interface{}) (string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

// Obtener Trabajo obtiene un trabajo por su ID.
func (ur *taskRepository) Usuario(id string) (*model.Usuario, error) {
	if id == "" {
		return nil, errors.New("El ID de usuario es requerido")
	}

	var tareaGORM model.TareaGORM
	//result := ur.db.GetConn().First(&usuarioGORM, id)
	result := ur.db.GetConn().First(&tareaGORM, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Printf("Error al obtener el usuario con ID %s: %v", id, result.Error)
		return nil, result.Error
	}

	return tareaGORM.ToGQL()
}

// Usuarios obtiene todos los usuarios de la base de datos.
func (ur *taskRepository) Usuarios() ([]*model.Tarea, error) {
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
		&model.UsuarioGORM{
			titulo:       input.Titulo,
			descripcion:  input.Descripcion,
			fechaInicio:  input.fechaInicio,
			fechaTermino: input.fechaTermino,
			userId:       input.userId,
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
	if input.fechaInicio != nil {
		tareaGORM.fechaInicio = *input.fechaInicio
	}
	if input.fechaTermino != nil {
		tareaGORM.fechaTermino = *input.fechaTermino
	}
	if input.userId != nil {
		tareaGORM.userId = *input.UserId
	}

	result = ur.db.GetConn().Save(&tareaGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Tarea actualizada: %v", tareaGORM)
	return tareaGORM.ToGQL()
}

// EliminarUsuario elimina un usuario de la base de datos por su ID.
func (ur *taskRepository) EliminarUsuario(id string) (*model.RespuestaEliminacion, error) {
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
		Mensaje: "Usuario eliminado con éxito",
	}
	return response, result.Error

}
