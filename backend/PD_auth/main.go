package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/PD_auth/auth"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/PD_auth/handlers_auth"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers_auth.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers_auth.LoginHandler).Methods("POST")
	router.HandleFunc("/users", auth.TokenVerificationMiddleware(handlers_auth.GetUsersHandler)).Methods("GET")

	fmt.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
