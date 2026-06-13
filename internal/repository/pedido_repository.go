package repository

import (
	"github.com/DiegoM739/ReSport/internal/models"
	"gorm.io/gorm"
)

// PedidoRepository maneja al acceso a la tabla de pedidos.

type PedidoRepository struct {
	db *gorm.DB
}

// NuevoPedidoRepository crea una nueva instancia
// Creamos el constructor para crear y entregar un repository listo para usar
// (db *gorm.DB) recirbe un parametro que es la conexión de la BD
// Este puntero lo pongo en todo el programa para que sea una misma
func NuevoPedidoRepository(db *gorm.DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

// Crear guarda un nuevo pedido en la base de datos.
// Se usa el puntero porque necesitamos la conexion a db
// Crear es el metodo que crea un nuevo pedido en base de datos
// *models.Pedido es el parametro que recibe
func (r *PedidoRepository) Crear(pedido *models.Pedido) error {
	return r.db.Create(pedido).Error
}

// CrearEnTransaccion guarda un pedido dentro de una transaccion
// la transaccion es crear pedido, reducir stock producto a-b-c, vaciar carrito
// La transaccion se controla desde el service
func (r *PedidoRepository) CrearEnTransaccion(tx *gorm.DB, pedido *models.Pedido) error {
	return tx.Create(pedido).Error
}

// ListarPorCliente devuelve todos los pedidos de un cliente
func (r *PedidoRepository) ListarPorCliente(clienteID uint) ([]models.Pedido, error) {
	var pedidos []models.Pedido
	err := r.db.
		Where("cliente_id = ?", clienteID).
		Preload("Items.Producto").
		Order("created_at DESC").
		Find(&pedidos).Error
	return pedidos, err
}

/*
   Order("created_at DESC"): ordena los pedidos del más
   reciente al más viejo.
   Es lo que el usuario espera ver al consultar su historial.
*/

// BuscarPorID devuelve un pedido especifico
func (r *PedidoRepository) BuscarPorID(id uint) (*models.Pedido, error) {
	var pedido models.Pedido
	err := r.db.
		Preload("Items.Producto").
		First(&pedido, id).Error
	if err != nil {
		return nil, err
	}
	return &pedido, nil
}

// Actualizar modifica un pedido existente
func (r *PedidoRepository) Actualizar(pedido *models.Pedido) error {
	return r.db.Save(pedido).Error
}
