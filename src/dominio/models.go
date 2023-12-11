package dominio

type ActualizarTareaInput struct {
	Titulo       *string `json:"titulo,omitempty"`
	Descripcion  *string `json:"descripcion,omitempty"`
	FechaInicio  *string `json:"fechaInicio,omitempty"`
	FechaTermino *string `json:"fechaTermino,omitempty"`
	UserId       *string `json:"userId,omitempty"`
}

type CrearTareaInput struct {
	Titulo       string `json:"titulo"`
	Descripcion  string `json:"descripcion"`
	FechaInicio  string `json:"fechaInicio"`
	FechaTermino string `json:"fechaTermino"`
	UserId       string `json:"userId"`
}

type RespuestaEliminacion struct {
	Mensaje string `json:"mensaje"`
}

type Tarea struct {
	ID           string `json:"id"`
	Titulo       string `json:"titulo"`
	Descripcion  string `json:"descripcion"`
	FechaInicio  string `json:"fechaInicio"`
	FechaTermino string `json:"fechaTermino"`
	UserId       string `json:"userId"`
}

func (Tarea) IsEntity() {}
