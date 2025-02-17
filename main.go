package main

import (
	"gin-api/db"
	"gin-api/handlers" // Importamos el nuevo paquete handlers
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
        db.ConnectDB()
        defer db.DB.Close()

        r := mux.NewRouter()

        r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")       
        r.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")       
        r.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")      
        r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")     
        r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE") 

        port := os.Getenv("PORT")
        if port == "" {
                port = "8088"
        }

        log.Printf("Servidor escuchando en el puerto %s", port)
        log.Fatal(http.ListenAndServe(":"+port, r))
}