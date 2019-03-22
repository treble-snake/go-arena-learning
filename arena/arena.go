package arena

import (
	"fmt"

	"fighting.pit/fighters"
)

// Duel is used for battle of 2 warriors
type Duel struct {
	fighter1      fighters.Fighter
	fighter2      fighters.Fighter
	endingChannel chan string
}

// Announce prints out duel description
func (duel *Duel) Announce() {
	fmt.Println(duel.fighter1, "\nversus\n", duel.fighter2, "\n")
}

// Start starts the fights
func (duel *Duel) Start() {
	fighter1, fighter2 := duel.fighter1, duel.fighter2

	// todo: fighter with highest agility starts the fight (or random)
	go fighter1.Fight()
	go fighter2.Fight()

	(func() {
		hasWinner := func() bool {
			return fighter1.IsDead() || fighter2.IsDead()
		}

		// returns true if the fight is over
		processDamage := func(
			dealer fighters.Fighter, receiver fighters.Fighter, amount uint) bool {
			fmt.Println(
				dealer.GetName(), "deals", amount, "damage to", receiver.GetName())
			receiver.TakeDamage(amount)
			return hasWinner()
		}

		for {
			select {
			case dmg := <-fighter1.AwaitStrikes():
				if processDamage(fighter1, fighter2, dmg) {
					return
				}
			case dmg := <-fighter2.AwaitStrikes():
				if processDamage(fighter2, fighter1, dmg) {
					return
				}
			}
		}
	})()

	var winner string
	if fighter1.IsDead() {
		fighter2.StandDown()
		winner = fighter2.GetName()
	} else {
		fighter1.StandDown()
		winner = fighter1.GetName()
	}
	duel.endingChannel <- winner
}

// AwaitEnding return channel in which winner's name will be sent
func (duel *Duel) AwaitEnding() <-chan string {
	return duel.endingChannel
}

// NewDuel creates a Duel between two fighters
func NewDuel(fighter1 fighters.Fighter, fighter2 fighters.Fighter) *Duel {
	// todo: validate unique names
	return &Duel{fighter1, fighter2, make(chan string)}
}
