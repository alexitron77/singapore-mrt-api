package main

import (
	"github.com/labstack/echo"
	handlers "zendesk.com/interview/mrt-backend/http/handlers"
)

func main() {

	e := echo.New()
	e.GET("/get-shortest-path", handlers.GetShortestPath)

	e.Logger.Fatal(e.Start(":1323"))
}
