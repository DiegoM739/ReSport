package handlers

// Este handler maneja registro y login. Son rutas públicas
/* 
RegistroRequest: definimos una struct específica para el 
JSON de entrada. Para no usar el model CLIENTE 
Porque queremos controlar EXACTAMENTE qué campos acepta el endpoint. 
Si recibieras un Cliente directo, el cliente podría 
enviar campos como ID o CreatedAt y filtrar basura.
*/
import (
	"net/http"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	clienteService *services.ClienteService
	authService    *services.AuthService
}

func NuevoAuthHandler(clienteService *services.ClienteService, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		clienteService: clienteService,
		authService:    authService,
	}
}

// RegistroRequest define los datos que recibimos del cliente al registrar.
type RegistroRequest struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Telefono string `json:"telefono"`
}

// Registro maneja POST /auth/registro
func (h *AuthHandler) Registro(c *gin.Context) {
	var req RegistroRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cliente := &models.Cliente{
		Persona: models.Persona{
			Nombre:   req.Nombre,
			Email:    req.Email,
			Password: req.Password,
		},
		Telefono: req.Telefono,
	}

	if err := h.clienteService.Registrar(cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// IMPORTANTE: borrar la password de la respuesta (no devolvemos el hash)
	cliente.Password = ""
	c.JSON(http.StatusCreated, cliente)
}

// LoginRequest define los datos del login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login maneja POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}