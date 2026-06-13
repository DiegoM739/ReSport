package main

import (
	"log"
	"net/http"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/database"
	"github.com/DiegoM739/ReSport/internal/handlers"
	"github.com/DiegoM739/ReSport/internal/middleware"
	"github.com/DiegoM739/ReSport/internal/repository"
	"github.com/DiegoM739/ReSport/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// === 1. Configuración y base de datos ===
	cfg := config.Cargar()
	log.Println("Configuración cargada. Entorno:", cfg.Env)

	db := database.Conectar(cfg.DBPath)
	database.Migrar(db)

	// === 2. Inicializar repositorios ===
	productoRepo := repository.NuevoProductoRepository(db)
	clienteRepo := repository.NuevoClienteRepository(db)
	carritoRepo := repository.NuevoCarritoRepository(db)
	pedidoRepo := repository.NuevoPedidoRepository(db)

	// === 3. Inicializar servicios ===
	productoService := services.NuevoProductoService(productoRepo)
	clienteService := services.NuevoClienteService(clienteRepo)
	authService := services.NuevoAuthService(clienteRepo, cfg)
	carritoService := services.NuevoCarritoService(carritoRepo, productoRepo)
	pedidoService := services.NuevoPedidoService(pedidoRepo, carritoRepo, productoRepo, db)

	// === 4. Inicializar handlers ===
	productoHandler := handlers.NuevoProductoHandler(productoService)
	authHandler := handlers.NuevoAuthHandler(clienteService, authService)
	clienteHandler := handlers.NuevoClienteHandler(clienteService)
	carritoHandler := handlers.NuevoCarritoHandler(carritoService)
	pedidoHandler := handlers.NuevoPedidoHandler(pedidoService)

	// === 5. Router ===
	router := gin.Default()
	router.SetTrustedProxies(nil) // silenciar el warning de proxies

	// Rutas básicas
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Bienvenido a ReSport"})
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Productos (públicos)
	productos := router.Group("/productos")
	{
		productos.POST("", productoHandler.Crear)
		productos.GET("", productoHandler.Listar)
		productos.GET("/:id", productoHandler.Obtener)
		productos.PUT("/:id", productoHandler.Actualizar)
		productos.DELETE("/:id", productoHandler.Eliminar)
	}

	// Autenticación (pública)
	auth := router.Group("/auth")
	{
		auth.POST("/registro", authHandler.Registro)
		auth.POST("/login", authHandler.Login)
	}

	// Cliente (protegida)
	clientes := router.Group("/clientes")
	clientes.Use(middleware.AuthRequired(authService))
	{
		clientes.GET("/perfil", clienteHandler.ObtenerPerfil)
		clientes.PUT("/perfil", clienteHandler.ActualizarPerfil)
	}

	// Carrito (protegida)
	carrito := router.Group("/carrito")
	carrito.Use(middleware.AuthRequired(authService))
	{
		carrito.GET("", carritoHandler.VerCarrito)
		carrito.POST("/items", carritoHandler.AgregarItem)
		carrito.PUT("/items/:id", carritoHandler.ModificarCantidad)
		carrito.DELETE("/items/:id", carritoHandler.EliminarItem)
		carrito.DELETE("", carritoHandler.VaciarCarrito)
	}

	// Pedidos (protegida)
	pedidos := router.Group("/pedidos")
	pedidos.Use(middleware.AuthRequired(authService))
	{
		pedidos.POST("", pedidoHandler.CrearPedido)
		pedidos.GET("", pedidoHandler.ListarPedidos)
		pedidos.GET("/:id", pedidoHandler.ObtenerPedido)
	}

	// === 6. Levantar servidor ===
	log.Println("Servidor iniciando en puerto:", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error al levantar servidor: %v", err)
	}
}
