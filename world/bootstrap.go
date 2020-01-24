package world

type Map struct {
	Cities map[string]*City
	Aliens map[int]*Alien
}

type City struct {
	Uid int
	Name string
	North bool
	East bool
	South bool
	West bool
	Alien *Alien
}

func (m *Map) bootstrap()  {

}
