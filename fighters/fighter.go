package fighters

// Fighter repesents a warrior for the arena
type Fighter interface {
	// Fight starts fighting process.
	Fight()

	// StandDown stops fighter from further actions.
	StandDown()

	// TakeDamage recieves an amount of damage the fighter's supposed
	// to receive. Handling may vary between implementations.
	TakeDamage(amount uint)

	// AwaitStrikes returns a channel where the fighters will send
	// his strikes to.
	AwaitStrikes() <-chan uint

	GetName() string

	// IsDead shows if the fighter is defeated.
	IsDead() bool
}
