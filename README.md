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

`curl -X GET "http://localhost:1323/get-itinerary?src=EW15&dest=DT12"`

```
Travel from Tanjong Pagar to Little India

Best route: [EW15 EW14 EW13 EW12 DT14 DT13 DT12]
Total Station travelled: 7

Travel on EW line from Tanjong Pagar to Raffles Place
Travel on EW line from Raffles Place to City Hall
Travel on EW line from City Hall to Bugis
Change from EW line to DT line
Travel on DT line from Bugis to Rochor
Travel on DT line from Rochor to Little India

----

Alternative route: [EW15 EW14 EW13 NS25 NS24 NE6 NE7 DT12]
Total Station travelled: 8

Travel on EW line from Tanjong Pagar to Raffles Place
Travel on EW line from Raffles Place to City Hall
Change from EW line to NS line
Travel on NS line from City Hall to Dhoby Ghaut
Change from NS line to NE line
Travel on NE line from Dhoby Ghaut to Little India
Change from NE line to DT line

```

- **/get-itinerary-time-cost**: Get travel paths with time cost computing

`curl -X GET "http://localhost:1323/get-itinerary-time-cost?src=EW27&dest=DT12&startTime=2018-01-20T19:35"`

```
Best route: [EW27 EW26 EW25 EW24 EW23 EW22 EW21 CC22 CC21 CC20 CC19 DT9 DT10 DT11 DT12]
Total Travel time: 160

Travel on EW line from Boon Lay to Lakeside
Travel on EW line from Lakeside to Chinese Garden
Travel on EW line from Chinese Garden to Jurong East
Travel on EW line from Jurong East to Clementi
Travel on EW line from Clementi to Dover
Travel on EW line from Dover to Buona Vista
Change from EW line to CC line
Travel on CC line from Buona Vista to Holland Village
Travel on CC line from Holland Village to Farrer Road
Travel on CC line from Farrer Road to Botanic Gardens
Change from CC line to DT line
Travel on DT line from Botanic Gardens to Stevens
Travel on DT line from Stevens to Newton
Travel on DT line from Newton to Little India

----

Alternative route: [EW27 EW26 EW25 EW24 EW23 EW22 EW21 EW20 EW19 EW18 EW17 EW16 NE3 NE4 NE5 NE6 NE7 DT12]
Total Travel time: 198

Travel on EW line from Boon Lay to Lakeside
Travel on EW line from Lakeside to Chinese Garden
Travel on EW line from Chinese Garden to Jurong East
Travel on EW line from Jurong East to Clementi
Travel on EW line from Clementi to Dover
Travel on EW line from Dover to Buona Vista
Travel on EW line from Buona Vista to Commonwealth
Travel on EW line from Commonwealth to Queenstown
Travel on EW line from Queenstown to Redhill
Travel on EW line from Redhill to Tiong Bahru
Travel on EW line from Tiong Bahru to Outram Park
Change from EW line to NE line
Travel on NE line from Outram Park to Chinatown
Travel on NE line from Chinatown to Clarke Quay
Travel on NE line from Clarke Quay to Dhoby Ghaut
Travel on NE line from Dhoby Ghaut to Little India
Change from NE line to DT line

```

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
