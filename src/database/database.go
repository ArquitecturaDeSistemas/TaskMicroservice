package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

func Connect() *DB {
	time.Sleep(5 * time.Second)
	// dsn := "root:mysql@tcp(172.16.238.10:3306)/db_users?parseTime=true"
	dsn := "root:mysql@tcp(localhost:3306)/db_tareas_arq?parseTime=true"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos MySQL: %v", err)
	}
	return &DB{conn}
}

// Función para obtener la conexión de GORM
func (db *DB) GetConn() *gorm.DB {
	return db.conn
}
