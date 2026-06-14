package main

/*
Script de siembra (seed) para crear el administrador inicial.
Se ejecuta UNA SOLA VEZ con: go run cmd/seed/main.go
*/

import (
	"log"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/database"
	"github.com/DiegoM739/ReSport/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// === 1. Cargar configuración del .env ===
	cfg := config.Cargar()
	log.Println("Configuración cargada para seed.")

	// === 2. Conectar a la base de datos ===
	db := database.Conectar(cfg.DBPath)
	log.Println("Conexión a base de datos establecida.")

	// === 3. Verificar si el admin ya existe (idempotencia) ===
	// Evita crear duplicados si se ejecuta el seed varias veces.
	var contador int64
	db.Model(&models.Administrador{}).
		Where("email = ?", "admin@resport.com").
		Count(&contador)

	if contador > 0 {
		log.Println("El administrador ya existe en la base de datos. No se crea de nuevo.")
		return
	}

	// === 4. Hashear la contraseña con bcrypt ===
	// Usamos el mismo algoritmo que para los clientes regulares.
	passwordPlana := "admin123"
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordPlana), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error al hashear la contraseña: %v", err)
	}

	// === 5. Crear el administrador con sus datos ===
	// Persona se embebe (herencia) y aporta los campos comunes:
	// Nombre, Email, Password, Telefono.
	admin := models.Administrador{
		Persona: models.Persona{
			Nombre:   "Administrador Principal",
			Email:    "admin@resport.com",
			Password: string(hash),
			Telefono: "0999111222",
		},
		Rol: "superadmin",
	}

	// === 6. Guardar en la base de datos ===
	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Error al crear el administrador: %v", err)
	}

	// === 7. Confirmación visual ===
	log.Println("==============================================")
	log.Println("✓ Administrador creado correctamente:")
	log.Println("  Email:    admin@gmail.com")
	log.Println("  Password: admin123")
	log.Println("  Rol:      superadmin")
	log.Println("==============================================")
}