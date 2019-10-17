// +build ignore,OMIT

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// INICIO MAIN OMIT
func main() {
	e := echo.New()

	e.PUT("/", novo)
	e.POST("/:id", atualiza)
	e.DELETE("/:id", apaga)
	e.GET("/:id", busca)
	e.GET("/lista", lista)

	e.Logger.Fatal(e.Start(":8000"))
}

// FIM MAIN OMIT

// INICIO STRUCT OMIT
type todo struct {
	ID   int    `json:"id"`
	Desc string `json:"desc"`
}

var todos []todo

// FIM STRUCT OMIT

// INICIO NOVO OMIT
func novo(c echo.Context) error {
	var t todo

	// Fazendo parse do corpo da requisição e preenchendo a struct.
	err := c.Bind(&t)

	// Checando erros.
	if err != nil || t.Desc == "" {
		log.Printf("erro criando novo todo (%+v): %q", t, err)
		return c.String(http.StatusBadRequest, "erro criando novo todo")
	}

	// Atualizando ID e lista de todos.
	t.ID = len(todos)
	todos = append(todos, t)

	// Retornando resposta para usuário.
	return c.JSON(http.StatusCreated, t)
}

// FIM NOVO OMIT

func atualiza(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func apaga(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func busca(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func lista(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
