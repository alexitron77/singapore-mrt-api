## Introduction

## Quickstart

## Algorithm explanation

First, we need to build the graph including each station (node) and their connections.

The best way is to use an adjacent list, each station is a key value in a dict, the value is the reference to the other stations which it is connected to:

graph = {
NS25: [NS24, NS26] // City Hall
NS26: [NS25, NS26, EW14] // Raffles Place
EW14: [EW13, EW15] // Raffles Place
}
