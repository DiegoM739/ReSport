package models

import "gorm.io/gorm"

// Categoria agrupa productos similares (ej: "Recuperación", "Suplementación").
type Categoria struct {
	gorm.Model
	Nombre      string `gorm:"unique;not null"`
	Descripcion string
	Productos   []Producto // una categoría tiene muchos productos
}

// Producto representa un artículo a la venta.
// El campo Tipo diferencia físicos de digitales (patrón Single Table Inheritance).
type Producto struct {
	gorm.Model
	Nombre      string  `gorm:"not null"`
	Descripcion string
	Precio      float64 `gorm:"not null"`
	Stock       int     `gorm:"default:0"`
	Imagen      string
	Tipo        string  `gorm:"not null"` // "fisico" o "digital"

	// Campos específicos de producto físico
	Peso        float64
	Dimensiones string

	// Campos específicos de producto digital
	URLDescarga string

	// Relación con categoría (un producto pertenece a una categoría)
	CategoriaID uint
	Categoria   Categoria
}

// VerificarStock comprueba si hay stock suficiente para una cantidad dada.
// Esto es ENCAPSULACIÓN: la regla vive dentro del producto.
func (p *Producto) VerificarStock(cantidad int) bool {
	return p.Stock >= cantidad
}

// AplicarDescuento devuelve el precio con un descuento aplicado.
func (p *Producto) AplicarDescuento(porcentaje float64) float64 {
	return p.Precio * (1 - porcentaje)
}