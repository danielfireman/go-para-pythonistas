package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// INICIO FLAGS OMIT
var (
	url         = flag.String("url", "http://localhost:8000", "URL do servidor de TODOS")
	corpo       = flag.String("corpo", "", "Corpo da requisição")
	metodo      = flag.String("metodo", "GET", "Método HTTP a ser utilização na requisição")
	contentType = flag.String("content-type", "application/json", "Cabeçalho Content-Type")
)

// FIM FLAGS OMIT

// INICIO MAIN OMIT
func main() {
	flag.Parse()

	// Criando requisição.
	// Atenção para strings.NewReader(). Interfaces é assunto para outro tutorial!
	// Atenção para o "*" antes do nome da variável.
	req, err := http.NewRequest(*metodo, *url, strings.NewReader(*corpo)) // Tratar erro!
	if *contentType != "" {
		req.Header.Add("Content-Type", *contentType)
	}

	// Disparando requisição.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("erro realizando requisiçãp - %+v):%q", req, err)
	}
	defer resp.Body.Close()

	// Lendo corpo da resposta.
	corpoResp, _ := ioutil.ReadAll(resp.Body) // Tratar erro!
	fmt.Printf("Status:%v | Corpo:%s\n", resp.StatusCode, string(corpoResp))
}

// FIM MAIN OMIT
