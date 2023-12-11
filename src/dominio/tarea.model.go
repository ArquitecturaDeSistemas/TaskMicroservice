package dominio

import (
	"strconv"
	"time"
)

// TareaGORM es el modelo de tarea para GORM de Tarea
type TareaGORM struct {
	ID           uint      `gorm:"primaryKey:autoIncrement" json:"id"`
	Titulo       string    `gorm:"type:varchar(255);not null"`
	Descripcion  string    `gorm:"type:varchar(255);not null"`
	fechaInicio  time.Time `gorm:"type:varchar(255);not null;unique"`
	fechaTermino time.Time `gorm:"type:varchar(255);not null"`
	userId       string    `gorm:"type:varchar(255);not null"`
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
		fechaInicio:  tareaGORM.fechaInicio,
		fechaTermino: tareaGORM.fechaTermino,
	}, nil
}
