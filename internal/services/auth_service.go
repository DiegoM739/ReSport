package services

import (
	"errors"
	"time"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/models"
	"github.com/DiegoM739/ReSport/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService maneja la autenticación de clientes.
type AuthService struct {
	clienteRepo repository.IClienteRepository
	config      *config.Config
}

// NuevoAuthService crea una nueva instancia del service.
func NuevoAuthService(clienteRepo repository.IClienteRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		clienteRepo: clienteRepo,
		config:      cfg,
	}
}

// Login valida las credenciales y devuelve un token JWT.
func (s *AuthService) Login(email, password string) (string, *models.Cliente, error) {
	// 1. Buscar cliente por email
	cliente, err := s.clienteRepo.BuscarPorEmail(email)
	if err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	// 2. Comparar el hash de la contraseña con bcrypt
	err = bcrypt.CompareHashAndPassword(
		[]byte(cliente.Password),
		[]byte(password),
	)
	if err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	// 3. Generar token JWT
	token, err := s.GenerarToken(cliente.ID)
	if err != nil {
		return "", nil, errors.New("error al generar token")
	}

	return token, cliente, nil
}

// GenerarToken crea un token JWT firmado con la clave secreta.
func (s *AuthService) GenerarToken(userID uint) (string, error) {
	// Calcular fecha de expiración
	expiracion := time.Now().Add(time.Duration(s.config.JWTExpirationHours) * time.Hour)

	// Crear los claims del token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiracion.Unix(),
		"iat":     time.Now().Unix(),
	}

	// Crear el token con el algoritmo HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidarToken verifica un token JWT y devuelve el user_id.
func (s *AuthService) ValidarToken(tokenString string) (uint, error) {
	// Parsear el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el algoritmo sea el esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return []byte(s.config.JWTSecret), nil
	})

	if err != nil {
		return 0, errors.New("token inválido")
	}

	// Extraer claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("token inválido")
	}

	// Extraer el user_id
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("token sin user_id")
	}

	return uint(userIDFloat), nil
}
