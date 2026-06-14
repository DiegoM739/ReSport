package handlers

import (
	"net/http"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

// AuthHandler maneja las rutas de registro y login.
type AuthHandler struct {
	clienteService *services.ClienteService
	authService    *services.AuthService
}

// NuevoAuthHandler crea una nueva instancia.
func NuevoAuthHandler(clienteService *services.ClienteService, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		clienteService: clienteService,
		authService:    authService,
	}
}

// RegistroRequest define los datos del body para registro.
type RegistroRequest struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Telefono string `json:"telefono"`
}

// Registro maneja POST /auth/registro
func (h *AuthHandler) Registro(c *gin.Context) {
	var req RegistroRequest

	// Parsear el body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear el cliente
	cliente := &models.Cliente{
		Persona: models.Persona{
			Nombre:   req.Nombre,
			Email:    req.Email,
			Telefono: req.Telefono,
		},
	}

	// Llamar al service
	if err := h.clienteService.Registrar(cliente, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Borrar la password de la respuesta
	cliente.Password = ""
	c.JSON(http.StatusCreated, cliente)
}

// LoginRequest define los datos del body para login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login maneja POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	// Parsear el body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login devuelve 3 valores: token, cliente, error
	token, cliente, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Responder con token + datos del cliente
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"user_id": cliente.ID,
		"email":   cliente.Email,
	})
}