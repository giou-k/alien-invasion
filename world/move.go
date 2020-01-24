package world

import (
	"math/rand"
	"time"
)

// Move implements the functionality for aliens to move to other cities.
// Firstly choose the city, then move the alien of the city, to another city based on the random
// direction(if exists).
func (m *Map) Move(currentCity *City) {
	// random seed
	rand.Seed(time.Now().UnixNano())
	chooseDirection := rand.Intn(len(currentCity.Directions))

	switch chooseDirection { // 0=north, 1=east, etc. clockwise.
	case 0:
		// Check if the direction exists in the current city.
		if newCity, ok := currentCity.Directions["north"]; ok {
			currentCity.Alien[0].CityName = newCity
			// Pass the alien to the new city.
			m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])
			// Old City should not be invaded anymore by this alien anymore.
			currentCity.Alien[0] = nil
		}
	case 1:
		if newCity, ok := currentCity.Directions["east"]; ok {
			currentCity.Alien[0].CityName = newCity
			m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])
			// Old City should not be invaded by this alien anymore.
			currentCity.Alien[0] = nil
		}
	case 2:
		if newCity, ok := currentCity.Directions["south"]; ok {
			currentCity.Alien[0].CityName = newCity
			m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])
			// Old City should not be invaded by this alien anymore.
			currentCity.Alien[0] = nil
		}
	case 3:
		if newCity, ok := currentCity.Directions["west"]; ok {
			currentCity.Alien[0].CityName = newCity
			m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])
			// Old City should not be invaded by this alien anymore.
			currentCity.Alien[0] = nil
		}
	}
	return
}
