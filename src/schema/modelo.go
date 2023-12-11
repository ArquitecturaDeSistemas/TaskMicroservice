package schema

import "time"

type Notification struct {
	ID           uint      `gorm:"primaryKey:autoIncrement" json:"id"`
	Titulo       string    `gorm:"type:varchar(255);not null"`
	descripcion  string    `gorm:"type:varchar(255);not null"`
	fechaInicio  time.Time `gorm:"not null" json:"fechaInicio"`
	fechaTermino time.Time `gorm:"not null" json:"fecha_termino"`
	userId       string    `gorm:"type:varchar(255);not null"`
}

type Notifications []Notification
