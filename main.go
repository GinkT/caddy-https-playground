package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", Handler)

	if os.Getenv("SSL") == "true" {
		log.Fatal(e.StartTLS(":12321", "/certs/cc.cert", "/certs/cc.key"))
	} else {
		log.Fatal(e.Start(":12321"))
	}
}

func Handler(e echo.Context) (err error) {
	log.Println("handler")

	return e.JSON(http.StatusOK, echo.Map{
		"hi": "test",
	})
}
