package models

import "gorm.io/gorm"

// Direccion representa una dirección de envío de un cliente.
type Direccion struct {
	gorm.Model
	Calle        string `gorm:"not null"`
	Ciudad       string `gorm:"not null"`
	Provincia    string
	CodigoPostal string
	Referencia   string

	// Relación: una dirección pertenece a un cliente
	ClienteID uint
}