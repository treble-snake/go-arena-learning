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
	fmt.Printf("%s\n\nversus\n\n%s\n\n", duel.fighter1, duel.fighter2)
}

// Start starts the fights
func (duel *Duel) Start() {
	fighter1, fighter2 := duel.fighter1, duel.fighter2

	startFight := func(first fighters.Fighter, second fighters.Fighter) {
		fmt.Printf("%s has the first move\n", first.GetName())
		go first.Fight()
		go second.Fight()
	}

	if fighter1.IsFasterThan(fighter2) {
		startFight(fighter1, fighter2)
	} else {
		startFight(fighter2, fighter1)
	}

	hasWinner := func() bool {
		return fighter1.IsDead() || fighter2.IsDead()
	}

	processDamage :=
		func(dealer fighters.Fighter, receiver fighters.Fighter, amount uint) {
			fmt.Printf(
				"%s deals %d damage to %s\n", dealer.GetName(), amount, receiver.GetName())
			receiver.TakeDamage(amount)
		}

	for !hasWinner() {
		select {
		case dmg := <-fighter1.AwaitStrikes():
			processDamage(fighter1, fighter2, dmg)
		case dmg := <-fighter2.AwaitStrikes():
			processDamage(fighter2, fighter1, dmg)
		}
	}

	determineWinner := func() string {
		firstFighterDead := fighter1.IsDead()
		if firstFighterDead && fighter2.IsDead() {
			return "Draw" // todo: is it possible?
		}

		if firstFighterDead {
			fighter2.StandDown()
			return fighter2.GetName()
		}

		fighter1.StandDown()
		return fighter1.GetName()
	}

	duel.endingChannel <- determineWinner()
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
