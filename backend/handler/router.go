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

type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}
