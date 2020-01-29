package world

import (
	"fmt"
	"math/rand"
	"time"
)

// Move implements the functionality for aliens to move to other cities.
// Firstly choose the city, then move the alien of the city, to another city based on the random direction(if exists).
// I assume that aliens move only to cities which have less than two aliens.
func (m *Map) Move(currentCity *City) {

	// random seed
	rand.Seed(time.Now().UnixNano())
	chooseDirection := rand.Intn(4)

	// Get the cities that have less than two aliens.
	peacefulCity := m.peacefulCities()

	switch chooseDirection { // 0=north, 1=east, etc. clockwise.
	case 0:
		// Check if the direction exists in the current city.
		if newCity, ok := currentCity.Directions["north"]; ok {
			// Check if this city has less than two aliens.
			if _, peace := peacefulCity.Cities[newCity]; peace {
				// Give to alien the name of the new city.
				currentCity.Alien[0].CityName = newCity

				// Pass the alien to the new city.
				m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])

				fmt.Println("\nAlien ", currentCity.Alien[0].Uid, "moved to city ", newCity)

				// Old City should not be invaded anymore by this alien anymore.
				currentCity.Alien = currentCity.Alien[1:]
			} else {
				fmt.Println("\nCity " + newCity + " has already two aliens.")
			}
		}
	case 1:
		// Check if the direction exists in the current city.
		if newCity, ok := currentCity.Directions["east"]; ok {
			// Check if this city has less than two aliens.
			if _, peace := peacefulCity.Cities[newCity]; peace {
				// Give to alien the name of the new city.
				currentCity.Alien[0].CityName = newCity

				// Pass the alien to the new city.
				m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])

				fmt.Println("\nAlien ", currentCity.Alien[0].Uid, "moved to city ", newCity)

				// Old City should not be invaded anymore by this alien anymore.
				currentCity.Alien = currentCity.Alien[1:]
			} else {
				fmt.Println("\nCity " + newCity + " has already two aliens.")
			}
		}
	case 2:
		// Check if the direction exists in the current city.
		if newCity, ok := currentCity.Directions["south"]; ok {
			// Check if this city has less than two aliens.
			if _, peace := peacefulCity.Cities[newCity]; peace {
				// Give to alien the name of the new city.
				currentCity.Alien[0].CityName = newCity

				// Pass the alien to the new city.
				m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])

				fmt.Println("\nAlien ", currentCity.Alien[0].Uid, "moved to city ", newCity)

				// Old City should not be invaded anymore by this alien anymore.
				currentCity.Alien = currentCity.Alien[1:]
			} else {
				fmt.Println("\nCity " + newCity + " has already two aliens.")
			}
		}
	case 3:
		// Check if the direction exists in the current city.F
		if newCity, ok := currentCity.Directions["west"]; ok {
			// Check if this city has less than two aliens.
			if _, peace := peacefulCity.Cities[newCity]; peace {
				// Give to alien the name of the new city.
				currentCity.Alien[0].CityName = newCity

				// Pass the alien to the new city.
				m.Cities[newCity].Alien = append(m.Cities[newCity].Alien, currentCity.Alien[0])

				fmt.Println("\nAlien ", currentCity.Alien[0].Uid, "moved to city ", newCity)

				// Old City should not be invaded anymore by this alien anymore.
				currentCity.Alien = currentCity.Alien[1:]
			} else {
				fmt.Println("\nCity " + newCity + " has already two aliens.")
			}
		}
	}
	return
}

// peacefulCities returns cities that have less than two aliens.
func (m *Map) peacefulCities() Map {
	var peacefulCities Map

	peacefulCities.Cities = make(map[string]*City)

	for _, city := range m.Cities {
		if len(city.Alien) < 2 {
			peacefulCities.Cities[city.Name] = city
		}
	}

	return peacefulCities
}
