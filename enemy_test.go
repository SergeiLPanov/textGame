package main

import "testing"

func TestEnemyTakeDamageNeverGoesNegative(t *testing.T) {
	e := newWolf()
	e.takeDamage(e.maxHealth * 10)
	if e.health != 0 {
		t.Fatalf("health = %d, want 0", e.health)
	}
	if e.isAlive() {
		t.Fatal("enemy with 0 health should not be alive")
	}
}

func TestTakeSpellDamageAppliesResistance(t *testing.T) {
	e := newWolf()
	dealt := e.takeSpellDamage(100, SpellIce)
	if dealt != int(100*resistMultiplier) {
		t.Fatalf("dealt = %d, want %d", dealt, int(100*resistMultiplier))
	}
}

func TestTakeSpellDamageAppliesVulnerability(t *testing.T) {
	e := newWolf()
	dealt := e.takeSpellDamage(100, SpellFire)
	if dealt != int(100*vulnerableMultiplier) {
		t.Fatalf("dealt = %d, want %d", dealt, int(100*vulnerableMultiplier))
	}
}

func TestTakeSpellDamageDefaultsToNeutral(t *testing.T) {
	e := newWolf()
	dealt := e.takeSpellDamage(100, SpellLightning)
	if dealt != 100 {
		t.Fatalf("dealt = %d, want 100 (neutral multiplier)", dealt)
	}
}

func TestUndeadIsHighlyResistantExceptLightAndLightning(t *testing.T) {
	skeleton := newSkeleton()
	if !skeleton.isUndead() {
		t.Fatal("skeleton should be classified as undead")
	}

	fireDamage := skeleton.takeSpellDamage(100, SpellFire)
	if fireDamage != int(100*highResistMultiplier) {
		t.Fatalf("fire damage = %d, want %d", fireDamage, int(100*highResistMultiplier))
	}

	skeleton = newSkeleton()
	lightDamage := skeleton.takeSpellDamage(100, SpellLight)
	if lightDamage != int(100*lightRayVulnMultiplier) {
		t.Fatalf("light damage = %d, want %d", lightDamage, int(100*lightRayVulnMultiplier))
	}
}

func TestNonUndeadIsNotUndead(t *testing.T) {
	wolf := newWolf()
	if wolf.isUndead() {
		t.Fatal("wolf should not be classified as undead")
	}
}

func TestScaleToLevelOneIsNoOp(t *testing.T) {
	wolf := newWolf()
	wantHealth, wantDamage, wantDrop, wantExp := wolf.health, wolf.damage, wolf.drop, wolf.expReward

	wolf.scaleToLevel(1)

	if wolf.health != wantHealth || wolf.maxHealth != wantHealth {
		t.Fatalf("health/maxHealth = %d/%d, want unchanged %d", wolf.health, wolf.maxHealth, wantHealth)
	}
	if wolf.damage != wantDamage {
		t.Fatalf("damage = %d, want unchanged %d", wolf.damage, wantDamage)
	}
	if wolf.drop != wantDrop {
		t.Fatalf("drop = %d, want unchanged %d", wolf.drop, wantDrop)
	}
	if wolf.expReward != wantExp {
		t.Fatalf("expReward = %d, want unchanged %d", wolf.expReward, wantExp)
	}
}

func TestScaleToLevelIncreasesStatsProportionally(t *testing.T) {
	wolf := newWolf()
	baseHealth, baseDamage, baseDrop, baseExp := wolf.health, wolf.damage, wolf.drop, wolf.expReward

	level := 5
	wolf.scaleToLevel(level)

	multiplier := 1 + float64(level-1)*enemyLevelScaling
	wantHealth := int(float64(baseHealth) * multiplier)
	wantDamage := int(float64(baseDamage) * multiplier)
	wantDrop := int(float64(baseDrop) * multiplier)
	wantExp := int(float64(baseExp) * multiplier)

	if wolf.maxHealth != wantHealth || wolf.health != wantHealth {
		t.Fatalf("health/maxHealth = %d/%d, want %d", wolf.health, wolf.maxHealth, wantHealth)
	}
	if wolf.damage != wantDamage {
		t.Fatalf("damage = %d, want %d", wolf.damage, wantDamage)
	}
	if wolf.drop != wantDrop {
		t.Fatalf("drop = %d, want %d", wolf.drop, wantDrop)
	}
	if wolf.expReward != wantExp {
		t.Fatalf("expReward = %d, want %d", wolf.expReward, wantExp)
	}
}

func TestScaleToLevelHealsToFullScaledHealth(t *testing.T) {
	wolf := newWolf()
	wolf.takeDamage(20)
	wolf.scaleToLevel(3)

	if wolf.health != wolf.maxHealth {
		t.Fatalf("health = %d, want full maxHealth %d after scaling", wolf.health, wolf.maxHealth)
	}
}

func TestRegularEnemiesHaveExpectedMaterialDrop(t *testing.T) {
	cases := []struct {
		name   string
		enemy  *Enemy
		wantID string
	}{
		{"wolf", newWolf(), "wolf_fang"},
		{"goblin", newGoblin(), "goblin_ear"},
		{"skeleton", newSkeleton(), "undead_ash"},
		{"zombie", newZombie(), "undead_ash"},
		{"soul", newSoul(), "undead_ash"},
		{"kobold", newKobold(), "kobold_eye"},
		{"armory guard", newArmoryGuard(), "guard_insignia"},
	}
	for _, c := range cases {
		if c.enemy.materialDrop != c.wantID {
			t.Errorf("%s: materialDrop = %q, want %q", c.name, c.enemy.materialDrop, c.wantID)
		}
		if _, ok := allItems[c.wantID]; !ok {
			t.Errorf("%s: materialDrop id %q not present in allItems", c.name, c.wantID)
		}
	}
}

func TestKnightHasNoMaterialDrop(t *testing.T) {

	knight := newKnight()
	if knight.materialDrop != "" {
		t.Fatalf("materialDrop = %q, want empty (knight already drops the amulet)", knight.materialDrop)
	}
}

func TestBossesHaveUniqueDropAndNoMaterialDrop(t *testing.T) {
	cases := []struct {
		name   string
		enemy  *Enemy
		wantID string
	}{
		{"necromancer", newNecromancer(), "necromancer_staff"},
		{"kobold king", newKoboldKing(), "kobold_king_crown"},
	}
	for _, c := range cases {
		if c.enemy.uniqueDrop != c.wantID {
			t.Errorf("%s: uniqueDrop = %q, want %q", c.name, c.enemy.uniqueDrop, c.wantID)
		}
		if _, ok := allItems[c.wantID]; !ok {
			t.Errorf("%s: uniqueDrop id %q not present in allItems", c.name, c.wantID)
		}
		if c.enemy.materialDrop != "" {
			t.Errorf("%s: materialDrop = %q, want empty (bosses drop a unique item instead)", c.name, c.enemy.materialDrop)
		}
	}
}
