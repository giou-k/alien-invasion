package main

import (
	"flag"
	"fmt"
	"github.com/giou-k/alien-invasion/world"
	"log"
)

func main() {
	fmt.Println("Invasion Starts!Get Covered!")

	var (
		alienNum int
		m        world.Map
	)

	// Firstly bootstrap the world.
	err := m.Bootstrap()
	if err != nil {
		log.Fatal("error in bootstrap of the world: ", err)
		return
	}

	flag.IntVar(&alienNum, "n", 16, "number of aliens. default: 3")
	flag.Parse()

	if alienNum == 0 || alienNum > 2*len(m.Cities) {
		log.Fatal("Invalid number of aliens.The number of aliens can't be 0 or greater than the double" +
			" number of cities.")
	}

	// Secondly input the aliens randomly.
	m.InputAliens(alienNum)

	// Iterate through cities.
	// If a city has no aliens, then do nothing
	// If a city has one alien, move that alien to another city. The new city must have none or one alien.
	// If a city has two aliens, let them fight and destroy the city. Delete the city and the paths to or out of it.
	for i := 0; i < 100; i++ {
		for _, currentCity := range m.Cities {
			switch len(currentCity.Alien) {
			case 0:
				continue
			case 1:
				trapped := m.Trap(currentCity)
				if trapped {
					continue
				}
				m.Move(currentCity)
			case 2:
				m.Fight(currentCity)
			}
		}
		if len(m.Cities) == len(m.CitiesTrapped) {
			fmt.Println("Aliens are trapped inside cities. The End.")
			return
		}
		if len(m.Aliens) == 0 {
			fmt.Println("All aliens are dead. The End.")
			return
		}
	}

	// Lastly print the outcome.
}
