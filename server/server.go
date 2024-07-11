package server

import (
	"client-server-api/repository"
	client "client-server-api/server/client"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func NewServer() {
	router := http.NewServeMux()
	router.HandleFunc("/cotacao", HandlerCotacao)

	http.ListenAndServe(":8080", router)
}

func HandlerCotacao(w http.ResponseWriter, r *http.Request) {
	inicio := time.Now()
	dolar, err := client.GetDollarPrice()
	if err != nil {
		log.Fatal("Falha ao obter cotacao: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = repository.InsertDollarPrice(dolar.Dolar)
	if err != nil {
		log.Fatal("Falha ao inserir cotacao do dolar: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Preço do dólar salvo no banco de dados, tempo de execução: ", time.Since(inicio))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dolar)
}
