package main

import (
	"os"
	"path/filepath"
	"testing"
)

func withTempSaveFile(t *testing.T) {
	t.Helper()
	original := saveFilePath
	saveFilePath = filepath.Join(t.TempDir(), "save.json")
	t.Cleanup(func() { saveFilePath = original })
}

func TestSaveExistsFalseWhenNoFile(t *testing.T) {
	withTempSaveFile(t)
	if saveExists() {
		t.Fatal("saveExists should be false before any save")
	}
}

func TestSaveAndLoadRoundTrip(t *testing.T) {
	withTempSaveFile(t)

	p := newPlayer("Аванюрист")
	p.level = 3
	p.exp = 40
	p.expToNext = 112
	p.health = 42
	p.mana = 10
	p.gold = 77
	p.baseDamage = 20
	p.hasKey = true
	p.spellPower = 25
	p.hasLightRay = true
	p.questElderDone = true
	p.questHuntressDone = true
	p.inventory.addItem(allItems["health_potion"])
	p.inventory.addItem(allItems["amulet"])
	p.inventory.addItem(allItems["iron_sword"])
	p.inventory.addItem(allItems["leather_armor"])
	p.equipItem(allItems["iron_sword"])
	p.equipItem(allItems["leather_armor"])

	if err := saveGame(p); err != nil {
		t.Fatalf("saveGame failed: %v", err)
	}
	if !saveExists() {
		t.Fatal("saveExists should be true after saving")
	}

	loaded, err := loadGame()
	if err != nil {
		t.Fatalf("loadGame failed: %v", err)
	}

	if loaded.name != p.name || loaded.level != p.level || loaded.exp != p.exp ||
		loaded.expToNext != p.expToNext || loaded.health != p.health || loaded.maxHealth != p.maxHealth ||
		loaded.mana != p.mana || loaded.maxMana != p.maxMana || loaded.gold != p.gold ||
		loaded.baseDamage != p.baseDamage || loaded.damage != p.damage ||
		loaded.armorDefense != p.armorDefense || loaded.hasKey != p.hasKey ||
		loaded.spellPower != p.spellPower || loaded.hasLightRay != p.hasLightRay ||
		loaded.questElderDone != p.questElderDone || loaded.questHealerDone != p.questHealerDone ||
		loaded.questHuntressDone != p.questHuntressDone {
		t.Fatalf("loaded player = %+v, want fields matching %+v", loaded, p)
	}

	if loaded.weapon == nil || loaded.weapon.id != "iron_sword" {
		t.Fatalf("loaded weapon = %+v, want iron_sword", loaded.weapon)
	}
	if loaded.armorItem == nil || loaded.armorItem.id != "leather_armor" {
		t.Fatalf("loaded armor = %+v, want leather_armor", loaded.armorItem)
	}

	if len(loaded.inventory.items) != 2 {
		t.Fatalf("loaded inventory has %d items, want 2", len(loaded.inventory.items))
	}
	if !loaded.inventory.hasItem("health_potion") || !loaded.inventory.hasItem("amulet") {
		t.Fatal("loaded inventory missing expected items")
	}
}

func TestLoadGameFailsWithoutSaveFile(t *testing.T) {
	withTempSaveFile(t)
	if _, err := loadGame(); err == nil {
		t.Fatal("loadGame should fail when no save file exists")
	}
}

func TestLoadGameMigratesPreLevelingSave(t *testing.T) {
	withTempSaveFile(t)

	legacy := `{"name":"Ветеран","health":80,"max_health":100,"mana":30,"max_mana":50,"gold":55,"damage":18}`
	if err := os.WriteFile(saveFilePath, []byte(legacy), 0644); err != nil {
		t.Fatalf("failed to write legacy save fixture: %v", err)
	}

	loaded, err := loadGame()
	if err != nil {
		t.Fatalf("loadGame failed on legacy save: %v", err)
	}
	if loaded.level != playerStartLevel {
		t.Fatalf("level = %d, want %d", loaded.level, playerStartLevel)
	}
	if loaded.expToNext != levelExpBase {
		t.Fatalf("expToNext = %d, want %d", loaded.expToNext, levelExpBase)
	}
	if loaded.exp != 0 {
		t.Fatalf("exp = %d, want 0", loaded.exp)
	}

	if loaded.baseDamage != 18 || loaded.damage != 18 {
		t.Fatalf("baseDamage/damage = %d/%d, want 18/18", loaded.baseDamage, loaded.damage)
	}
}
