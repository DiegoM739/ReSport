package main

import (
	"log"
	"net/http"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/database"
	"github.com/DiegoM739/ReSport/internal/handlers"
	"github.com/DiegoM739/ReSport/internal/repository"
	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Cargar configuración
	cfg := config.Cargar()
	log.Println("Configuración cargada. Entorno:", cfg.Env)

	// 2. Conectar a base de datos y migrar
	db := database.Conectar(cfg.DBPath)
	database.Migrar(db)

	// 3. Inicializar capas (inyección de dependencias)
	productoRepo := repository.NuevoProductoRepository(db)
	productoService := services.NuevoProductoService(productoRepo)
	productoHandler := handlers.NuevoProductoHandler(productoService)

	// 4. Crear router de Gin
	router := gin.Default()

	// 5. Rutas básicas
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Bienvenido a ReSport"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 6. Rutas de productos (agrupadas)
	productos := router.Group("/productos")
	{
		productos.POST("", productoHandler.Crear)
		productos.GET("", productoHandler.Listar)
		productos.GET("/:id", productoHandler.Obtener)
		productos.PUT("/:id", productoHandler.Actualizar)
		productos.DELETE("/:id", productoHandler.Eliminar)
	}

	// 7. Levantar servidor
	log.Println("Servidor iniciando en puerto:", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error al levantar servidor: %v", err)
	}
}

