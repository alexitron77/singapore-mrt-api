package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"zendesk.com/interview/mrt-backend/cmd"
)

func GetShortestPath(c echo.Context) error {
	src := c.QueryParam("src")
	dest := c.QueryParam("dest")

	computePath(src, dest)
	return c.String(http.StatusOK, "ans")
}

func computePath(src string, dest string) ([][]string, error) {
	visited := map[string]bool{}
	graph := cmd.BuildGraph()

	res := [][]string{}
	queue := [][]string{{src}}

	for len(queue) > 0 {
		// Select first element in the queue
		path := queue[0]

		// Shift element of the queue, equivalent of a .pop operation
		queue = queue[1:]

		// Select last node in the path
		lastNode := path[len(path)-1]
		visited[lastNode] = true

		// Check if the dest is reached
		if lastNode == dest {
			res = append(res, path)
			visited[lastNode] = false
		} else {
			for _, stn := range graph[lastNode] {
				if !visited[stn] {
					temp := make([]string, len(path))
					copy(temp, path)
					path := append(temp, stn)
					queue = append(queue, path)

					visited[stn] = true
				}
			}
		}
	}

	fmt.Print("Final res", res)

	return res, nil
}
