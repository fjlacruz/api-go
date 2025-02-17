package handlers

import (
	"encoding/json"
	"gin-api/db"
	"gin-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
        var p models.Product
        if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        if err := db.CreateProduct(&p); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(p)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id, err := strconv.Atoi(params["id"])
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