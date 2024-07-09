package server

import (
	model "client-server-api/server/model"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func GetDollarPrice() (model.Cotacao, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/last/USD-BRL", nil)
	if err != nil {
		log.Fatal("Erro na requisicao: ", err)
		return model.Cotacao{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Falha ao buscar cotacao: ", err)
		return model.Cotacao{}, err
	}
	defer resp.Body.Close()

	return readBody(resp)
}

func readBody(resp *http.Response) (model.Cotacao, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Falha ao ler resposta")
		return model.Cotacao{}, err
	}

	var cotacao model.CotacaoResponse

	json.Unmarshal(body, &cotacao)

	return model.ToDomain(cotacao), nil
}
