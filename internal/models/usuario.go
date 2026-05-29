package models

import "gorm.io/gorm"

// Persona contiene los datos comunes a todos los usuarios.
// NO es una tabla por sí misma: se embebe en Cliente y Administrador.
// Esto es HERENCIA al estilo Go (composición).
type Persona struct {
	Nombre   string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// Cliente representa a un usuario que compra en la tienda.
// Hereda los campos de Persona mediante composición.
type Cliente struct {
	gorm.Model
	Persona               // composición = herencia
	Telefono    string
	Direcciones []Direccion // un cliente tiene muchas direcciones (relación 1:N)
}

// Administrador representa a un usuario con permisos de gestión.
type Administrador struct {
	gorm.Model
	Persona               // composición = herencia
	Rol string `gorm:"default:admin"`
}


