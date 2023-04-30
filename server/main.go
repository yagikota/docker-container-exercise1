package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Message string `json:"message"`
}

func greet(c echo.Context) error {
	response := Response{"Hello, World"}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", greet)
	e.Logger.Fatal(e.Start(":8080"))
}
