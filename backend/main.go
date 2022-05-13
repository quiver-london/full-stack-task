package main

import (
	"backend/handler"
	"backend/storage/mongostorage"
	"log"
	"net/http"
	"os"
	"time"
)

func NewServer() *http.Server {

	mongoUrl := os.Getenv("MONGO_URL")
	mongoStorage := mongostorage.DatabaseStorage(mongoUrl)
	router := handler.CreateRouterFromStorage(mongoStorage)

	return &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func main() {
	srv := NewServer()
	log.Printf("Start serving on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
