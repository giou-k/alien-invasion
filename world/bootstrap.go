package world

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Map struct {
	Cities map[string]*City
	Aliens map[int]*Alien
}

type City struct {
	Uid int
	Name string
	Directions map[string]string
	Alien []*Alien
}

// Bootstrap creates the world X and its cities based on the input file worldxin.txt.
func (m *Map) Bootstrap() error {
	file, err := os.Open("./worldxin.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	m.Cities = make(map[string]*City)

	scanner := bufio.NewScanner(file)
	cityUid := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		s := strings.Split(currentLine, " ")

		// If there is an empty line just continue.
		if len(s) == 1 {
			continue
		}

		// Assume that the line of file must start with the name of the city.
		city := &City{
			Uid:  cityUid,
			Name: s[0],
		}
		city.Directions = make(map[string]string)

		// Check if city has directions to other cities.(should we assume that it has, even though they are not in the
		// input file?)
		if len(s) > 1 {
			for _, v := range s[1:] {
				directionTo := strings.Split(v, "=")
				if len(directionTo) != 2 {
					return errors.New("directions for city: " + m.Cities[s[0]].Name + "in file worldxin.txt have" +
						" wrong format")
				}
				city.Directions[directionTo[0]] = directionTo[1]
			}

			m.Cities[city.Name] = city
			cityUid++
		}

	}

	return nil
}
