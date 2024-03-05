package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	e := echo.New()
	text := `
        Hello
    `

	// Define routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, text)
	})

	// Start the server
	e.Start(":1234")
}
