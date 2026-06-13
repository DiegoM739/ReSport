package handlers

import (
	"net/http"
	"strconv"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	service *services.PedidoService
}

func NuevoPedidoHandler(service *services.PedidoService) *PedidoHandler {
	return &PedidoHandler{service: service}
}

// CrearPedidoRequest define los datos del body.
type CrearPedidoRequest struct {
	MetodoPago string `json:"metodo_pago"` // "tarjeta" o "transferencia"

	// Datos opcionales según el método de pago
	DatosTarjeta  *DatosTarjetaInput  `json:"datos_tarjeta,omitempty"`
	DatosTransfer *DatosTransferInput `json:"datos_transferencia,omitempty"`
}

type DatosTarjetaInput struct {
	NumeroTarjeta   string `json:"numero_tarjeta"`
	Titular         string `json:"titular"`
	FechaExpiracion string `json:"fecha_expiracion"`
	CVV             string `json:"cvv"`
}

type DatosTransferInput struct {
	Banco        string `json:"banco"`
	NumeroCuenta string `json:"numero_cuenta"`
	Comprobante  string `json:"comprobante"`
}

// CrearPedido maneja POST /pedidos
// Convierte el carrito en pedido y procesa el pago.
func (h *PedidoHandler) CrearPedido(c *gin.Context) {
	// === Paso 1: Obtener el cliente ===
	clienteID, ok := obtenerUserIDDelContexto(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	// === Paso 2: Parsear el body ===
	var req CrearPedidoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// === Paso 3: Construir el método de pago según el tipo ===
	// Aquí es donde se ve el POLIMORFISMO:
	// Dependiendo del tipo, creamos UN objeto u OTRO, pero ambos
	// implementan la interface MetodoPago.
	var metodoPago models.MetodoPago

	switch req.MetodoPago {
	case "tarjeta":
		// Verificar que vengan los datos de tarjeta
		if req.DatosTarjeta == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "faltan datos de tarjeta"})
			return
		}
		// Crear objeto PagoTarjeta
		metodoPago = models.PagoTarjeta{
			NumeroTarjeta:   req.DatosTarjeta.NumeroTarjeta,
			Titular:         req.DatosTarjeta.Titular,
			FechaExpiracion: req.DatosTarjeta.FechaExpiracion,
			CVV:             req.DatosTarjeta.CVV,
		}

	case "transferencia":
		if req.DatosTransfer == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "faltan datos de transferencia"})
			return
		}
		metodoPago = models.PagoTransferencia{
			Banco:        req.DatosTransfer.Banco,
			NumeroCuenta: req.DatosTransfer.NumeroCuenta,
			Comprobante:  req.DatosTransfer.Comprobante,
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "método de pago inválido"})
		return
	}

	// === Paso 4: Llamar al service ===
	// El service no sabe ni le importa si es PagoTarjeta o PagoTransferencia.
	// Solo sabe que metodoPago cumple la interface MetodoPago.
	pedido, err := h.service.CrearPedidoDesdeCarrito(clienteID, metodoPago, req.MetodoPago)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// === Paso 5: Responder con el pedido creado ===
	c.JSON(http.StatusCreated, pedido)
}

// ListarPedidos maneja GET /pedidos
func (h *PedidoHandler) ListarPedidos(c *gin.Context) {
	clienteID, ok := obtenerUserIDDelContexto(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	pedidos, err := h.service.ListarPedidosCliente(clienteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pedidos)
}

// ObtenerPedido maneja GET /pedidos/:id
func (h *PedidoHandler) ObtenerPedido(c *gin.Context) {
	clienteID, ok := obtenerUserIDDelContexto(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	pedidoIDStr := c.Param("id")
	pedidoID, err := strconv.ParseUint(pedidoIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	pedido, err := h.service.ObtenerPedido(uint(pedidoID), clienteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pedido)
}
