package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func insertMockData(db *sql.DB) error {
	// Preparar a declaração SQL de inserção
	stmt, err := db.Prepare("INSERT INTO Gas (id_gas, id_estacao, CO2, CO, NO2, MP10, MP25) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Dados mockados para inserção
	mockData := []struct {
		IDEstacao int
		CO2       float64
		CO        float64
		NO2       float64
		MP10      float64
		MP25      float64
	}{
		{1, 100, 50, 25, 10, 5},
		{1, 150, 60, 30, 15, 8},
		{1, 200, 70, 35, 20, 10},
	}

	// Inserir dados mockados na tabela
	for _, data := range mockData {
		_, err := stmt.Exec(4, data.IDEstacao, data.CO2, data.CO, data.NO2, data.MP10, data.MP25)
		if err != nil {
			return err
		}
	}

	fmt.Println("Dados mockados inseridos com sucesso.")
	return nil
}

func test() {
	// String de conexão com o banco de dados PostgreSQL
	connStr := "user=postgres password=admin1234 dbname=postgres sslmode=disable"

	// Tenta abrir uma conexão com o banco de dados PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Tenta pingar o banco de dados para testar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao pingar o banco de dados:", err)
	}

	// Chama a função para inserir dados mockados
	err = insertMockData(db)
	if err != nil {
		log.Fatal("Erro ao inserir dados mockados:", err)
	}

	fmt.Println("Teste de inserção de dados mockados concluído com sucesso.")
}
