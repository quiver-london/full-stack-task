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
	r.HandleFunc("/api/v1/products", handler.HandleListProducts).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/products/{productId}", handler.HandleGetProduct).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/products/{productId}", handler.HandleUpdateProduct).Methods(http.MethodPatch)
	r.HandleFunc("/api/v1/products/{productId}", handler.HandleDeleteProduct).Methods(http.MethodDelete)

	return r
}
