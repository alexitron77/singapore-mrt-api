package http

import (
	"net/http"

	"github.com/labstack/echo"
	ctrl "zendesk.com/interview/mrt-backend/controllers"
	"zendesk.com/interview/mrt-backend/utils"
)

// Get all paths from src to destination
func GetPaths(c echo.Context) error {
	src := c.QueryParam("src")
	dest := c.QueryParam("dest")

	if src == "" || dest == "" {
		return c.String(http.StatusNotAcceptable, "Please set the 'src' and 'dest' query parameters in the url")
	}

	paths := ctrl.ComputePath(src, dest)
	utils.ItineraryInfo(src, dest, paths, []int{})
	return c.String(http.StatusOK, "See the logs")
}

// Get all paths from src to destination including travel time
func GetPathsWithTimeCost(c echo.Context) error {
	src := c.QueryParam("src")
	dest := c.QueryParam("dest")
	startTime := c.QueryParam("startTime")

	if src == "" || dest == "" || startTime == "" {
		return c.String(http.StatusBadRequest, "Please set the 'src' and 'dest' query parameters in the url")
	}

	paths := ctrl.ComputePath(src, dest)
	lineStnTime, lineChangeTime, err := ctrl.GetTravelShift(paths, startTime)

	if err != nil {
		return c.String(http.StatusBadRequest, "The start date format in incorrect")
	}

	computedTravelTime := ctrl.ComputeTravelTime(paths, lineStnTime, lineChangeTime)

	utils.ItineraryInfo(src, dest, paths, computedTravelTime)
	return c.String(http.StatusOK, "See the logs")
}
