package main

import (
	"github.com/danisbagus/golang-elasticsearch/interface/api"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	apiRoute := e.Group("/api")
	api.API(apiRoute)

	if err := e.Start(":" + "8080"); err != nil {
		e.Logger.Info("Shutting down the server")
	}
}
