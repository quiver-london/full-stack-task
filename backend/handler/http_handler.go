package handler

import (
	"backend/storage"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type HttpHandler struct {
	Storage storage.Storage
}

func (h *HttpHandler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product storage.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.Id = primitive.NewObjectID()

	err = h.Storage.Create(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rawResponse, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(rawResponse)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (h *HttpHandler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	productId := parts[len(parts)-1]

	post, err := h.Storage.GetProduct(r.Context(), productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	rawResponse, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(rawResponse)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
