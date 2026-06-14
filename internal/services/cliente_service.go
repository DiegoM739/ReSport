package services

// Crea el usuario, encripta y hashea la contrseñ a
// Obtiene un perfil y actualiza el perfil
// Sevice del registro

import (
	"errors"

	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type ClienteService struct {
	repo repository.IClienteRepository
}

func NuevoClienteService(repo repository.IClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

// Registrar valida los datos, encripta la contraseña y crea el cliente.
func (s *ClienteService) Registrar(cliente *models.Cliente) error {
	// Validaciones
	if cliente.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if cliente.Email == "" {
		return errors.New("el email es obligatorio")
	}
	if len(cliente.Password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	// Verificar que el email no esté registrado
	existente, _ := s.repo.BuscarPorEmail(cliente.Email)
	if existente != nil {
		return errors.New("el email ya está registrado")
	}

	// Encriptar contraseña con bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(cliente.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error al procesar la contraseña")
	}
	cliente.Password = string(hash)

	return s.repo.Crear(cliente)
}

func (s *ClienteService) ObtenerPerfil(id uint) (*models.Cliente, error) {
	return s.repo.BuscarPorID(id)
}

func (s *ClienteService) ActualizarPerfil(cliente *models.Cliente) error {
	if cliente.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	return s.repo.Actualizar(cliente)
}
