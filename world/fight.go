package world

// Fight implements a fight between two aliens. First the aliens turn dead and secondly the city gets deleted from the
// Map and the direction to this city from the other cities that exist in our Map.
func (m *Map) Fight(city *City) {
	m.killAliens(m.Cities[city.Name].Alien)

	m.removeCity(m.Cities[city.Name].Name)
}

func (m *Map) killAliens(aliens []*Alien) {
	for _, alien := range aliens {
		alien.Dead = true
	}
}

func (m *Map) removeCity(cityName string)  {
	// remove the city from world.
	delete(m.Cities, cityName)

	// Search through cities and delete removed city's directions.
	for _,city := range m.Cities {
		for _,directionCity := range city.Directions  {
			if directionCity ==  cityName {
				delete(city.Directions, cityName)
			}
		}

	}
}
