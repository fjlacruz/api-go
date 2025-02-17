package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB establece la conexión con la base de datos.
func ConnectDB() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Verificación de que las variables de entorno se han cargado
	log.Println("Verificando las variables de entorno...")

	// Obtener las variables de entorno necesarias
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Comprobar si las variables de entorno están vacías
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Faltan variables de entorno necesarias")
	}


	// Construir la cadena de conexión usando las variables de entorno
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", 
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Intentar abrir la conexión
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir la conexión:", err)
	}

	// Intentar hacer ping a la base de datos
	if err = DB.Ping(); err != nil {
		log.Fatal("Error al hacer ping a la base de datos:", err)
	}

	// Imprimir mensaje si la conexión es exitosa
	fmt.Println("Conexión a la base de datos establecida.")
}
