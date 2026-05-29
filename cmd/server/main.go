package main

import (
	"log"
	"net/http"

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

	// 3. Migrar las tablas (CREA TODAS LAS TABLAS AUTOMÁTICAMENTE)
	database.Migrar(db)

	_ = db // todavía no la usamos en los handlers

	// 4. Crear router de Gin
	router := gin.Default()

	// 5. Ruta de bienvenida
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Bienvenido a ReSport",
			"version": "0.1.0",
			"estado":  "funcionando",
		})
	})

	// 6. Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 7. Levantar servidor
	log.Println("Servidor iniciando en puerto:", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error al levantar servidor: %v", err)
	}
}