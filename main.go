package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	text := `
        Hello
		Test
		Test
    `
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, text)
	})
	e.Start(":1234")
}
