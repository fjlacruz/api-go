package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	_ = godotenv.Load() // No falla si no encuentra .env

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Faltan variables de entorno necesarias")
	}

	// Detectar si la app está corriendo en Docker
	if os.Getenv("IN_DOCKER") != "true" {
		dbHost = "localhost" // Conexión desde fuera de Docker
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=5",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	var err error
	for i := 0; i < 5; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err == nil && DB.Ping() == nil {
			log.Println("✅ Conexión exitosa a la base de datos")
			return
		}
		log.Printf("Intento %d: Error al conectar con la base de datos: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("❌ No se pudo establecer conexión con la base de datos después de varios intentos")
}
