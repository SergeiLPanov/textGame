package main

import "testing"

func TestEquipWeaponIncreasesDamage(t *testing.T) {
	p := newPlayer("Hero")
	start := p.damage

	p.equipItem(allItems["iron_sword"])

	if p.weapon == nil || p.weapon.id != "iron_sword" {
		t.Fatalf("weapon = %+v, want iron_sword equipped", p.weapon)
	}
	want := start + ironSwordDamageBonus
	if p.damage != want {
		t.Fatalf("damage = %d, want %d", p.damage, want)
	}
}

func TestEquipArmorSetsArmorDefense(t *testing.T) {
	p := newPlayer("Hero")
	p.equipItem(allItems["leather_armor"])

	if p.armorItem == nil || p.armorItem.id != "leather_armor" {
		t.Fatalf("armorItem = %+v, want leather_armor equipped", p.armorItem)
	}
	if p.armorDefense != leatherArmorDefense {
		t.Fatalf("armorDefense = %d, want %d", p.armorDefense, leatherArmorDefense)
	}
}

func TestEquipWeaponSwapReturnsOldWeaponToInventory(t *testing.T) {
	p := newPlayer("Hero")
	p.equipItem(allItems["iron_sword"])
	p.equipItem(allItems["steel_sword"])

	if p.weapon == nil || p.weapon.id != "steel_sword" {
		t.Fatalf("weapon = %+v, want steel_sword equipped", p.weapon)
	}
	if !p.inventory.hasItem("iron_sword") {
		t.Fatal("swapping weapons should return the old one to the inventory")
	}
	want := playerStartDamage + steelSwordDamageBonus
	if p.damage != want {
		t.Fatalf("damage = %d, want %d", p.damage, want)
	}
}

func TestEquipItemRemovesItFromInventory(t *testing.T) {
	p := newPlayer("Hero")
	p.inventory.addItem(allItems["iron_sword"])

	p.equipItem(allItems["iron_sword"])

	if p.inventory.hasItem("iron_sword") {
		t.Fatal("equipped item should be removed from the loose inventory list")
	}
}

func TestEquipConsumableItemIsRejected(t *testing.T) {
	p := newPlayer("Hero")
	p.equipItem(allItems["health_potion"])

	if p.weapon != nil || p.armorItem != nil {
		t.Fatal("a consumable should not be equippable as gear")
	}
}

func TestLevelUpPreservesWeaponBonusOnTopOfBaseDamage(t *testing.T) {
	p := newPlayer("Hero")
	p.equipItem(allItems["iron_sword"])
	p.addExp(p.expToNext)

	want := playerStartDamage + levelDamageBonus + ironSwordDamageBonus
	if p.damage != want {
		t.Fatalf("damage after level-up = %d, want %d", p.damage, want)
	}
}

func TestEquippableItemsFiltersGearOnly(t *testing.T) {
	inv := newInventory()
	inv.addItem(allItems["health_potion"])
	inv.addItem(allItems["iron_sword"])
	inv.addItem(allItems["leather_armor"])
	inv.addItem(allItems["amulet"])

	gear := inv.equippableItems()
	if len(gear) != 2 {
		t.Fatalf("got %d equippable items, want 2", len(gear))
	}
	for _, item := range gear {
		if item.itemType != ItemWeapon && item.itemType != ItemArmor {
			t.Fatalf("non-gear item %q returned by equippableItems", item.id)
		}
	}
}
