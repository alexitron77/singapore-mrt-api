package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"zendesk.com/interview/mrt-backend/conf"
)

func GetShortestPath(c echo.Context) error {
	src := c.QueryParam("src")
	dest := c.QueryParam("dest")

	computePath(src, dest)
	return c.String(http.StatusOK, "ans")
}

func computePath(src string, dest string) ([][]string, error) {
	visited := map[string]bool{}
	graph := conf.BuildGraph()

	res := [][]string{}
	queue := [][]string{}

	queue = append(queue, []string{src})

	for len(queue) > 0 {
		first := queue[0]
		lastNode := first[len(first)-1]
		queue = queue[1:]

		if visited[lastNode] {
			continue
		}
		visited[lastNode] = true

		if lastNode == dest {
			res = append(res, first)
		} else {
			next := graph[lastNode]
			for _, stn := range next {
				path := append(first, stn)
				queue = append(queue, path)
			}
		}
	}

	fmt.Print(res)

	return res, nil
}
