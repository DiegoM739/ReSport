package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config almacena toda la configuración del sistema
type Config struct {
	Port               string
	DBPath             string
	JWTSecret          string
	JWTExpirationHours int
	Env                string
}

// Cargar lee el archivo .env y devuelve la configuración
func Cargar() *Config {
	// Cargar variables del archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: no se encontró archivo .env, usando variables del sistema")
	}

	// Leer cada variable, con valores por defecto si no existen
	jwtHours, err := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))
	if err != nil {
		jwtHours = 24
	}

	return &Config{
		Port:               getEnv("PORT", "8080"),
		DBPath:             getEnv("DB_PATH", "resport.db"),
		JWTSecret:          getEnv("JWT_SECRET", "secreto_por_defecto_inseguro"),
		JWTExpirationHours: jwtHours,
		Env:                getEnv("ENV", "development"),
	}
}

// getEnv lee una variable de entorno; si no existe, devuelve el valor por defecto
func getEnv(clave, valorDefault string) string {
	if valor, existe := os.LookupEnv(clave); existe {
		return valor
	}
	return valorDefault
}