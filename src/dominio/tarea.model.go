package dominio

import (
	"strconv"
)

// TareaGORM es el modelo de tarea para GORM de Tarea
type TareaGORM struct {
	ID           uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Titulo       string `gorm:"type:varchar(255);not null"`
	Descripcion  string `gorm:"type:varchar(255);not null"`
	FechaInicio  string `gorm:"type:varchar(255);not null"`
	FechaTermino string `gorm:"type:varchar(255);not null"`
	UserId       string `gorm:"type:varchar(255);not null"`
}

// TableName especifica el nombre de la tabla para UsuarioGORM
func (TareaGORM) TableName() string {
	return "tareas"
}

func (tareaGORM *TareaGORM) ToGQL() (*Tarea, error) {

	return &Tarea{
		ID:           strconv.Itoa(int(tareaGORM.ID)),
		Titulo:       tareaGORM.Titulo,
		Descripcion:  tareaGORM.Descripcion,
		FechaInicio:  tareaGORM.FechaInicio,
		FechaTermino: tareaGORM.FechaTermino,
		UserId:       tareaGORM.UserId,
	}, nil
}
