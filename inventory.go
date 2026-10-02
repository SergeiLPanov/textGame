package main

import "fmt"

type ItemType string

const (
	ItemConsumable ItemType = "consumable"
	ItemQuestItem  ItemType = "quest"
	ItemPassive    ItemType = "passive"
	ItemWeapon     ItemType = "weapon"
	ItemArmor      ItemType = "armor"
	ItemTrophy     ItemType = "trophy"
)

type Item struct {
	id          string
	name        string
	description string
	itemType    ItemType
	value       int
}

var allItems = map[string]Item{
	"health_potion": {
		id:          "health_potion",
		name:        "🧪 Зелье здоровья",
		description: fmt.Sprintf("Восстанавливает %d HP", healthPotionRestore),
		itemType:    ItemConsumable,
		value:       healthPotionRestore,
	},
	"mana_potion": {
		id:          "mana_potion",
		name:        "💙 Зелье маны",
		description: fmt.Sprintf("Восстанавливает %d маны", manaPotionRestore),
		itemType:    ItemConsumable,
		value:       manaPotionRestore,
	},
	"amulet": {
		id:          "amulet",
		name:        "🔮 Тёмный амулет",
		description: "Старинный амулет с чёрным камнем. Кому-то нужен...",
		itemType:    ItemQuestItem,
		value:       0,
	},
	"spell_power_potion": {
		id:          "spell_power_potion",
		name:        "✨ Зелье силы заклинаний",
		description: fmt.Sprintf("Навсегда увеличивает урон заклинаний на %d%%", spellPowerBonus),
		itemType:    ItemPassive,
		value:       spellPowerBonus,
	},
	"iron_sword": {
		id:          "iron_sword",
		name:        "⚔️  Железный меч",
		description: fmt.Sprintf("Оружие. +%d к урону в бою", ironSwordDamageBonus),
		itemType:    ItemWeapon,
		value:       ironSwordDamageBonus,
	},
	"steel_sword": {
		id:          "steel_sword",
		name:        "⚔️  Стальной меч",
		description: fmt.Sprintf("Оружие. +%d к урону в бою", steelSwordDamageBonus),
		itemType:    ItemWeapon,
		value:       steelSwordDamageBonus,
	},
	"leather_armor": {
		id:          "leather_armor",
		name:        "🛡️  Кожаная броня",
		description: fmt.Sprintf("Броня. -%d к входящему урону", leatherArmorDefense),
		itemType:    ItemArmor,
		value:       leatherArmorDefense,
	},
	"chainmail_armor": {
		id:          "chainmail_armor",
		name:        "🛡️  Кольчуга",
		description: fmt.Sprintf("Броня. -%d к входящему урону", chainmailArmorDefense),
		itemType:    ItemArmor,
		value:       chainmailArmorDefense,
	},
	"wolf_fang": {
		id:          "wolf_fang",
		name:        "🦷 Клык волка",
		description: "Острый клык поверженного волка. Просто трофей.",
		itemType:    ItemTrophy,
	},
	"goblin_ear": {
		id:          "goblin_ear",
		name:        "👂 Ухо гоблина",
		description: "Трофей охотника на гоблинов.",
		itemType:    ItemTrophy,
	},
	"undead_ash": {
		id:          "undead_ash",
		name:        "💀 Прах нежити",
		description: "Серый прах, оставшийся после уничтожения нежити.",
		itemType:    ItemTrophy,
	},
	"kobold_eye": {
		id:          "kobold_eye",
		name:        "👁️  Глаз кобольда",
		description: "Мутно-жёлтый глаз, всё ещё слабо светится в темноте.",
		itemType:    ItemTrophy,
	},
	"necromancer_staff": {
		id:          "necromancer_staff",
		name:        "🦯 Посох некроманта",
		description: "Уникальный трофей. Посох, пропитанный тёмной магией павшего некроманта.",
		itemType:    ItemTrophy,
	},
	"kobold_king_crown": {
		id:          "kobold_king_crown",
		name:        "👑 Корона короля кобольдов",
		description: "Уникальный трофей. Грубая, но золотая корона повелителя кобольдов.",
		itemType:    ItemTrophy,
	},
	"guard_insignia": {
		id:          "guard_insignia",
		name:        "🎖️  Нашивка стража",
		description: "Нашивка с гербом замка, снятая с поверженного стража оружейной.",
		itemType:    ItemTrophy,
	},
}

type Inventory struct {
	items []Item
}

func newInventory() *Inventory {
	return &Inventory{items: []Item{}}
}

func (inv *Inventory) addItem(item Item) {
	inv.items = append(inv.items, item)
	fmt.Printf("📦 Получен предмет: %s\n", item.name)
}

func (inv *Inventory) removeItem(id string) bool {
	for i, item := range inv.items {
		if item.id == id {
			inv.items = append(inv.items[:i], inv.items[i+1:]...)
			return true
		}
	}
	return false
}

func (inv *Inventory) hasItem(id string) bool {
	for _, item := range inv.items {
		if item.id == id {
			return true
		}
	}
	return false
}

func (inv *Inventory) countItem(id string) int {
	count := 0
	for _, item := range inv.items {
		if item.id == id {
			count++
		}
	}
	return count
}

func (inv *Inventory) getConsumables() []Item {
	var result []Item
	for _, item := range inv.items {
		if item.itemType == ItemConsumable {
			result = append(result, item)
		}
	}
	return result
}

func (inv *Inventory) equippableItems() []Item {
	var result []Item
	for _, item := range inv.items {
		if item.itemType == ItemWeapon || item.itemType == ItemArmor {
			result = append(result, item)
		}
	}
	return result
}

func (inv *Inventory) print() {
	fmt.Println("\n╔═══════════════════════════════╗")
	fmt.Println("║         📦 ИНВЕНТАРЬ          ║")
	fmt.Println("╚═══════════════════════════════╝")

	if len(inv.items) == 0 {
		fmt.Println("  Инвентарь пуст")
		return
	}

	for i, item := range inv.items {
		fmt.Printf("  %d. %s\n", i+1, item.name)
		fmt.Printf("     %s\n", item.description)
	}
}
