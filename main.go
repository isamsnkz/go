package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isamsnkz/go/config"
	"github.com/isamsnkz/go/handler"
	"github.com/isamsnkz/go/middleware"
)

func main() {
	config.ConnectDatabase()

	r := mux.NewRouter()
	r.Use(middleware.LogProtocol)
	r.Use(middleware.CORSMiddleware)

	r.HandleFunc("/clans", handler.GetClans).Methods("GET")
	r.HandleFunc("/clans", handler.CreateClan).Methods("POST")
	r.HandleFunc("/clans/{id:[0-9]+}", handler.GetClanByID).Methods("GET")
	r.HandleFunc("/clans/{id:[0-9]+}", handler.UpdateClan).Methods("PUT")
	r.HandleFunc("/clans/{id:[0-9]+}", handler.DeleteClan).Methods("DELETE")

	log.Println("🚀 Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
