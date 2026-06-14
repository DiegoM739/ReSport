package repository

import (
	"github.com/DiegoM739/ReSport/internal/models"
	"gorm.io/gorm"
)

// IProductoRepository define el contrato que cualquier repositorio
// de productos debe cumplir. Esto permite que el service no dependa
// de una implementación concreta, sino de esta abstracción.
type IProductoRepository interface {
	Crear(producto *models.Producto) error
	Listar() ([]models.Producto, error)
	BuscarPorID(id uint) (*models.Producto, error)
	Actualizar(producto *models.Producto) error
	Eliminar(id uint) error
}

// IClienteRepository define el contrato del repositorio de clientes.
type IClienteRepository interface {
	Crear(cliente *models.Cliente) error
	BuscarPorID(id uint) (*models.Cliente, error)
	BuscarPorEmail(email string) (*models.Cliente, error)
	Actualizar(cliente *models.Cliente) error
}

// ICarritoRepository define el contrato del repositorio de carritos.
type ICarritoRepository interface {
	ObtenerOCrearPorCliente(clienteID uint) (*models.Carrito, error)
	AgregarItem(item *models.ItemCarrito) error
	ActualizarItem(item *models.ItemCarrito) error
	EliminarItem(itemID uint) error
	BuscarItem(itemID uint) (*models.ItemCarrito, error)
	VaciarCarrito(carritoID uint) error
}

// IPedidoRepository define el contrato del repositorio de pedidos.
type IPedidoRepository interface {
	Crear(pedido *models.Pedido) error
	CrearEnTransaccion(tx *gorm.DB, pedido *models.Pedido) error
	ListarPorCliente(clienteID uint) ([]models.Pedido, error)
	BuscarPorID(id uint) (*models.Pedido, error)
	Actualizar(pedido *models.Pedido) error
}