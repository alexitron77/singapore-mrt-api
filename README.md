# Introduction

You are provided data on the stations and lines of Singapore's urban rail system, including planned additions over the next few years. Your task is to use this data to build a routing service, to help users find routes from any station to any other station on this future network.

The app should expose an API to find and display one or more routes from a specified origin to a specified destination, ordered by some efficiency heuristic. Routes should have one or more steps, like "Take [line] from [station] to [station]" or "Change to [line]". You may add other relevant information to the results.

# Quickstart

## Requirements

[golang](https://golang.org/dl/) - v1.15

Go on the root of the project and run the application using

`go run main.go`

The server is now running on **port :1323**

Then you can simply call the API endpoint with a curl command:

- **/get-itinerary**: Get travel paths

`curl -X GET "http://localhost:1323/get-itinerary?src=EW27&dest=DT12"`

- **/get-itinerary-time-cost**: Get travel paths with time cost computing

`curl -X GET "http://localhost:1323/get-itinerary-time-cost?src=EW27&dest=DT12&startTime=2018-01-20T19:35"`

# Tests

Run tests with `go test -v ./...`

Run

`go test -coverprofile=coverage.out`

following by

`go tool cover -html=coverage.out`

to visualize the html coverage report.

Currently the coverage is:

| folders     | coverage |     |     |     |
| ----------- | -------- | --- | --- | --- |
| controllers | 78.73%   |     |     |     |
| core/models | 98%      |     |     |     |

# Algorithm overview

First, we need to build the graph including each station (node) and their connections.

The best way is to use an adjacent list, each station is a key value in a dict, the value is the reference to the other stations which it is connected to:

```
graph = {
NS25: [NS24, NS26] // City Hall
NS26: [NS25, NS26, EW14] // Raffles Place
EW14: [EW13, EW15] // Raffles Place
}
```

From the source station, we simply need to visit the connected stations, and from the connected stations, visit their connected stations. We make sure to mark the visited station to not revisited an already visited station.

In the end, we will have traverse all the MRT network and select the results, in other words, the itineraries which end with the destination station.
