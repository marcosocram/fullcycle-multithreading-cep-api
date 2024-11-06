package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Endereco interface {
	getCep() string
	getLogradouro() string
	getBairro() string
	getLocalidade() string
	getUf() string
}

type EnderecoViaCep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func (e EnderecoViaCep) getCep() string {
	return e.Cep
}

func (e EnderecoViaCep) getLogradouro() string {
	return e.Logradouro
}

func (e EnderecoViaCep) getBairro() string {
	return e.Bairro
}

func (e EnderecoViaCep) getLocalidade() string {
	return e.Localidade
}

func (e EnderecoViaCep) getUf() string {
	return e.Uf
}

type EnderecoBrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (e EnderecoBrasilAPI) getCep() string {
	return e.Cep
}

func (e EnderecoBrasilAPI) getLogradouro() string {
	return e.Street
}

func (e EnderecoBrasilAPI) getBairro() string {
	return e.Neighborhood
}

func (e EnderecoBrasilAPI) getLocalidade() string {
	return e.City
}

func (e EnderecoBrasilAPI) getUf() string {
	return e.State
}

type Resultado struct {
	Endereco Endereco
	API      string
	Err      error
}

func buscarBrasilAPI(ctx context.Context, cep string, ch chan Resultado) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ch <- Resultado{Err: err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- Resultado{Err: err}
		return
	}
	defer resp.Body.Close()

	var endereco EnderecoBrasilAPI
	if err := json.NewDecoder(resp.Body).Decode(&endereco); err != nil {
		ch <- Resultado{Err: err}
		return
	}

	ch <- Resultado{Endereco: endereco, API: "BrasilAPI"}
}

func buscarViaCEP(ctx context.Context, cep string, ch chan Resultado) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ch <- Resultado{Err: err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- Resultado{Err: err}
		return
	}
	defer resp.Body.Close()

	var endereco EnderecoViaCep
	if err := json.NewDecoder(resp.Body).Decode(&endereco); err != nil {
		ch <- Resultado{Err: err}
		return
	}

	ch <- Resultado{Endereco: endereco, API: "ViaCEP"}
}

func main() {
	cep := "88110798"
	ch := make(chan Resultado, 2)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go buscarBrasilAPI(ctx, cep, ch)
	go buscarViaCEP(ctx, cep, ch)

	select {
	case resultado := <-ch:
		if resultado.Err != nil {
			fmt.Printf("Erro: %v\n", resultado.Err)
			return
		}
		fmt.Printf("Resposta mais rápida da API %s:\n", resultado.API)
		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s\n",
			resultado.Endereco.getCep(),
			resultado.Endereco.getLogradouro(),
			resultado.Endereco.getBairro(),
			resultado.Endereco.getLocalidade(),
			resultado.Endereco.getUf())
	case <-ctx.Done():
		fmt.Println("Erro: timeout após 1 segundo.")
	}
}
