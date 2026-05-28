package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Conectar establece la conexión con SQLite y devuelve la instancia de GORM
func Conectar(rutaDB string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	log.Println("Conexión a base de datos establecida:", rutaDB)
	return db
}