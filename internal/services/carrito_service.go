package services

import (
	"errors"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/repository"
)

// CarritoService contiene la lógica de negocios del carrito
type CarritoService struct {
	carritoRepo  *repository.CarritoRepository
	productoRepo *repository.ProductoRepository
}

// Herencia. NuevoCarritoService crea una nueva instancia del service
// Necesita ambos repositorios porque trabaja con porductos y carrito
func NuevoCarritoService(
	carritoRepo *repository.CarritoRepository,
	productoRepo *repository.ProductoRepository,
) *CarritoService {
	return &CarritoService{
		carritoRepo:  carritoRepo,
		productoRepo: productoRepo,
	}
}

// ObtenerCarrito devuelve el carrito del cliente con todos sus items.
// Si el cliente nunca ha tenido carrito, le crea uno vacio
func (s *CarritoService) ObtenerCarrito(clienteID uint) (*models.Carrito, error) {
	// Pedirle al repository el carrito (lo crea si no existe)
	carrito, err := s.carritoRepo.ObtenerOCrearPorCliente(clienteID)
	if err != nil {
		return nil, err
	}
	// Calcular el total sumando los subtotales al consultar
	// Esta lógica está en el método del producto Carrito, es decir usamos encapsulación
	total := carrito.CalcularTotal()

	// El total no se guarda en DB, se calcula al consultar
	//Esto evita que el total quede desincronizado con los items

	_ = total // Esto es para que el total solo se utilice aqui

	return carrito, nil

}

// AgregarProducto agrega un producto al carrito del cliente.
func (s *CarritoService) AgregarProducto(clienteID uint, productoID uint, cantidad int) error {
	// === Paso 1: Validar la cantidad ===
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	// === Paso 2: Buscar el producto en la base ===
	producto, err := s.productoRepo.BuscarPorID(productoID)
	if err != nil {
		return errors.New("producto no encontrado")
	}

	// === Paso 3: Verificar que haya stock suficiente ===
	// Usamos el método del producto (encapsulación de la regla)
	if !producto.VerificarStock(cantidad) {
		return errors.New("stock insuficiente")
	}

	// === Paso 4: Obtener el carrito del cliente ===
	carrito, err := s.carritoRepo.ObtenerOCrearPorCliente(clienteID)
	if err != nil {
		return errors.New("error al obtener el carrito")
	}

	// === Paso 5: Crear el item con los datos ===
	subtotal := producto.Precio * float64(cantidad)

	nuevoItem := &models.ItemCarrito{
		CarritoID:  carrito.ID,
		ProductoID: producto.ID,
		Cantidad:   cantidad,
		Subtotal:   subtotal,
	}

	// === Paso 6: Guardar el item en la base ===
	err = s.carritoRepo.AgregarItem(nuevoItem)
	if err != nil {
		return errors.New("error al agregar el producto al carrito")
	}

	return nil
}

// ModificarCantidadItem cambia la cantidad de un item del carrito.
func (s *CarritoService) ModificarCantidadItem(itemID uint, nuevaCantidad int) error {
	// === Paso 1: Validar ===
	if nuevaCantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	// === Paso 2: Buscar el item ===
	item, err := s.carritoRepo.BuscarItem(itemID)
	if err != nil {
		return errors.New("item no encontrado")
	}

	// === Paso 3: Verificar stock del producto asociado ===
	if !item.Producto.VerificarStock(nuevaCantidad) {
		return errors.New("stock insuficiente")
	}

	// === Paso 4: Actualizar cantidad y subtotal ===
	item.Cantidad = nuevaCantidad
	item.Subtotal = item.Producto.Precio * float64(nuevaCantidad)

	// === Paso 5: Guardar ===
	return s.carritoRepo.ActualizarItem(item)
}

// EliminarItem borra un item del carrito.
func (s *CarritoService) EliminarItem(itemID uint) error {
	return s.carritoRepo.EliminarItem(itemID)
}

// VaciarCarrito remueve todos los items del carrito del cliente.
func (s *CarritoService) VaciarCarrito(clienteID uint) error {
	carrito, err := s.carritoRepo.ObtenerOCrearPorCliente(clienteID)
	if err != nil {
		return err
	}
	return s.carritoRepo.VaciarCarrito(carrito.ID)
}
