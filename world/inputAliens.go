package world

import (
	"math/rand"
	"time"
)

type Alien struct {
	Uid      int
	CityName string
	Trapped  bool
}

// InputAliens creates the aliens and places them randomly in a city. If a city has already two aliens, then we don't
// input any other alien.
func (m *Map) InputAliens(alienNum int) {
	m.Aliens = make(map[int]*Alien)

	rand.Seed(time.Now().UnixNano())

	for alienUid := 0; alienUid < alienNum; alienUid++ {
		randCityUid := rand.Intn(len(m.Cities))

		for _, currentCity := range m.Cities {
			if currentCity.Uid == randCityUid {
				// Check if city has already two aliens.
				if len(currentCity.Alien) == 2 {
					// If not a new random city uid is given, then this aliens will not invade any city.
					rand.Seed(time.Now().UnixNano())

					// The range of the random number must start from the cities that have not yet been used.
					randCityUid = rand.Intn(len(m.Cities)-currentCity.Uid+1) + currentCity.Uid
					continue
				}

				// Create new alien.
				alien := newAlien(alienUid, currentCity.Name)

				// Pass alien value to city.
				m.setAlien(alien, currentCity.Name)

				// Pass alien value to Map.
				m.Aliens[alien.Uid] = alien

				// no need to go to next iterations
				break
			}
		}
	}
	return
}

func newAlien(i int, cityName string) *Alien {
	return &Alien{
		Uid:      i,
		CityName: cityName,
		Trapped:  false,
	}
}

func (m *Map) setAlien(alien *Alien, cityName string) {
	m.Cities[cityName].Alien = append(m.Cities[cityName].Alien, alien)
}
