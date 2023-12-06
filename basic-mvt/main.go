package main

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed static/*
var static embed.FS

func main() {
	e := echo.New()
	e = setMiddlewares(e)
	e = setRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
