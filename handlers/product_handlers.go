package handlers

import (
	"encoding/json"
	"fmt"
	"gin-api/db"
	"gin-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var p models.Product

    // Mejor manejo de errores y logging
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
        // Log the error for debugging purposes
                log.Printf("Error decoding request body: %v", err) // or your preferred logging method
        return
    }

        // Log the received product data (for debugging/audit trails if needed)
        log.Printf("Received product data: %+v", p) // %+v prints struct fields

    if err := db.CreateProduct(&p); err != nil {
        http.Error(w, fmt.Sprintf("Error creating product: %v", err), http.StatusInternalServerError)
        // Log the error for debugging
                log.Printf("Error creating product: %v", err)
        return
    }

    w.WriteHeader(http.StatusCreated)

    // Return the created product with a 201 Created status.  This is good practice.
    if err := json.NewEncoder(w).Encode(p); err != nil {
        http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
                log.Printf("Error encoding response: %v", err)
        return
    }
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        log.Printf("Received params data: %+v", params) 
        
        id, err := strconv.Atoi(params["id"])
        log.Printf("Received params data: %+v", id) 
        if err != nil {
                http.Error(w, "ID inválido", http.StatusBadRequest)
                return
        }

        product, err := db.GetProduct(id)
        if err != nil {
                http.Error(w, err.Error(), http.StatusNotFound)
                return
        }

        json.NewEncoder(w).Encode(product)
		
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
        products, err := db.GetAllProducts()
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        json.NewEncoder(w).Encode(products)
		fmt.Println(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id, err := strconv.Atoi(params["id"])
        if err != nil {
                http.Error(w, "ID inválido", http.StatusBadRequest)
                return
        }

        var p models.Product
        if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        p.ID = id

        if err := db.UpdateProduct(&p); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id, err := strconv.Atoi(params["id"])
        if err != nil {
                http.Error(w, "ID inválido", http.StatusBadRequest)
                return
        }

        if err := db.DeleteProduct(id); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusNoContent)
}