package stats

import (

)

type Stats struct {
	TotalHP int
	CurrentHP int
	BaseAC int
	Initiative int
	Speed int
	HitDice int
	Inspiration bool
	ProficiencyBonus int
}

func (s *Stats) IncreaseTotalHP(amount int) {
	// TODO: Does there need to be some checking here?
	s.TotalHP += amount
}

func (s *Stats) IncreaseCurrentHP(amount int) {
	if (s.TotalHP - s.CurrentHP) < amount {
		s.CurrentHP = c.TotalHP
	} else {
		s.TotalHP += amount
	}
}

func (s *Stats) DecreaseCurrentHP(amount int) {
	// Need to add loads of death stuff here
	s.CurrentHP -= amount
	if s.CurrentHP < 0 {
		s.CureentHP = 0
	}
}

func (s *Stats) AddToAC(amount int) {
	s.BaseAC += amount
}

func (s *Stats) AddInspiration() {
	s.Inspiration = true
}
