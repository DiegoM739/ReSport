package handlers

import (
	"net/http"
	"strconv"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

// ProductoHandler maneja las peticiones HTTP de productos.
type ProductoHandler struct {
	service *services.ProductoService
}

// NuevoProductoHandler crea una nueva instancia del handler.
func NuevoProductoHandler(service *services.ProductoService) *ProductoHandler {
	return &ProductoHandler{service: service}
}

// Crear maneja POST /productos
func (h *ProductoHandler) Crear(c *gin.Context) {
	var producto models.Producto

	// Convertir el JSON del body al struct Producto
	// Si JSON esta mal formado devolver error
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al service
	if err := h.service.CrearProducto(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Responder con el producto creado (ya con ID)
	c.JSON(http.StatusCreated, producto)
}

// Listar maneja GET /productos
func (h *ProductoHandler) Listar(c *gin.Context) {
	productos, err := h.service.ListarProductos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}

// Obtener maneja GET /productos/:id
func (h *ProductoHandler) Obtener(c *gin.Context) {
	id, err := obtenerIDDeParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	producto, err := h.service.ObtenerProducto(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}

// Actualizar maneja PUT /productos/:id
func (h *ProductoHandler) Actualizar(c *gin.Context) {
	id, err := obtenerIDDeParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var producto models.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	producto.ID = id // forzar el ID del param, no del body

	if err := h.service.ActualizarProducto(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, producto)
}

// Eliminar maneja DELETE /productos/:id
func (h *ProductoHandler) Eliminar(c *gin.Context) {
	id, err := obtenerIDDeParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.EliminarProducto(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado"})
}

// obtenerIDDeParam extrae el ID de la URL y lo convierte a uint.
func obtenerIDDeParam(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
