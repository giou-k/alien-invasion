package main

import (
	"flag"
	"fmt"
	"github.com/cosmos/alien-invasion/world"
	"log"
)

func main() {
	fmt.Println("Invasion Starts!Get Covered!")

	var (
		alienNum int
		m world.Map
	)

	// Max number of aliens must be the total number of cities.
	flag.IntVar(&alienNum, "n", 3, "number of aliens. default: 3")
	flag.Parse()

	if alienNum == 0 || alienNum > 3 {
		log.Fatal("Invalid number of aliens.The number of aliens can't be 0 or greater than the number of cities.")
	}

	// Firstly bootstrap the world.
	err := m.Bootstrap()
	if err != nil {
		log.Fatal("error in bootstrap of the world: ",err)
		return
	}

	// Secondly input the aliens randomly.
	err = m.InputAliens(alienNum)
	if err != nil {
		log.Fatal("error while matching aliens with cities: ", err)
	}

	// Thirdly start the game.
	for _, currentCity := range m.Cities {
		switch len(currentCity.Alien) {
		case 0:
			continue
		case 1:
			continue // Move aliens
		case 2:
			m.Fight(currentCity)
		}
	}

	// Lastly print the outcome.
}
