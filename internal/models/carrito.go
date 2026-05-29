package models

import "gorm.io/gorm"

// Carrito representa el carrito de compras de un cliente.
type Carrito struct {
	gorm.Model
	ClienteID uint
	Items     []ItemCarrito
}

// ItemCarrito representa un producto con su cantidad dentro de un carrito.
type ItemCarrito struct {
	gorm.Model
	CarritoID  uint
	ProductoID uint
	Producto   Producto
	Cantidad   int
	Subtotal   float64
}

// CalcularSubtotal calcula el subtotal de este item (precio × cantidad).
// ENCAPSULACIÓN: el item sabe calcular su propio subtotal.
func (i *ItemCarrito) CalcularSubtotal() {
	i.Subtotal = i.Producto.Precio * float64(i.Cantidad)
}

// CalcularTotal suma los subtotales de todos los items del carrito.
func (c *Carrito) CalcularTotal() float64 {
	var total float64
	for _, item := range c.Items {
		total += item.Subtotal
	}
	return total
}