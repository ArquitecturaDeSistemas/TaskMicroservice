package database

import (
	"log"

	model "github.com/ArquitecturaDeSistemas/taskmicroservice/src/dominio"
	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.TareaGORM{})

	log.Println("Migraciones completadas")
}
