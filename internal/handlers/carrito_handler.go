package handlers

import (
	"net/http"
	"strconv"

	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

type CarritoHandler struct {
	service *services.CarritoService
}

func NuevoCarritoHandler(service *services.CarritoService) *CarritoHandler {
	return &CarritoHandler{service: service}
}

// obtenerUserIDDelContexto extrae el user_id que puso el middleware.
func obtenerUserIDDelContexto(c *gin.Context) (uint, bool) {
	userID, existe := c.Get("user_id")
	if !existe {
		return 0, false
	}
	return userID.(uint), true
}

// VerCarrito maneja GET /carrito
func (h *CarritoHandler) VerCarrito(c *gin.Context) {
	// Obtener el ID del cliente desde el contexto (puesto por el middleware)
	clienteID, ok := obtenerUserIDDelContexto(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	// Pedir el carrito al service
	carrito, err := h.service.ObtenerCarrito(clienteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver el carrito con su total calculado
	respuesta := gin.H{
		"carrito": carrito,
		"total":   carrito.CalcularTotal(),
	}

	c.JSON(http.StatusOK, respuesta)
}

// AgregarItemRequest define los datos que esperamos en el body.
type AgregarItemRequest struct {
	ProductoID uint `json:"producto_id"`
	Cantidad   int  `json:"cantidad"`
}

// AgregarItem maneja POST /carrito/items
func (h *CarritoHandler) AgregarItem(c *gin.Context) {
	// 1. Obtener el cliente
	clienteID, ok := obtenerUserIDDelContexto(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	// 2. Parsear el body
	var req AgregarItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Llamar al service
	err := h.service.AgregarProducto(clienteID, req.ProductoID, req.Cantidad)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. Responder con éxito
	c.JSON(http.StatusCreated, gin.H{"mensaje": "producto agregado al carrito"})
}

// ModificarCantidadRequest define los datos del body.
type ModificarCantidadRequest struct {
	Cantidad int `json:"cantidad"`
}

// ModificarCantidad maneja PUT /carrito/items/:id
func (h *CarritoHandler) ModificarCantidad(c *gin.Context) {
	// Obtener el ID del item de la URL
	itemIDStr := c.Param("id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Parsear el body
	var req ModificarCantidadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al service
	err = h.service.ModificarCantidadItem(uint(itemID), req.Cantidad)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "cantidad actualizada"})
}

// EliminarItem maneja DELETE /carrito/items/:id
func (h *CarritoHandler) EliminarItem(c *gin.Context) {
	itemIDStr := c.Param("id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.service.EliminarItem(uint(itemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "item eliminado"})
}

// VaciarCarrito maneja DELETE /carrito
func (h *CarritoHandler) VaciarCarrito(c *gin.Context) {
	clienteID, ok := obtenerUserIDDelContexto(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	err := h.service.VaciarCarrito(clienteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "carrito vaciado"})
}