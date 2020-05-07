package sword

import (
	"fmt"
	"testing"
)

type damageTest struct {
	sword interface {
		Damage() int
	}
	expected int
}

var damageTests = []damageTest{
	{
		sword:    Sword{name: "Silver Saber"},
		expected: 2,
	},
	{
		sword:    EnchantedSword{Sword{name: "Dragon's Greatsword"}},
		expected: 42,
	},
}

func TestDamage(t *testing.T) {
	for idx, spec := range damageTests {
		damage := spec.sword.Damage()
		if damage != spec.expected {
			t.Errorf("[spec %d] expected to get damage %d; got %d", idx, spec.expected, damage)
		}
	}
}

type stringTest struct {
	sword fmt.Stringer
	exp   string
}

var stringTests = []stringTest{
	{
		sword: Sword{name: "Silver Saber"},
		exp:   "Silver Saber is a sword that can deal 2 points of damage to opponents",
	},
	{
		sword: EnchantedSword{Sword{name: "Dragon's Greatsword"}},
		exp:   "Dragon's Greatsword is a sword that can deal 42 points of damage to opponents",
	},
}

func TestSwordToString(t *testing.T) {
	for idx, spec := range stringTests {
		got := spec.sword.String()
		if got != spec.exp {
			t.Errorf("[spec %d] expected to get\n%q\ngot:\n%q", idx, spec.exp, got)
		}
	}
}
