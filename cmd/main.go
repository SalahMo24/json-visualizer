package main

import (
	"json-visualizer/pkg/handler"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	path, err := filepath.Abs("static")
	if err != nil {
		panic(path)
	}
	app.Static("/static", path)
	userHandeler := handler.UserHandler{}
	app.GET("/", userHandeler.HandleUser)

	app.Start(":3000")
}
