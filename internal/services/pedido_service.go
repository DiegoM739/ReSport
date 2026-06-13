package services

import (
	"errors"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/repository"
	"gorm.io/gorm"
)

// PedidoService contiene la lógica de negocio de pedidos.
type PedidoService struct {
	pedidoRepo   *repository.PedidoRepository
	carritoRepo  *repository.CarritoRepository
	productoRepo *repository.ProductoRepository
	db           *gorm.DB
}

// NuevoPedidoService crea el service.
// Recibe los 3 repositorios + la DB (para manejar transacciones).
func NuevoPedidoService(
	pedidoRepo *repository.PedidoRepository,
	carritoRepo *repository.CarritoRepository,
	productoRepo *repository.ProductoRepository,
	db *gorm.DB,
) *PedidoService {
	return &PedidoService{
		pedidoRepo:   pedidoRepo,
		carritoRepo:  carritoRepo,
		productoRepo: productoRepo,
		db:           db,
	}
}

// CrearPedidoDesdeCarrito convierte el carrito del cliente en un pedido formal.
// Procesa el pago, reduce el stock y vacía el carrito.
// Recibe el método de pago como interface (POLIMORFISMO).
func (s *PedidoService) CrearPedidoDesdeCarrito(
	clienteID uint,
	metodoPago models.MetodoPago,
	tipoPago string,
) (*models.Pedido, error) {

	// === Paso 1: Obtener el carrito del cliente ===
	carrito, err := s.carritoRepo.ObtenerOCrearPorCliente(clienteID)
	if err != nil {
		return nil, errors.New("error al obtener el carrito")
	}

	// === Paso 2: Verificar que el carrito tenga items ===
	if len(carrito.Items) == 0 {
		return nil, errors.New("el carrito está vacío")
	}

	// === Paso 3: Validar stock de TODOS los productos ===
	// Antes de hacer cambios, verificamos que todo esté disponible.
	for _, item := range carrito.Items {
		if !item.Producto.VerificarStock(item.Cantidad) {
			mensaje := "stock insuficiente para el producto: " + item.Producto.Nombre
			return nil, errors.New(mensaje)
		}
	}

	// === Paso 4: Calcular el total del pedido ===
	totalPedido := carrito.CalcularTotal()

	// === Paso 5: Procesar el pago (POLIMORFISMO en acción) ===
	// metodoPago es una INTERFACE. Puede ser PagoTarjeta o PagoTransferencia.
	// El service NO sabe cuál es. Solo llama al método Procesar().
	pagoExitoso, mensajePago := metodoPago.Procesar(totalPedido)
	if !pagoExitoso {
		return nil, errors.New("el pago falló: " + mensajePago)
	}

	// === Paso 6: Crear el pedido en una TRANSACCIÓN ===
	// Una transacción garantiza que todo se hace o nada se hace.
	// Si algo falla en el medio, GORM revierte automáticamente.
	var pedidoCreado *models.Pedido

	errorTransaccion := s.db.Transaction(func(tx *gorm.DB) error {
		// 6.1: Crear los items del pedido (copiando del carrito)
		var itemsPedido []models.ItemPedido
		for _, itemCarrito := range carrito.Items {
			nuevoItemPedido := models.ItemPedido{
				ProductoID:   itemCarrito.ProductoID,
				Cantidad:     itemCarrito.Cantidad,
				PrecioUnidad: itemCarrito.Producto.Precio,
				Subtotal:     itemCarrito.Subtotal,
			}
			itemsPedido = append(itemsPedido, nuevoItemPedido)
		}

		// 6.2: Crear el pedido
		nuevoPedido := &models.Pedido{
			ClienteID:  clienteID,
			Items:      itemsPedido,
			Total:      totalPedido,
			Estado:     models.EstadoPagado, // ya está pagado
			MetodoPago: tipoPago,
		}

		errorCrear := tx.Create(nuevoPedido).Error
		if errorCrear != nil {
			return errorCrear
		}

		// 6.3: Reducir el stock de cada producto
		for _, itemCarrito := range carrito.Items {
			producto := itemCarrito.Producto
			producto.Stock = producto.Stock - itemCarrito.Cantidad

			errorActualizar := tx.Save(&producto).Error
			if errorActualizar != nil {
				return errorActualizar
			}
		}

		// 6.4: Vaciar el carrito
		errorVaciar := tx.Where("carrito_id = ?", carrito.ID).
			Delete(&models.ItemCarrito{}).Error
		if errorVaciar != nil {
			return errorVaciar
		}

		// Guardamos el pedido creado para devolverlo
		pedidoCreado = nuevoPedido

		// Si llegamos aquí, todo bien. La transacción se confirma.
		return nil
	})

	if errorTransaccion != nil {
		return nil, errors.New("error al crear el pedido: " + errorTransaccion.Error())
	}

	return pedidoCreado, nil
}

// ListarPedidosCliente devuelve todos los pedidos del cliente.
func (s *PedidoService) ListarPedidosCliente(clienteID uint) ([]models.Pedido, error) {
	return s.pedidoRepo.ListarPorCliente(clienteID)
}

// ObtenerPedido busca un pedido específico, validando que sea del cliente.
func (s *PedidoService) ObtenerPedido(pedidoID uint, clienteID uint) (*models.Pedido, error) {
	pedido, err := s.pedidoRepo.BuscarPorID(pedidoID)
	if err != nil {
		return nil, errors.New("pedido no encontrado")
	}

	// Verificar que el pedido pertenezca al cliente que pregunta
	if pedido.ClienteID != clienteID {
		return nil, errors.New("no tienes permiso para ver este pedido")
	}

	return pedido, nil
}
