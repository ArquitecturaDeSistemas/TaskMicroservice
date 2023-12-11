package dominio

import (
	"time"
)

type ActualizarTareaInput struct {
	Titulo       *string    `json:"titulo,omitempty"`
	Descripcion  *string    `json:"descripcion,omitempty"`
	FechaInicio  *time.Time `json:"fechaInicio,omitempty"`
	FechaTermino *time.Time `json:"fechaTermino,omitempty"`
	UserId       *string    `json:"userId,omitempty"`
}

type CrearTareaInput struct {
	Titulo       string    `json:"titulo"`
	Descripcion  string    `json:"descripcion"`
	FechaInicio  time.Time `json:"fechaInicio"`
	FechaTermino time.Time `json:"fechaTermino"`
	UserId       string    `json:"userId"`
}

type RespuestaEliminacion struct {
	Mensaje string `json:"mensaje"`
}

type Tarea struct {
	ID           string    `json:"id"`
	Titulo       string    `json:"titulo"`
	Descripcion  string    `json:"descripcion"`
	FechaInicio  time.Time `json:"fechaInicio"`
	FechaTermino time.Time `json:"fechaTermino"`
	UserId       string    `json:"userId"`
}

func (Tarea) IsEntity() {}
