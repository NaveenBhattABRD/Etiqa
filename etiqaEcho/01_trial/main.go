package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.GET("/products", func(c echo.Context) error {

	// 	return c.JSON(http.StatusOK, products)

	// })

	// add middleware and routes
	// ...
	if err := e.Start(":4442"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
