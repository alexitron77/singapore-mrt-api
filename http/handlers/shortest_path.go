package http

import (
	"net/http"

	"github.com/labstack/echo"
	ctrl "zendesk.com/interview/mrt-backend/controllers"
	"zendesk.com/interview/mrt-backend/utils"
)

func GetShortestPath(c echo.Context) error {
	src := c.QueryParam("src")
	dest := c.QueryParam("dest")

	paths := ctrl.ComputePath(src, dest)
	utils.ItineraryInfo(src, dest, paths)
	return c.String(http.StatusOK, "ans")
}
