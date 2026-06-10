package middleware

// Autenticación y validación de tokens del usuario

import (
	"net/http"
	"strings"

	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

// AuthRequired es un middleware que valida el token JWT en el header Authorization.
// Si el token es válido, deja pasar y agrega el user_id al contexto.
// Si no, devuelve 401 Unauthorized.
func AuthRequired(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Obtener el header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación requerido"})
			c.Abort()
			return
		}

		// 2. Formato esperado: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		// 3. Validar el token
		tokenString := parts[1]
		userID, err := authService.ValidarToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		// 4. Guardar el user_id en el contexto para que los handlers lo usen
		c.Set("user_id", userID)
		c.Next()
	}
}