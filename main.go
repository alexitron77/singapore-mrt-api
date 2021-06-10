package main

import (
	"github.com/labstack/echo"
	handlers "zendesk.com/interview/mrt-backend/http/handlers"
)

func main() {

	e := echo.New()

	e.GET("/get-itinerary", handlers.GetShortestPath)
	e.GET("/get-itinerary-time-cost", handlers.GetShortestPathWithTimeCost)

	e.Logger.Fatal(e.Start(":1323"))
}
