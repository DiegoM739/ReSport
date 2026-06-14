package repository

import (
	"github.com/DiegoM739/ReSport/internal/models"
	"gorm.io/gorm"
)

// ProductoRepository maneja el acceso a la tabla productos.
type ProductoRepository struct {
	db *gorm.DB
}

// NuevoProductoRepository crea una nueva instancia.
func NuevoProductoRepository(db *gorm.DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}

// Crear guarda un producto nuevo.
func (r *ProductoRepository) Crear(producto *models.Producto) error {
	return r.db.Create(producto).Error
}

// Listar devuelve todos los productos.
func (r *ProductoRepository) Listar() ([]models.Producto, error) {
	var productos []models.Producto
	err := r.db.Find(&productos).Error
	return productos, err
}

// BuscarPorID encuentra un producto por su ID.
func (r *ProductoRepository) BuscarPorID(id uint) (*models.Producto, error) {
	var producto models.Producto
	err := r.db.First(&producto, id).Error
	if err != nil {
		return nil, err
	}
	return &producto, nil
}

// Actualizar modifica un producto existente.
func (r *ProductoRepository) Actualizar(producto *models.Producto) error {
	return r.db.Save(producto).Error
}

// Eliminar borra un producto por su ID.
func (r *ProductoRepository) Eliminar(id uint) error {
	return r.db.Delete(&models.Producto{}, id).Error
}