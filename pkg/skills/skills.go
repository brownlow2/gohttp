package skills

import (

)

const (
	STR = "strength"
	DEX = "dexterity"
	CON = "constitution"
	INT = "intelligence"
	WIS = "wisdom"
	CHA = "charisma"
)

type Skill struct {
	Name string `json:"name"`
	Skills `json:"skills"`
	SavingThrow int `json:"saving_throw"`
	Value int `json:"value"`
}

type Skills struct {
	Skills []string `json:"names"`
	SkillValues map[string]int `json:"map"`
}
