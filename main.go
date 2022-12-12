package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusCodeHandlerPayload struct {
	Code int `json:"code"`
}

func main() {
	e := echo.New()
	e.GET("/", helloWorldHandler)
	e.POST("/status", statusCodeHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func statusCodeHandler(c echo.Context) error {
	payload := new(StatusCodeHandlerPayload)
	if err := c.Bind(payload); err != nil {
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("Status Code: %d", payload.Code))
}
