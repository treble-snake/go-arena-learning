package main

import (
	"fmt"
	"math/rand"
	"time"

	"fighting.pit/arena"
	"fighting.pit/fighters"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	thor := fighters.NewSimpleGrunt("Thor", 10, 5, 9)
	odin := fighters.NewSimpleGrunt("Odin", 8, 8, 11)

	duel := arena.NewDuel(thor, odin)
	duel.Announce()

	go duel.Start()
	winner := <-duel.AwaitEnding()

	fmt.Printf("\n%s won!\n", winner)
}
