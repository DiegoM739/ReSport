package main

import (
	"log"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/database"
	"github.com/DiegoM739/ReSport/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Programa inicial, el script es para poder demostrar mi trabajo en el Autónomo 2
/* Esta parte será desarrollada de forma completa de tal forma que
seed/main.go esté separado del sistema. Entonces el admin es independiente, aún falta crear
el panel funcional para el ADMIN. 
Donde donde voy a agregar endpoints/admin/login y 
admin/productos, es decir un middleware que valide el rol de superadmin 
y vistas HTML para gestionar productos y ver pedidos pendientes
*/
func main() {
	// 1. Cargar configuración
	cfg := config.Cargar()

	// 2. Conectar a la base de datos
	db := database.Conectar(cfg.DBPath)

	// 3. Verificar si el admin ya existe (para no duplicarlo)
	var contador int64
	db.Model(&models.Administrador{}).
		Where("email = ?", "admin@resport.com").
		Count(&contador)

	if contador > 0 {
		log.Println("El administrador ya existe en la base de datos.")
		return
	}

	// 4. Hashear la contraseña con bcrypt (igual que los clientes)
	passwordPlana := "admin123"
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordPlana), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error al hashear la contraseña: %v", err)
	}

	// 5. Crear el administrador con sus datos
	admin := models.Administrador{
		Persona: models.Persona{
			Nombre:   "Administrador Principal",
			Email:    "admin@resport.com",
			Telefono: "0999111222",
		},
		Rol:      "superadmin",
		Password: string(hash),
	}

	// 6. Guardar en la base de datos
	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Error al crear administrador: %v", err)
	}

	log.Println("✓ Administrador creado correctamente:")
	log.Println("  Email:    admin@resport.com")
	log.Println("  Password: admin123")
	log.Println("  Rol:      superadmin")
}