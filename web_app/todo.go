// +build ignore,OMIT

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// INICIO MAIN OMIT
func main() {
	e := echo.New()

	e.PUT("/", novo)
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
// Atenção para a declaração de função
func novo(c echo.Context) error {
	// Declarando e construindo variável
	t := todo{}

	// Fazendo parse do corpo da requisição e preenchendo a struct
	// Atenção para o "&" e para o ":="
	err := c.Bind(&t)

	// Estrutura condicional e checagem de erros
	if err != nil || t.Desc == "" {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "erro criando novo todo")
	}

	// Atualizando ID e lista de todos
	// Atenção para o "="
	t.ID = len(todos)
	todos = append(todos, t)

	return c.JSON(http.StatusCreated, t)
}

// FIM NOVO OMIT

// INICIO BUSCA OMIT
func busca(c echo.Context) error {
	// Acessando parâmetro da URL
	idStr := c.Param("id")

	// Conversão de tipo
	// Atenção para função retornando dois valores: resultado e erro
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("identificador inválido:%q", err))
	}

	// Estrutura de repetição
	for _, t := range todos {
		if t.ID == id {
			return c.JSON(http.StatusFound, t)
		}
	}
	return c.String(http.StatusNotFound, idStr)
}

// FIM BUSCA OMIT

// INICIO LISTA OMIT
func lista(c echo.Context) error {
	// Caso especial para quando o lista for invocado antes de alguma adição.
	if todos == nil {
		// Atenção para o construtor.
		return c.JSON(http.StatusOK, []todo{})
	}
	return c.JSON(http.StatusOK, todos)
}

// FIM LISTA OMIT
