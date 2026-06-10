package repository

import (
	"github.com/DiegoM739/ReSport/internal/models"
	"gorm.io/gorm"
)

// ProductoRepository maneja todo el acceso a la tabla productos.
type ProductoRepository struct {
	db *gorm.DB
}

// NuevoProductoRepository crea una nueva instancia del repositorio.
// Recibe la conexión a la base de datos por parámetro (esto se llama "inyección de dependencias").
func NuevoProductoRepository(db *gorm.DB) *ProductoRepository {
	return &ProductoRepository{db: db} // diccionario clave valor punterl al producto respository 
}

// Crear guarda un nuevo producto en la base de datos.
func (r *ProductoRepository) Crear(producto *models.Producto) error {
	return r.db.Create(producto).Error
}

// ListarTodos devuelve todos los productos junto con su categoría.
func (r *ProductoRepository) ListarTodos() ([]models.Producto, error) {
	var productos []models.Producto
	err := r.db.Preload("Categoria").Find(&productos).Error
	return productos, err
}

// BuscarPorID devuelve un producto específico por su ID.
func (r *ProductoRepository) BuscarPorID(id uint) (*models.Producto, error) {
	var producto models.Producto
	err := r.db.Preload("Categoria").First(&producto, id).Error
	if err != nil {
		return nil, err
	}
	return &producto, nil
}

// Actualizar modifica un producto existente.
func (r *ProductoRepository) Actualizar(producto *models.Producto) error {
	return r.db.Save(producto).Error
}

// Eliminar borra un producto (para evitar problemas en las BD se relaliza
// soft delete: es decir no se borra realmente, se marca como borrado).
func (r *ProductoRepository) Eliminar(id uint) error {
	return r.db.Delete(&models.Producto{}, id).Error
}
