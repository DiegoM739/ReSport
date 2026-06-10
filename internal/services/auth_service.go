package services
// Este service maneja el login y los tokens JWT.
import (
	"errors"
	"time"

	"github.com/DiegoM739/ReSport/internal/config"
	"github.com/DiegoM739/ReSport/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.ClienteRepository
	cfg  *config.Config
}

func NuevoAuthService(repo *repository.ClienteRepository, cfg *config.Config) *AuthService {
	return &AuthService{repo: repo, cfg: cfg}
}

// LoginResponse es la respuesta que se devuelve al hacer login exitoso.
type LoginResponse struct {
	Token  string `json:"token"`
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

// Login valida las credenciales y devuelve un token JWT.
func (s *AuthService) Login(email, password string) (*LoginResponse, error) {
	if email == "" || password == "" {
		return nil, errors.New("email y contraseña son obligatorios")
	}

	cliente, err := s.repo.BuscarPorEmail(email)
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Comparar contraseña ingresada vs hash guardado
	err = bcrypt.CompareHashAndPassword([]byte(cliente.Password), []byte(password))
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Generar token JWT
	token, err := s.GenerarToken(cliente.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:  token,
		UserID: cliente.ID,
		Email:  cliente.Email,
	}, nil
}

// GenerarToken crea un JWT con el ID del usuario.
func (s *AuthService) GenerarToken(userID uint) (string, error) {
	duracion := time.Hour * time.Duration(s.cfg.JWTExpirationHours)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duracion).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

// ValidarToken verifica que un token sea válido y devuelve el userID.
func (s *AuthService) ValidarToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWTSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("claims inválidos")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id inválido")
	}

	return uint(userID), nil
}