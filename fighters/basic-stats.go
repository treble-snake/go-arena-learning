package fighters

import (
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

func (grunt *BasicStats) AwaitStrikes() <-chan uint {
	return grunt.strikesChannel
}

func (grunt *BasicStats) GetName() string {
	return grunt.name
}

func (grunt *BasicStats) IsDead() bool {
	return grunt.health <= 0
}
