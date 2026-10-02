package main

import (
	"encoding/json"
	"os"
)

var saveFilePath = "save.json"

type saveData struct {
	Name      string `json:"name"`
	Level     int    `json:"level"`
	Exp       int    `json:"exp"`
	ExpToNext int    `json:"exp_to_next"`
	Health    int    `json:"health"`
	MaxHealth int    `json:"max_health"`
	Mana      int    `json:"mana"`
	MaxMana   int    `json:"max_mana"`
	Gold      int    `json:"gold"`

	BaseDamage        int      `json:"damage"`
	HasKey            bool     `json:"has_key"`
	SpellPower        int      `json:"spell_power"`
	HasLightRay       bool     `json:"has_light_ray"`
	Items             []string `json:"items"`
	Weapon            string   `json:"weapon,omitempty"`
	Armor             string   `json:"armor,omitempty"`
	QuestElderDone    bool     `json:"quest_elder_done,omitempty"`
	QuestHealerDone   bool     `json:"quest_healer_done,omitempty"`
	QuestHuntressDone bool     `json:"quest_huntress_done,omitempty"`
	TreasuryLooted    bool     `json:"treasury_looted,omitempty"`
	ForestGladeFound  bool     `json:"forest_glade_found,omitempty"`
}

func saveExists() bool {
	_, err := os.Stat(saveFilePath)
	return err == nil
}

func saveGame(p *Player) error {
	data := saveData{
		Name:              p.name,
		Level:             p.level,
		Exp:               p.exp,
		ExpToNext:         p.expToNext,
		Health:            p.health,
		MaxHealth:         p.maxHealth,
		Mana:              p.mana,
		MaxMana:           p.maxMana,
		Gold:              p.gold,
		BaseDamage:        p.baseDamage,
		HasKey:            p.hasKey,
		SpellPower:        p.spellPower,
		HasLightRay:       p.hasLightRay,
		QuestElderDone:    p.questElderDone,
		QuestHealerDone:   p.questHealerDone,
		QuestHuntressDone: p.questHuntressDone,
		TreasuryLooted:    p.treasuryLooted,
		ForestGladeFound:  p.forestGladeFound,
	}
	for _, item := range p.inventory.items {
		data.Items = append(data.Items, item.id)
	}
	if p.weapon != nil {
		data.Weapon = p.weapon.id
	}
	if p.armorItem != nil {
		data.Armor = p.armorItem.id
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(saveFilePath, bytes, 0644)
}

func loadGame() (*Player, error) {
	bytes, err := os.ReadFile(saveFilePath)
	if err != nil {
		return nil, err
	}

	var data saveData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	if data.Level == 0 {
		data.Level = playerStartLevel
		data.ExpToNext = levelExpBase
	}

	p := &Player{
		name:              data.Name,
		level:             data.Level,
		exp:               data.Exp,
		expToNext:         data.ExpToNext,
		health:            data.Health,
		maxHealth:         data.MaxHealth,
		mana:              data.Mana,
		maxMana:           data.MaxMana,
		gold:              data.Gold,
		baseDamage:        data.BaseDamage,
		hasKey:            data.HasKey,
		spellPower:        data.SpellPower,
		hasLightRay:       data.HasLightRay,
		questElderDone:    data.QuestElderDone,
		questHealerDone:   data.QuestHealerDone,
		questHuntressDone: data.QuestHuntressDone,
		treasuryLooted:    data.TreasuryLooted,
		forestGladeFound:  data.ForestGladeFound,
		inventory:         newInventory(),
	}
	for _, id := range data.Items {
		if item, ok := allItems[id]; ok {
			p.inventory.items = append(p.inventory.items, item)
		}
	}
	if data.Weapon != "" {
		if item, ok := allItems[data.Weapon]; ok {
			p.weapon = &item
		}
	}
	if data.Armor != "" {
		if item, ok := allItems[data.Armor]; ok {
			p.armorItem = &item
			p.armorDefense = item.value
		}
	}
	p.recomputeDamage()
	return p, nil
}
