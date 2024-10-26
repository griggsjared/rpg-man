package main

import "fmt"

func main() {

	player := NewCharacter("Gordon", 100, 10, 5)
	enemy := NewCharacter("Zombie", 50, 5, 2)

	fmt.Printf("%s attacks %s\n", player.Name, enemy.Name)
}

// Character is a struct that represents a character in a game.
// The can have a name and stats like health, attack, and defense.
type Character struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

// NewCharacter creates a new character with the given name and stats.
func NewCharacter(name string, health, attack, defense int) *Character {
	return &Character{
		Name:    name,
		Health:  health,
		Attack:  attack,
		Defense: defense,
	}
}
