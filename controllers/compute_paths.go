package controllers

import (
	models "zendesk.com/interview/mrt-backend/core/models"
)

// Compute all possible paths from source to destination
func ComputePath(src string, dest string) [][]string {
	visited := map[string]bool{}
	graph := models.BuildGraph()

	res := [][]string{}
	queue := [][]string{{src}}

	for len(queue) > 0 {
		// Select first element in the queue, equivalent of a pop operation
		path := queue[0]
		queue = queue[1:]

		// Select last station in the path
		lastNode := path[len(path)-1]
		visited[lastNode] = true

		// Check if the dest is reached
		if lastNode == dest {
			res = append(res, path)
			// Reset the visited value for future possible paths to be discovered
			visited[lastNode] = false
		} else {
			// Add the next stations to the possible paths and append those path to the queue
			for _, stn := range graph[lastNode] {
				if !visited[stn] {
					temp := make([]string, len(path))
					copy(temp, path)
					path := append(temp, stn)
					queue = append(queue, path)

					// Make sure that we can't revisit an already visited station
					visited[stn] = true
				}
			}
		}
	}
	return res
}
