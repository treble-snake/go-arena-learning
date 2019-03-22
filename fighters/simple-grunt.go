package fighters

import (
	"fmt"
	"math/rand"
	"time"
)

func getMaxHealth(vitality uint) uint {
	return vitality * 10
}

// SimpleGrunt has one simple weapon and no armor
type SimpleGrunt struct {
	BasicStats
}

func (grunt SimpleGrunt) String() string {
	return fmt.Sprintf("Warrior (Grunt) %s:\nStr: %d\nAgl: %d\nVit: %d",
		grunt.name, grunt.strength, grunt.agility, grunt.vitality)
}

func (grunt *SimpleGrunt) TakeDamage(amount uint) {
	grunt.health -= int(amount)
	fmt.Printf("%s: %d/%d\n", grunt.name, grunt.health, grunt.maxHealth)
}

func (grunt *SimpleGrunt) getStrikeTime() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(50)+50-int(2*grunt.agility))
}

func (grunt *SimpleGrunt) Fight() {
	grunt.strikeTimer = time.NewTimer(grunt.getStrikeTime())
	for range grunt.strikeTimer.C {
		if !grunt.IsDead() {
			grunt.strikesChannel <- grunt.strength
			grunt.strikeTimer.Reset(grunt.getStrikeTime())
		} else {
			grunt.closeAll()
			return
		}
	}
}

func (grunt *SimpleGrunt) closeAll() {
	// TODO wtf ?
	if grunt.strikeTimer != nil && !grunt.strikeTimer.Stop() {
		<-grunt.strikeTimer.C
	}

	close(grunt.strikesChannel)
}

func (grunt *SimpleGrunt) StandDown() {
	if grunt.IsDead() {
		return
	}

	grunt.closeAll()
}

// NewSimpleGrunt creates an instance of this type of a fighter
func NewSimpleGrunt(name string, strength uint, agility uint, vitality uint) *SimpleGrunt {
	maxHealth := getMaxHealth(vitality)
	return &SimpleGrunt{
		BasicStats{
			name,
			strength,
			agility,
			vitality,
			int(maxHealth),
			maxHealth,
			make(chan uint),
			nil}}
}
