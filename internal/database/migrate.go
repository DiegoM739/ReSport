package database

import (
	"log"

	"github.com/DiegoM739/ReSport/internal/models"
	"gorm.io/gorm"
)

// Migrar crea o actualiza las tablas en la base de datos según los modelos.
func Migrar(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Cliente{},
		&models.Administrador{},
		&models.Categoria{},
		&models.Producto{},
		&models.Direccion{},
		&models.Carrito{},
		&models.ItemCarrito{},
		&models.Pedido{},
		&models.ItemPedido{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}
	log.Println("Migración completada: todas las tablas creadas")
}