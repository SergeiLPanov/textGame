package main

import "testing"

func TestAddAndHasItem(t *testing.T) {
	inv := newInventory()
	inv.addItem(allItems["health_potion"])
	if !inv.hasItem("health_potion") {
		t.Fatal("inventory should contain the added item")
	}
	if inv.hasItem("mana_potion") {
		t.Fatal("inventory should not contain an item that was never added")
	}
}

func TestRemoveItem(t *testing.T) {
	inv := newInventory()
	inv.addItem(allItems["health_potion"])

	if !inv.removeItem("health_potion") {
		t.Fatal("removeItem should return true when the item exists")
	}
	if inv.hasItem("health_potion") {
		t.Fatal("item should be gone after removal")
	}
	if inv.removeItem("health_potion") {
		t.Fatal("removeItem should return false when the item no longer exists")
	}
}

func TestCountItem(t *testing.T) {
	inv := newInventory()
	if inv.countItem("wolf_fang") != 0 {
		t.Fatal("countItem should be 0 for an empty inventory")
	}

	inv.addItem(allItems["wolf_fang"])
	inv.addItem(allItems["wolf_fang"])
	inv.addItem(allItems["goblin_ear"])

	if got := inv.countItem("wolf_fang"); got != 2 {
		t.Fatalf("countItem(wolf_fang) = %d, want 2", got)
	}
	if got := inv.countItem("goblin_ear"); got != 1 {
		t.Fatalf("countItem(goblin_ear) = %d, want 1", got)
	}
}

func TestGetConsumablesFiltersByType(t *testing.T) {
	inv := newInventory()
	inv.addItem(allItems["health_potion"])
	inv.addItem(allItems["amulet"])
	inv.addItem(allItems["mana_potion"])

	consumables := inv.getConsumables()
	if len(consumables) != 2 {
		t.Fatalf("got %d consumables, want 2", len(consumables))
	}
	for _, item := range consumables {
		if item.itemType != ItemConsumable {
			t.Fatalf("non-consumable item %q returned by getConsumables", item.id)
		}
	}
}
