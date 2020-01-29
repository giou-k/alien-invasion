# Alien Invasion

----
Alien Invasion is a exercise, done using golang. Mad aliens are about to invade the earth and you are tasked with 
simulating the invasion. 
                                                 You are given a map containing the names of cities in the non-existent
                                                  world of X. The map is in a file, with one city per line. The city
                                                   name is first, followed by 1-4 directions (north, south, east, or 
                                                   west). Each one represents a road to another city that lies in that
                                                    direction. 
When two aliens end up in the same place, they fight, and in the process kill each other and destroy the city. When a
 city is destroyed, it is removed from the map, and so are any roads that lead into or out of it.Once a city is 
 destroyed, aliens can no longer travel to or through it. This may lead to aliens getting "trapped". 

----

### Requirements
* [Golang installed](https://golang.org/doc/install)

### How to run
* `git clone https://github.com/giou-k/alien-invasion.git`
* `cd alien-invasion`
* `go run main.go -n=<number of aliens>, default: n=16`

### Assumptions
* Assumed that the input file has fixed format, like that the first word is the name of the city and after that we have
 the possible directions. Also that all the possible combinations of the directions between the cities is inserted in 
 input file.
 * Assumed that the max number of aliens in a city is two.
 * Assumed that the max number of aliens inserted in the game via a command line argument is double the number of the
  cities.
