package schema

type Tarea struct {
	ID           uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Titulo       string `gorm:"type:varchar(255);not null"`
	Descripcion  string `gorm:"type:varchar(255);not null"`
	FechaInicio  string `gorm:"not null" json:"fechaInicio"`
	FechaTermino string `gorm:"not null" json:"fecha_termino"`
	UserId       string `gorm:"type:varchar(255);not null"`
}

type Notifications []Tarea
