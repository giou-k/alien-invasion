package world

import (
	"math/rand"
	"time"
)

type Alien struct {
	Uid      int
	CityName string
	Dead     bool
	Trapped  bool
}

// InputAliens creates the aliens and places them randomly in a city. If a city has already two aliens, then we don't
// input any other alien.
func (m *Map) InputAliens(alienNum int) error {
	m.Aliens = make(map[int]*Alien)

	rand.Seed(time.Now().UnixNano())

	for alienUid := 0; alienUid < alienNum; alienUid++ {
		randCityUid := rand.Intn(len(m.Cities))

		for _, currentCity := range m.Cities {
			if currentCity.Uid == randCityUid {
				// Check if city has already two aliens.
				if len(currentCity.Alien) == 2 {
					continue
				}

				// Create new alien.
				alien := newAlien(alienUid, currentCity.Name)

				// Pass alien value to city
				m.setAlien(alien, currentCity.Name)

				// Pass alien value to Map.
				m.Aliens[alienUid] = alien

				// no need to go to next iterations
				break
			}
		}
	}
	return nil
}

func newAlien(i int, cityName string) *Alien {
	return &Alien{
		Uid:      i,
		CityName: cityName,
		Dead:     false,
		Trapped:  false,
	}
}

func (m *Map) setAlien(alien *Alien, cityName string) {
	m.Cities[cityName].Alien = append(m.Cities[cityName].Alien, alien)
}
