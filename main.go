package main

import (
	"github.com/labstack/echo"
	handlers "zendesk.com/interview/mrt-backend/http/handlers"
)

func main() {

	e := echo.New()

	e.GET("/get-itinerary", handlers.GetPaths)
	e.GET("/get-itinerary-time-cost", handlers.GetPathsWithTimeCost)

	e.Logger.Fatal(e.Start(":1323"))
}
