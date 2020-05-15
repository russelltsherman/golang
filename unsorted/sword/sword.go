package sword

import (
	"fmt"
)

// Sword - the sword object
type Sword struct {
	name string // Important tip for RPG players: always name your swords!
}

// Damage - returns the damage dealt by this sword.
func (Sword) Damage() int {
	return 2
}

// String - implements fmt.Stringer for the Sword type.
func (s Sword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points of damage to opponents", s.name, s.Damage())
}

// EnchantedSword - the magic sword object which composes type Sword
type EnchantedSword struct {
	Sword // Embed the Sword type
}

// Damage returns the damage dealt by the enchanted sword.
func (EnchantedSword) Damage() int {
	return 42
}
