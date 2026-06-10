// Este handler maneja las rutas de perfil. Son protegidas (requieren token).

package handlers


/* 
clientes.Use(middleware.AuthRequired(authService)):
 aplica el middleware a TODO el grupo. 
 Cualquier ruta dentro de clientes requiere token válido
*/
import (
	"net/http"

	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	service *services.ClienteService
}

func NuevoClienteHandler(service *services.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: service}
}

// ObtenerPerfil maneja GET /clientes/perfil
func (h *ClienteHandler) ObtenerPerfil(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	cliente, err := h.service.ObtenerPerfil(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
		return
	}

	cliente.Password = ""
	c.JSON(http.StatusOK, cliente)
}

// ActualizarPerfil maneja PUT /clientes/perfil
func (h *ClienteHandler) ActualizarPerfil(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	cliente, err := h.service.ObtenerPerfil(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
		return
	}

	var datosActualizados struct {
		Nombre   string `json:"nombre"`
		Telefono string `json:"telefono"`
	}

	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cliente.Nombre = datosActualizados.Nombre
	cliente.Telefono = datosActualizados.Telefono

	if err := h.service.ActualizarPerfil(cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cliente.Password = ""
	c.JSON(http.StatusOK, cliente)
}