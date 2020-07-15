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

type Skills struct {
	Skills []Skill `json:"skills"`
}

type Skill struct {
	SkillName string `json:"name"`
	Skills []string `json:"names"`
	SkillValues map[string]int `json:"map"`
	SavingThrow int `json:"saving_throw"`
	Value int `json:"value"`
}
