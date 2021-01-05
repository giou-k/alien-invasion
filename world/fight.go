package world

import "fmt"

// Fight implements a fight between two aliens. First the aliens get deleted from city's alien array and secondly
// the city gets deleted from the Map, plus the directions towards this city from other cities that exist in our Map.
func (m *Map) Fight(city *City) {
	fmt.Println("\nAlien ", m.Cities[city.Name].Alien[0].Uid, " is fighting with alien ",
		+m.Cities[city.Name].Alien[1].Uid, " inside city ", city.Name)

	m.killAliens(m.Cities[city.Name].Alien)

	m.removeCity(m.Cities[city.Name].Name)
	fmt.Println("City ", city.Name, " has been destroyed.")
}

func (m *Map) killAliens(aliens []*Alien) {
	for _, alien := range aliens {
		delete(m.Aliens, alien.Uid)
	}
}

func (m *Map) removeCity(cityName string) {
	// remove the city from world.
	delete(m.Cities, cityName)

	// Search through cities and delete removed city's directions.
	for _, city := range m.Cities {
		for key, directionCity := range city.Directions {
			if directionCity == cityName {
				delete(city.Directions, key)
			}
		}

	}
}
