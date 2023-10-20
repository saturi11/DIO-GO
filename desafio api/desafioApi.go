package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Cliente é uma estrutura para representar um cliente
type Cliente struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var clientes []Cliente

func main() {
	r := mux.NewRouter()
	clientes = append(clientes, Cliente{ID: "1", Nome: "Cliente1", Email: "cliente1@example.com"})

	r.HandleFunc("/clientes", GetClientes).Methods("GET")
	r.HandleFunc("/clientes/{id}", GetCliente).Methods("GET")
	r.HandleFunc("/clientes", CreateCliente).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Servidor está rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// GetClientes retorna a lista de todos os clientes
func GetClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

// GetCliente retorna os detalhes de um cliente específico com base no ID
func GetCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, cliente := range clientes {
		if cliente.ID == params["id"] {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// CreateCliente cria um novo cliente e o adiciona à lista
func CreateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cliente Cliente
	_ = json.NewDecoder(r.Body).Decode(&cliente)
	clientes = append(clientes, cliente)
	json.NewEncoder(w).Encode(cliente)
}
