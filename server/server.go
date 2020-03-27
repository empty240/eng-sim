package main

import (
	"net/http"
	"github.com/labstack/echo"
	"app/handler"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!2\n")
	})
	e.GET("/hello", handler.MainPage())
	e.GET("/api/json", handler.JsonTest())
	e.Logger.Fatal(e.Start(":8080"))
}