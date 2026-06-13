package repository

import (
	"github.com/DiegoM739/ReSport/internal/models"
	"gorm.io/gorm"
)

// CarritRepository maneja el acceso a la tabla carritos
type CarritoRepository struct {
	db *gorm.DB
}

// NuevoCarritoRepository crea una nueva instancia del repositorio
func NuevoCarritoRepository(db *gorm.DB) *CarritoRepository {
	return &CarritoRepository{db: db}
}

// ObtenerOCrearPorCliente busca el carrito del cliente
// Si no existe, lo crea. Esto es get or create

func (r *CarritoRepository) ObtenerOCrearPorCliente(clienteID uint) (*models.Carrito, error) {
	// 1. Crear variable vacía para guardar el resultado
	var carrito models.Carrito

	// 2. Buscar el carrito del cliente en la base
	resultado := r.db.
		Where("cliente_id = ?", clienteID).
		Preload("Items.Producto").
		First(&carrito)

	// 3. Si no se encontró, crear uno nuevo
	if resultado.Error == gorm.ErrRecordNotFound {
		carrito = models.Carrito{
			ClienteID: clienteID,
		}
		errorCrear := r.db.Create(&carrito).Error
		if errorCrear != nil {
			return nil, errorCrear
		}
		return &carrito, nil
	}

	// 4. Si hubo otro error, devolverlo
	if resultado.Error != nil {
		return nil, resultado.Error
	}

	// 5. Devolver el carrito encontrado
	return &carrito, nil
}

// AgregarItem guarda un nuevo item en la base.
func (r *CarritoRepository) AgregarItem(item *models.ItemCarrito) error {
	return r.db.Create(item).Error
}

// ActualizarItem modifica un item existente.
func (r *CarritoRepository) ActualizarItem(item *models.ItemCarrito) error {
	return r.db.Save(item).Error
}

// EliminarItem borra un item del carrito por su ID.
func (r *CarritoRepository) EliminarItem(itemID uint) error {
	return r.db.Delete(&models.ItemCarrito{}, itemID).Error
}

// BuscarItem encuentra un item específico por su ID.
func (r *CarritoRepository) BuscarItem(itemID uint) (*models.ItemCarrito, error) {
	var item models.ItemCarrito
	err := r.db.Preload("Producto").First(&item, itemID).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// VaciarCarrito borra todos los items de un carrito.
func (r *CarritoRepository) VaciarCarrito(carritoID uint) error {
	return r.db.Where("carrito_id = ?", carritoID).Delete(&models.ItemCarrito{}).Error
}
