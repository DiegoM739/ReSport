package main

/*
Script inicial para crear el administrador.
Se ejecuta UNA SOLA VEZ con: go run cmd/seed/main.go
*/

import (
	"log"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/database"
	"github.com/DiegoM739/ReSport/internal/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	ADMIN_EMAIL    = "admin@resport.com"
	ADMIN_PASSWORD = "admin123"
	ADMIN_NOMBRE   = "Administrador Principal"
	ADMIN_TELEFONO = "0999111222"
)

func main() {
	cfg := config.Cargar()
	log.Println("Configuración cargada para seed.")

	db := database.Conectar(cfg.DBPath)
	log.Println("Conexión a base de datos establecida.")

	// Verificar si el admin ya existe
	var contador int64
	db.Model(&models.Administrador{}).
		Where("email = ?", ADMIN_EMAIL).
		Count(&contador)

	if contador > 0 {
		log.Println("El administrador ya existe en la base de datos.")
		return
	}

	// Hashear la contraseña con bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(ADMIN_PASSWORD), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error al hashear la contraseña: %v", err)
	}

	// Crear el administrador
	admin := models.Administrador{
		Persona: models.Persona{
			Nombre:   ADMIN_NOMBRE,
			Email:    ADMIN_EMAIL,
			Password: string(hash),
			Telefono: ADMIN_TELEFONO,
		},
		Rol: "superadmin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Error al crear el administrador: %v", err)
	}

	log.Println("==============================================")
	log.Println(" Administrador creado correctamente:")
	log.Println("  Email:    " + ADMIN_EMAIL)
	log.Println("  Password: " + ADMIN_PASSWORD)
	log.Println("  Rol:      superadmin")
	log.Println("==============================================")
}