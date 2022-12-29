package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", Handler)

	log.Fatal(e.StartTLS(":12321", "/certs/cc.cert", "/certs/cc.key"))
}

func Handler(e echo.Context) (err error) {
	return e.JSON(http.StatusOK, echo.Map{
		"hi": "test",
	})
}
