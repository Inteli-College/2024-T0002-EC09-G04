package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Inteli-College/2024-T0002-EC09-G04@56-pd---auth/backend/P&D/auth"
	"github.com/Inteli-College/2024-T0002-EC09-G04@56-pd---auth/backend/P&D/handlers_auth"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/users", auth.TokenVerificationMiddleware(handlers.GetUsersHandler)).Methods("GET")

	fmt.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
