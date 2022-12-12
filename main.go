package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/codigo35/golang-api-swagger/docs"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

type StatusCodeHandlerPayload struct {
	Code int `json:"code"`
}

//	@title			Golang API Swagger
//	@version		1.0
//	@description	Exemplo de API Golang usando Echo integrado com Swagger via Swaggo

//	@contact.url	https://www.codigo35.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:1323
//	@BasePath	/
func main() {
	e := echo.New()
	e.GET("/", helloWorldHandler)
	e.POST("/status", statusCodeHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// HelloWorldHandler godoc
// @Summary      retorna um simples "Hello, World"
// @Description  endpoint para exemplo de uma chamada GET
// @Tags 		 codigo35
// @Success      200  {object}  string
// @Router       / [get]
func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// StatusCodeHandler godoc
// @Summary      retorna um status code recebido
// @Description  ao receber um status code via POST (JSON), retorna o valor recebido
// @Tags 		 codigo35
// @Success      200  {object} StatusCodeHandlerPayload
// @Failure      400  {object} string
// @Router       /status [post]
func statusCodeHandler(c echo.Context) error {
	payload := new(StatusCodeHandlerPayload)
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("Status Code: %d", payload.Code))
}
