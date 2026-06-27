package handlers

import (
	"net/http"

	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

// AdminHandler maneja las rutas de administradores.
type AdminHandler struct {
	authService *services.AuthService
}

// NuevoAdminHandler crea una nueva instancia.
func NuevoAdminHandler(authService *services.AuthService) *AdminHandler {
	return &AdminHandler{authService: authService}
}

// LoginAdminRequest define los datos del body.
type LoginAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login maneja POST /admin/login
func (h *AdminHandler) Login(c *gin.Context) {
	var req LoginAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al servicio (devuelve 3 valores: token, admin, error)
	token, admin, err := h.authService.LoginAdmin(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"admin_id": admin.ID,
		"email":    admin.Email,
		"rol":      admin.Rol,
		"mensaje":  "Bienvenido administrador",
	})
}