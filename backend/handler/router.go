package handler

import (
	"backend/storage"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateRouterFromStorage(cachedStorage storage.Storage) *mux.Router {
	handler := &HttpHandler{
		Storage: cachedStorage,
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/products", handler.HandleCreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/products/{postId}", handler.HandleGetProduct).Methods(http.MethodGet)

	return r
}
