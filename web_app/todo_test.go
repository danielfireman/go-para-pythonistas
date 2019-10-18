package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// INICIO NOVO OMIT
// Atenção para o nome especial da função e o ponteiro como parâmetro.
func TestNovo(t *testing.T) {
	// Item a ser inserido
	todoJSON := `{"id":0,"desc":"tapioca"}`

	// Requisição e Contexto
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()

	// Invocação
	err := novo(echo.New().NewContext(req, resp))

	// Checagens
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Equal(t, todoJSON, strings.Trim(resp.Body.String(), "\n"))
}

// FIM NOVO OMIT

// INICIO BUSCA OMIT
func TestBusca(t *testing.T) {
	// Construtor.
	novo := todo{
		ID:   0,
		Desc: "tapioca",
	}

	// Atualizando diretamente lista de todos.
	todos = append(todos, novo)
	defer func() { todos = todos[:0] }() // Limpar lista depois da execução.

	// Requisição e Contexto
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := echo.New().NewContext(req, resp)
	c.SetParamNames("id")
	c.SetParamValues("0")

	// Invocação
	err := busca(c)

	// Checagens
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.Code)

	// Atenção para a conversão do corpo da resposta para JSON e
	// preenchimento da struct ret.
	var ret todo
	assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &ret))
	assert.Equal(t, novo, ret)
}

// FIM BUSCA OMIT
