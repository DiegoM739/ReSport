package middleware

import (
	"net/http"
	"strings"

	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

// AdminRequired valida que la petición venga de un administrador.
func AdminRequired(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Leer el header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			c.Abort()
			return
		}

		// 2. Extraer el token "Bearer xxx"
		partes := strings.Split(authHeader, " ")
		if len(partes) != 2 || partes[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}
		token := partes[1]

		// 3. Validar token y obtener admin_id y rol
		adminID, rol, err := authService.ValidarTokenAdmin(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// 4. Verificar que sea un admin (no un cliente)
		if rol == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores"})
			c.Abort()
			return
		}

		// 5. Guardar info en el contexto
		c.Set("admin_id", adminID)
		c.Set("rol", rol)

		c.Next()
	}
}