package main

import "testing"

func TestTakeDamageReducesHealth(t *testing.T) {
	p := newPlayer("Hero")
	p.takeDamage(30)
	if p.health != playerStartHealth-30 {
		t.Fatalf("health = %d, want %d", p.health, playerStartHealth-30)
	}
}

func TestTakeDamageNeverGoesNegative(t *testing.T) {
	p := newPlayer("Hero")
	p.takeDamage(p.maxHealth * 10)
	if p.health != 0 {
		t.Fatalf("health = %d, want 0", p.health)
	}
}

func TestTakeDamageWhileDefendingIsHalvedAndConsumed(t *testing.T) {
	p := newPlayer("Hero")
	p.defense = true
	p.takeDamage(20)
	want := playerStartHealth - int(20*defenseMultiplier)
	if p.health != want {
		t.Fatalf("health = %d, want %d", p.health, want)
	}
	if p.defense {
		t.Fatal("defense should be consumed after taking damage")
	}
}

func TestTakeDamageReducedByArmor(t *testing.T) {
	p := newPlayer("Hero")
	p.armorDefense = 4
	p.takeDamage(10)
	want := playerStartHealth - 6
	if p.health != want {
		t.Fatalf("health = %d, want %d", p.health, want)
	}
}

func TestTakeDamageArmorClampsAtZeroNotNegative(t *testing.T) {
	p := newPlayer("Hero")
	p.armorDefense = 100
	p.takeDamage(10)
	if p.health != playerStartHealth {
		t.Fatalf("health = %d, want %d (armor should fully absorb the hit, not heal)", p.health, playerStartHealth)
	}
}

func TestTakeDamageArmorThenBlockStack(t *testing.T) {
	p := newPlayer("Hero")
	p.armorDefense = 4
	p.defense = true
	p.takeDamage(10)
	want := playerStartHealth - 3
	if p.health != want {
		t.Fatalf("health = %d, want %d", p.health, want)
	}
}

func TestHealCapsAtMaxHealth(t *testing.T) {
	p := newPlayer("Hero")
	p.health = p.maxHealth - 5
	p.heal(100)
	if p.health != p.maxHealth {
		t.Fatalf("health = %d, want %d", p.health, p.maxHealth)
	}
}

func TestAddManaCapsAtMaxMana(t *testing.T) {
	p := newPlayer("Hero")
	p.mana = p.maxMana - 5
	p.addMana(100)
	if p.mana != p.maxMana {
		t.Fatalf("mana = %d, want %d", p.mana, p.maxMana)
	}
}

func TestCastSpellFailsWithoutEnoughMana(t *testing.T) {
	p := newPlayer("Hero")
	p.mana = 5
	enemy := newWolf()
	if p.castSpell(enemy, SpellFire, fireballDamage, fireballManaCost) {
		t.Fatal("castSpell should fail when mana is insufficient")
	}
	if p.mana != 5 {
		t.Fatalf("mana should be unchanged on failed cast, got %d", p.mana)
	}
	if enemy.health != enemy.maxHealth {
		t.Fatal("enemy should take no damage on failed cast")
	}
}

func TestCastSpellDeductsManaAndDamagesEnemy(t *testing.T) {
	p := newPlayer("Hero")
	enemy := newWolf()
	startMana := p.mana

	ok := p.castSpell(enemy, SpellFire, fireballDamage, fireballManaCost)
	if !ok {
		t.Fatal("castSpell should succeed with enough mana")
	}
	if p.mana != startMana-fireballManaCost {
		t.Fatalf("mana = %d, want %d", p.mana, startMana-fireballManaCost)
	}
	if enemy.health >= enemy.maxHealth {
		t.Fatal("enemy should have taken damage")
	}
}

func TestApplySpellPowerScalesSpellDamage(t *testing.T) {
	p := newPlayer("Hero")
	p.applySpellPower(spellPowerBonus)
	if p.spellPower != spellPowerBonus {
		t.Fatalf("spellPower = %d, want %d", p.spellPower, spellPowerBonus)
	}

	enemy := newKoboldKing()
	p.castSpell(enemy, SpellLightning, lightningDamage, lightningManaCost)

	baseDamage, bonus := lightningDamage, spellPowerBonus
	want := enemy.maxHealth - int(float64(baseDamage)*(1.0+float64(bonus)/100.0))
	if enemy.health != want {
		t.Fatalf("enemy health = %d, want %d", enemy.health, want)
	}
}

func TestIsAlive(t *testing.T) {
	p := newPlayer("Hero")
	if !p.isAlive() {
		t.Fatal("freshly created player should be alive")
	}
	p.health = 0
	if p.isAlive() {
		t.Fatal("player with 0 health should not be alive")
	}
}

func TestAddExpBelowThresholdDoesNotLevelUp(t *testing.T) {
	p := newPlayer("Hero")
	p.addExp(p.expToNext - 1)
	if p.level != playerStartLevel {
		t.Fatalf("level = %d, want %d", p.level, playerStartLevel)
	}
	if p.exp != p.expToNext-1 {
		t.Fatalf("exp = %d, want %d", p.exp, p.expToNext-1)
	}
}

func TestAddExpTriggersLevelUpAndBoostsStats(t *testing.T) {
	p := newPlayer("Hero")
	startMaxHealth, startMaxMana, startDamage := p.maxHealth, p.maxMana, p.damage
	threshold := p.expToNext

	p.addExp(threshold)

	if p.level != playerStartLevel+1 {
		t.Fatalf("level = %d, want %d", p.level, playerStartLevel+1)
	}
	if p.exp != 0 {
		t.Fatalf("exp = %d, want 0 (exactly consumed by level-up)", p.exp)
	}
	if p.maxHealth != startMaxHealth+levelHealthBonus {
		t.Fatalf("maxHealth = %d, want %d", p.maxHealth, startMaxHealth+levelHealthBonus)
	}
	if p.health != p.maxHealth {
		t.Fatal("leveling up should fully heal the player")
	}
	if p.maxMana != startMaxMana+levelManaBonus {
		t.Fatalf("maxMana = %d, want %d", p.maxMana, startMaxMana+levelManaBonus)
	}
	if p.mana != p.maxMana {
		t.Fatal("leveling up should fully restore mana")
	}
	if p.damage != startDamage+levelDamageBonus {
		t.Fatalf("damage = %d, want %d", p.damage, startDamage+levelDamageBonus)
	}
	wantExpToNext := int(float64(threshold) * levelExpGrowth)
	if p.expToNext != wantExpToNext {
		t.Fatalf("expToNext = %d, want %d", p.expToNext, wantExpToNext)
	}
}

func TestAddExpCanTriggerMultipleLevelUps(t *testing.T) {
	p := newPlayer("Hero")

	huge := p.expToNext + int(float64(p.expToNext)*levelExpGrowth)
	p.addExp(huge)

	if p.level != playerStartLevel+2 {
		t.Fatalf("level = %d, want %d", p.level, playerStartLevel+2)
	}
}

func TestAddExpIgnoresNonPositiveAmounts(t *testing.T) {
	p := newPlayer("Hero")
	p.addExp(0)
	p.addExp(-10)
	if p.exp != 0 || p.level != playerStartLevel {
		t.Fatalf("addExp with non-positive amount should be a no-op, got exp=%d level=%d", p.exp, p.level)
	}
}
