 


package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Cargar configuración
	cfg := config.Cargar()
	log.Println("Configuración cargada. Entorno:", cfg.Env)

	// 2. Conectar a base de datos
	db := database.Conectar(cfg.DBPath)
	_ = db // Por ahora no la usamos, pero ya está lista

	// 3. Crear router de Gin
	router := gin.Default()

	// 4. Definir ruta de prueba
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje":   "Bienvenido a ReSport",
			"version":   "0.1.0",
			"estado":    "funcionando",
		})
	})

	// 5. Health check (estándar en backends profesionales)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 6. Levantar servidor
	log.Println("Servidor iniciando en puerto:", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error al levantar servidor: %v", err)
	}

	fmt.Println("Bienvenido")
}

