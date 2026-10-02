package main

import "testing"

func TestForestGladeGrantsGoldOnFirstVisitOnly(t *testing.T) {
	p := newPlayer("Hero")

	forestGlade(p)
	if !p.forestGladeFound {
		t.Fatal("forestGladeFound should be true after the first visit")
	}
	if p.gold != forestGladeGoldReward {
		t.Fatalf("gold = %d, want %d after first visit", p.gold, forestGladeGoldReward)
	}

	forestGlade(p)
	if p.gold != forestGladeGoldReward {
		t.Fatalf("gold = %d, want unchanged %d after second visit", p.gold, forestGladeGoldReward)
	}
}

func TestTreasuryGrantsRewardOnFirstVisitOnly(t *testing.T) {
	p := newPlayer("Hero")

	treasury(p)
	if !p.treasuryLooted {
		t.Fatal("treasuryLooted should be true after the first visit")
	}
	if p.gold != treasuryGoldReward {
		t.Fatalf("gold = %d, want %d after first visit", p.gold, treasuryGoldReward)
	}
	if !p.inventory.hasItem("health_potion") {
		t.Fatal("first treasury visit should grant a health potion")
	}
	potionsAfterFirst := p.inventory.countItem("health_potion")

	treasury(p)
	if p.gold != treasuryGoldReward {
		t.Fatalf("gold = %d, want unchanged %d after second visit", p.gold, treasuryGoldReward)
	}
	if got := p.inventory.countItem("health_potion"); got != potionsAfterFirst {
		t.Fatalf("second treasury visit should not grant another potion: got %d, want %d", got, potionsAfterFirst)
	}
}
