package atheletes

import "strings"

// Info ...
type Info struct {
	Country   string
	HairColor string
}

// Player ...
type Player struct {
	Name  string
	Sport string
	Age   int
	Info  Info
}

// ToLowerCase ...
func (p *Player) ToLowerCase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Info.Country = strings.ToLower(p.Info.Country)
	p.Info.HairColor = strings.ToLower(p.Info.HairColor)

	return p
}
