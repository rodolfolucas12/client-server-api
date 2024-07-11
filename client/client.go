package client

import (
	model "client-server-api/server/model"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func BuscarCotacao() {
	inicio := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Falha ao criar requisicao: ", err)
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Falha ao buscar cotacao: ", err)
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Falha ao ler resposta")
		return
	}
	defer resp.Body.Close()

	var cotacao model.Cotacao

	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Fatal("Erro ao deserializar response cotacao", err)
		return
	}

	salvarCotacao(cotacao, inicio)
}

func salvarCotacao(cotacao model.Cotacao, inicio time.Time) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal("erro ao criar arquivo de cotacao: ", err)
		return err
	}
	defer file.Close()
	file.WriteString("Dólar: " + cotacao.Dolar)

	log.Println("Arquivo de cotacao salvo com sucesso, tempo de execução: ", time.Since(inicio))
	return nil
}
