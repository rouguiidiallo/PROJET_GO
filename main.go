package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"restaurant-order-management/db"
	"restaurant-order-management/handlers"
)

func main() {
	router := mux.NewRouter()

	db.InitDB()
	defer db.CloseDB()

	router.HandleFunc("/api/orders", handlers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/{id}", handlers.GetOrder).Methods("GET")
	router.HandleFunc("/api/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/orders/{id}", handlers.DeleteOrder).Methods("DELETE")

	// Servir les fichiers statiques
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
	// Message dans la console indiquant que le serveur est démarré
    
}
