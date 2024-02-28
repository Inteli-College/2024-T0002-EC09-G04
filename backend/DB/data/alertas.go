package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq" // Driver PostgreSQL para Go
)

// Parâmetros de conexão com o banco de dados
const (
	host     = "database-stations.cxic0so62a43.us-east-1.rds.amazonaws.com"
	username = "postgres"
	password = "admin1234"
	database = "postgres"
)

// Tipo para representar um registro de alerta
type Alerta struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
	Opcoes    string    `json:"opcoes"`
}

// Função para inserir dados na tabela "alerta"
func inserirDados(alerta Alerta) error {
	sqlStatement := `INSERT INTO alerta (latitude, longitude, timestamp, opcoes) VALUES ($1, $2, $3, $4)`

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, username, password, database))
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement, alerta.Latitude, alerta.Longitude, alerta.Timestamp, alerta.Opcoes)
	if err != nil {
		return err
	}

	return nil
}

// Handler para requisições POST
func handlePost(w http.ResponseWriter, r *http.Request) {
	var alerta Alerta
	err := json.NewDecoder(r.Body).Decode(&alerta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = inserirDados(alerta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Dados inseridos com sucesso!")
}

func main() {
	http.HandleFunc("/alerta", handlePost)
	fmt.Println("Servidor ouvindo na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
