package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeanroths/jwtTeste666/auth"
	"github.com/jeanroths/jwtTeste666/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/users", auth.TokenVerificationMiddleware(handlers.GetUsersHandler)).Methods("GET")

	fmt.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
