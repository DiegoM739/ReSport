package services

import (
	"errors"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// ClienteService maneja la lógica de clientes.
type ClienteService struct {
	repo repository.IClienteRepository
}

// NuevoClienteService crea una nueva instancia.
func NuevoClienteService(repo repository.IClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

// Registrar crea un cliente nuevo con contraseña hasheada.
// Recibe el cliente y la contraseña en texto plano.
func (s *ClienteService) Registrar(cliente *models.Cliente, passwordPlana string) error {
	// 1. Validar datos básicos
	if cliente.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if cliente.Email == "" {
		return errors.New("el email es obligatorio")
	}
	if passwordPlana == "" {
		return errors.New("la contraseña es obligatoria")
	}
	if len(passwordPlana) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	// 2. Verificar que el email no esté registrado
	existente, _ := s.repo.BuscarPorEmail(cliente.Email)
	if existente != nil {
		return errors.New("el email ya está registrado")
	}

	// 3. Hashear la contraseña con bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordPlana), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error al procesar la contraseña")
	}
	cliente.Password = string(hash)

	// 4. Guardar en la base de datos
	return s.repo.Crear(cliente)
}

// ObtenerPerfil busca un cliente por su ID.
func (s *ClienteService) ObtenerPerfil(id uint) (*models.Cliente, error) {
	cliente, err := s.repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("cliente no encontrado")
	}
	return cliente, nil
}

// ActualizarPerfil modifica los datos de un cliente.
func (s *ClienteService) ActualizarPerfil(cliente *models.Cliente) error {
	if cliente.ID == 0 {
		return errors.New("ID del cliente es obligatorio")
	}
	if cliente.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	return s.repo.Actualizar(cliente)
}