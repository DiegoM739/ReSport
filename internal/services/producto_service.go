package services

import (
	"errors"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/repository"
)

// ProductoService contiene la lógica de negocio relacionada con productos.
type ProductoService struct {
	repo *repository.ProductoRepository
}

// NuevoProductoService crea una nueva instancia del service.
func NuevoProductoService(repo *repository.ProductoRepository) *ProductoService {
	return &ProductoService{repo: repo}
}

// CrearProducto valida los datos y crea un producto.
func (s *ProductoService) CrearProducto(producto *models.Producto) error {
	// === Reglas de negocio ===
	if producto.Nombre == "" {
		return errors.New("el nombre del producto es obligatorio")
	}
	if producto.Precio <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	if producto.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	if producto.Tipo != "fisico" && producto.Tipo != "digital" {
		return errors.New("el tipo debe ser 'fisico' o 'digital'")
	}

	// Si todo es válido, mandamos al repositorio
	return s.repo.Crear(producto)
}

// ListarProductos devuelve todos los productos.
func (s *ProductoService) ListarProductos() ([]models.Producto, error) {
	return s.repo.ListarTodos()
}

// ObtenerProducto busca un producto por ID.
func (s *ProductoService) ObtenerProducto(id uint) (*models.Producto, error) {
	return s.repo.BuscarPorID(id)
}

// ActualizarProducto valida y modifica un producto existente.
func (s *ProductoService) ActualizarProducto(producto *models.Producto) error {
	if producto.Nombre == "" {
		return errors.New("el nombre del producto es obligatorio")
	}
	if producto.Precio <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	if producto.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return s.repo.Actualizar(producto)
}

// EliminarProducto borra un producto del sistema.
func (s *ProductoService) EliminarProducto(id uint) error {
	return s.repo.Eliminar(id)
}
