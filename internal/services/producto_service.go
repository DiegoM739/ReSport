package services

import (
	"errors"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/repository"
)

// ProductoService contiene la lógica de negocio para productos.
type ProductoService struct {
	repo repository.IProductoRepository
}

// NuevoProductoService crea una nueva instancia del service.
func NuevoProductoService(repo repository.IProductoRepository) *ProductoService {
	return &ProductoService{repo: repo}
}

// CrearProducto valida y guarda un producto nuevo.
func (s *ProductoService) CrearProducto(producto *models.Producto) error {
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
	return s.repo.Crear(producto)
}

// ListarProductos devuelve todos los productos.
func (s *ProductoService) ListarProductos() ([]models.Producto, error) {
	return s.repo.Listar()
}

// ObtenerProducto busca un producto por su ID.
func (s *ProductoService) ObtenerProducto(id uint) (*models.Producto, error) {
	producto, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("producto no encontrado")
	}
	return producto, nil
}

// ActualizarProducto modifica un producto existente.
func (s *ProductoService) ActualizarProducto(producto *models.Producto) error {
	if producto.ID == 0 {
		return errors.New("ID del producto es obligatorio")
	}
	if producto.Precio <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	return s.repo.Actualizar(producto)
}

// EliminarProducto borra un producto por su ID.
func (s *ProductoService) EliminarProducto(id uint) error {
	return s.repo.Eliminar(id)
}