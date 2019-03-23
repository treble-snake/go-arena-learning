package fighters

import (
	"math/rand"
	"time"
)

// WarriorBasics incapsulates every warrior basic attributes
type WarriorBasics struct {
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

func (warrior *WarriorBasics) AwaitStrikes() <-chan uint {
	return warrior.strikesChannel
}

func (warrior *WarriorBasics) GetName() string {
	return warrior.name
}

func (warrior *WarriorBasics) GetAgility() uint {
	return warrior.agility
}

func (warrior *WarriorBasics) IsDead() bool {
	return warrior.health <= 0
}

func (warrior *WarriorBasics) IsFasterThan(other Fighter) bool {
	otherAgility := other.GetAgility()
	if warrior.agility == otherAgility {
		return rand.Intn(2) > 0
	}
	return warrior.agility > otherAgility
}
