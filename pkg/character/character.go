package character

import (
	"github.com/brownlow2/gohttp/pkg/overview"
	"github.com/brownlow2/gohttp/pkg/stats"
	"github.com/brownlow2/gohttp/pkg/skills"
)

type Character struct {
	overview.Overview `json:"overview"`
	stats.Stats `json:"stats"`
	skills.Skills `json:"skill"`
}
