package models

import "gorm.io/gorm"

// Estados posibles de un pedido (constantes para evitar errores de tipeo).
const (
	EstadoPendiente = "pendiente"
	EstadoPagado    = "pagado"
	EstadoEnviado   = "enviado"
	EstadoEntregado = "entregado"
	EstadoCancelado = "cancelado"
)

// Pedido representa una orden de compra confirmada por un cliente.
type Pedido struct {
	gorm.Model
	ClienteID   uint
	Cliente     Cliente
	Items       []ItemPedido
	Total       float64
	Estado      string `gorm:"default:pendiente"`
	MetodoPago  string // "tarjeta" o "transferencia"
	DireccionID uint
}

// ItemPedido representa un producto comprado dentro de un pedido.
// Se separa del ItemCarrito porque el precio se "congela" al momento de comprar.
type ItemPedido struct {
	gorm.Model
	PedidoID     uint
	ProductoID   uint
	Producto     Producto
	Cantidad     int
	PrecioUnidad float64 // precio al momento de la compra (congelado)
	Subtotal     float64
}

// CambiarEstado actualiza el estado del pedido validando que sea válido.
// ENCAPSULACIÓN: solo se puede cambiar a un estado permitido.
func (p *Pedido) CambiarEstado(nuevoEstado string) bool {
	estadosValidos := []string{
		EstadoPendiente, EstadoPagado, EstadoEnviado,
		EstadoEntregado, EstadoCancelado,
	}
	for _, estado := range estadosValidos {
		if estado == nuevoEstado {
			p.Estado = nuevoEstado
			return true
		}
	}
	return false // estado inválido, no se cambia
}
