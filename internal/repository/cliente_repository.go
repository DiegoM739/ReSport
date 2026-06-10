package repository

//archivo que habla directamente con la base de datos 
//Es para todo lo relacionado con clientes 
// Realiza crea, buscar por email, buscar por ID, actualizar
// Algo importante es que sin importar si se modifica la base de datos esto funciona porque tiene secciones 
//que permitan que sea remplazable la BD y reutilizable el código. 
import (
	"github.com/DiegoM739/ReSport/internal/models"
	"gorm.io/gorm"
)

type ClienteRepository struct {
	db *gorm.DB
}

func NuevoClienteRepository(db *gorm.DB) *ClienteRepository {
	return &ClienteRepository{db: db}
}

func (r *ClienteRepository) Crear(cliente *models.Cliente) error {
	return r.db.Create(cliente).Error
}

// BuscarPorEmail busca un cliente por su email (sirve para login).
func (r *ClienteRepository) BuscarPorEmail(email string) (*models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.Where("email = ?", email).First(&cliente).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}
// Busca por ID 
func (r *ClienteRepository) BuscarPorID(id uint) (*models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.Preload("Direcciones").First(&cliente, id).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

// Actualiza 
func (r *ClienteRepository) Actualizar(cliente *models.Cliente) error {
	return r.db.Save(cliente).Error
}