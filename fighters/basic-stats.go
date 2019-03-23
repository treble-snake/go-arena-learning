package fighters

import (
	"math/rand"
	"time"
)

// BasicStats incapsulates every warrior basic attributes
type BasicStats struct {
	// must be unieuq inside arena
	name           string
	strength       uint
	agility        uint
	vitality       uint
	health         int
	maxHealth      uint
	strikesChannel chan uint
	strikeTimer    *time.Timer
}

func (stats *BasicStats) AwaitStrikes() <-chan uint {
	return stats.strikesChannel
}

func (stats *BasicStats) GetName() string {
	return stats.name
}

func (stats *BasicStats) GetAgility() uint {
	return stats.agility
}

func (stats *BasicStats) IsDead() bool {
	return stats.health <= 0
}

func (stats *BasicStats) IsFasterThan(other Fighter) bool {
	otherAgility := other.GetAgility()
	if stats.agility == otherAgility {
		return rand.Intn(2) > 0
	}
	return stats.agility > otherAgility
}
