package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Invasion Starts!Get Covered!")

	var (
		alienNum int
		//m world.Map
	)

	// Max number of aliens must be the total number of cities.
	flag.IntVar(&alienNum, "n", 2, "number of aliens. default: 3")
	flag.Parse()

	if alienNum == 0 || alienNum > 2 {
		log.Fatal("Invalid number of aliens.The number of aliens can't be 0 or greater than the number of cities.")
	}

	// Firstly bootstrap the world.

	// Secondly input the aliens randomly.

	// Thirdly start the game.

	// Lastly print the outcome.
}
