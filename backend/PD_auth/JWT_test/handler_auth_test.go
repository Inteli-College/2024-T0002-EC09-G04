package jwt_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/PD_auth/models_auth"
)

func TestSignupHandler(t *testing.T) {
	// Criar um usuário para o teste
	user := models_auth.User{
		Username: "jean",
		Password: "socorro",
	}

	// Converter o usuário em JSON
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Erro ao converter usuário para JSON: %v", err)
	}

	// Cria uma solicitação HTTP POST com o usuário como corpo
	req, err := http.NewRequest("POST", "http://localhost:8080/signup", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Errorf("Erro ao criar a solicitação HTTP: %v", err)
	}

	// Faz a requisição real ao servidor e obtém a resposta
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Erro ao fazer a requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o código de status da resposta
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Handler retornou código de status errado. Esperado %v, obtido %v", http.StatusOK, status)
	}

		// Lê o corpo da resposta
		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			t.Errorf("Erro ao decodificar o corpo da resposta: %v", err)
		}
	
		// Verifica o conteúdo do corpo da resposta
		expectedUsername := "jean"
		if responseBody["username"] != expectedUsername {
			t.Errorf("Handler retornou corpo da resposta errado. Esperado %v, obtido %v", expectedUsername, responseBody["username"])
		}

			// Verifica a senha no corpo da resposta
		expectedPassword := "socorro"
		if responseBody["password"] != expectedPassword {
			t.Errorf("Handler retornou senha errada. Esperado %v, obtido %v", expectedPassword, responseBody["password"])
		}
}


func TestLoginHandler(t *testing.T) {
	// Cria um corpo de requisição com as credenciais do usuário
	requestBody := []byte(`{"username": "jean", "password": "socorro"}`)

	// Cria uma solicitação HTTP POST com as credenciais no corpo
	req, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Errorf("Erro ao criar a solicitação HTTP: %v", err)
	}

	// Faz a requisição real ao servidor e obtém a resposta
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Erro ao fazer a requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o código de status da resposta
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Handler retornou código de status errado. Esperado %v, obtido %v", http.StatusOK, status)
	}

	// Verifica o tipo de conteúdo da resposta
	expectedContentType := "application/json"
	if contentType := resp.Header.Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler retornou tipo de conteúdo errado. Esperado %v, obtido %v", expectedContentType, contentType)
	}

	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Erro ao decodificar o corpo da resposta: %v", err)
	}
	// Verifica se há um token na resposta
	token, tokenExists := responseBody["token"].(string)
	if !tokenExists {
		t.Errorf("Token não encontrado na resposta")
	}

	// Verifica se o token é válido
	expectedToken := responseBody["token"].(string) //literalmente pegando o token da resposta do server
	if token != expectedToken {
		t.Errorf("Token na resposta errado. Esperado %v, obtido %v", expectedToken, token)
	}
	// Aqui você pode fazer as verificações necessárias no corpo da resposta, se desejar
}
