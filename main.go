package main

import (
	"fmt"
	"math/rand"
)

func main() {

	player := NewRandomCharacter("Gordon")
	enemy := NewRandomCharacter("Jax")

	fmt.Printf("Player: %s\n", player.Name)
	fmt.Printf("Health: %d\n", player.Stats.Health)
	fmt.Printf("Attack: %d\n", player.Stats.Attack)
	fmt.Printf("Defense: %d\n", player.Stats.Defense)

	fmt.Printf("Enemy: %s\n", enemy.Name)
	fmt.Printf("Health: %d\n", enemy.Stats.Health)
	fmt.Printf("Attack: %d\n", enemy.Stats.Attack)
	fmt.Printf("Defense: %d\n", enemy.Stats.Defense)
}

// CharacterStats is a struct that represents the stats of a character.
// For now these are just health, attack, and defense.
type CharacterStats struct {
	Health struct {
		Max     int
		Current int
	}
	Attack  int
	Defense int
	Level   int
}

// NewCharacterStats creates a new CharacterStats struct with the given values.
func NewCharacterStats(health, attack, defense int) CharacterStats {
	return CharacterStats{
		Health:  struct{ Max, Current int }{Max: health, Current: health},
		Attack:  attack,
		Defense: defense,
		Level:   1,
	}
}

// Character is a struct that represents a character in a game.
// The can have a collection of stats.
type Character struct {
	Name  string
	Stats CharacterStats
}

// NewCharacter creates a new character with the given name and stats.
func NewCharacter(name string, health, attack, defense int) *Character {
	return &Character{
		Name:  name,
		Stats: NewCharacterStats(health, attack, defense),
	}
}

// randomNumberBias generates a random number between 1 and max with a bias towards the bias value.
// The influence value determines how much the bias affects the result.
// The closer to 1, the more likely the result will be the bias value.
func randomIntnBias(max int, bias, influence float64) int {
	min := 1
	rnd := float64(rand.Intn(max-min) + min)
	mix := rand.Float64() * influence
	value := rnd*(1-mix) + bias*mix
	return int(value)
}

// NewRandomCharacter creates a new character with random stats.
func NewRandomCharacter(name string) *Character {
	health := randomIntnBias(100, 100, 1)
	attack := randomIntnBias(10, 5, 1)
	defense := randomIntnBias(5, 2, 1)
	return NewCharacter(name, health, attack, defense)
}

// Attack calculates the damage dealt by the character.
// It does not modify the target character.
func (c *Character) Attack(target *Character) int {
	damage := c.Stats.Attack - target.Stats.Defense
	if damage < 0 {
		damage = 0
	}
	return damage
}

// AttackEvent is a struct that represents an attack event.
// In a Battle we will store a series of AttackEvents and source the events every time we want to display
type BattleCombatLogEntry struct {
	Attacker *Character
	Target   *Character
	Damage   int
}

// BattleCombatLog is a slice of BattleCombatLogEntry structs, They are ordered by the time they were added to the log.
type BattleCombatLog []BattleCombatLogEntry

type BattleCharacter struct {
	Character *Character
	Health    int //The current health of the character in the battle
}

// Battle is a struct that represents a battle between two characters.
type Battle struct {
	Subject   BattleCharacter //The character that initiated the battle
	Opponent  BattleCharacter //The character that is being attacked
	CombatLog BattleCombatLog //The log of all attacks that have been made in the battle
}

// NewBattle creates a new battle between two characters and initializes the combat log.
func NewBattle(subject, opponent *Character) *Battle {
	return &Battle{
		Subject:   BattleCharacter{Character: subject, Health: subject.Stats.Health.Current},
		Opponent:  BattleCharacter{Character: opponent, Health: opponent.Stats.Health.Current},
		CombatLog: make(BattleCombatLog, 0),
	}
}

func (b *Battle) Winner() *Character {
	switch {
	case b.Subject.Health > 0 && b.Opponent.Health == 0:
		return b.Subject.Character //The subject wins
	case b.Subject.Health == 0 && b.Opponent.Health > 0:
		return b.Opponent.Character //The opponent wins
	default:
		return nil //There is no current winner
	}
}

func (b *Battle) Attack(attacker, target *Character) {
	damage := attacker.Attack(target)
	target.Stats.Health.Current -= damage
	b.CombatLog = append(b.CombatLog, BattleCombatLogEntry{Attacker: attacker, Target: target, Damage: damage})
}

// Just ideas?

// Simulate is the main function that runs the battle between the two characters.
// THe battle will step through turns, a character will attack the other via a new entry in the combat log.
// After the turn, we will calculate the health of the characters.
// The battle will continue until one of the characters has no health left.
// func (b *Battle) Battle()

// Finalize is will end a battle
// func (b *Battle) Finalize()
