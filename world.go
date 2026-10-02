package main

import "fmt"

func darkForest(player *Player) {
	renderLocation("forest", "🌲 ТЁМНЫЙ ЛЕС")
	describe("forest_enter")

	fmt.Println("1 - Исследовать опушку")
	fmt.Println("2 - Углубиться в чащу")
	fmt.Println("3 - Осмотреть поляну")
	fmt.Println("4 - Вернуться назад")
	fmt.Print("\nВыбор: ")

	switch readChoice() {
	case 1:
		describe("forest_explore")
		describe("wolf_appear")
		wolf := newWolf()
		victory := fight(player, wolf)
		if victory {
			describe("wolf_death")
		} else {
			describe("forest_retreat")
		}
	case 2:
		wolfDen(player)
	case 3:
		forestGlade(player)
	case 4:
		describe("forest_retreat")
	default:
		fmt.Println("Неизвестная команда")
	}
}

func wolfDen(player *Player) {
	describe("forest_den_enter")
	for i := 1; i <= 2; i++ {
		fmt.Printf("\n⚠️  Волк %d из 2 бросается на тебя!\n", i)
		if !fight(player, newWolf()) || !player.isAlive() {
			return
		}
	}
	describe("forest_den_clear")
}

func forestGlade(player *Player) {
	describe("forest_glade_enter")
	if player.forestGladeFound {
		fmt.Println("Здесь уже пусто — ты подобрал всё в прошлый раз.")
		return
	}
	player.forestGladeFound = true
	describe("forest_glade_found")
	player.addGold(forestGladeGoldReward)
}

func swamp(player *Player) {
	renderLocation("swamp", "🌿 ГНИЛОЕ БОЛОТО")
	describe("swamp_enter")

	fmt.Println("1 - Исследовать болото")
	fmt.Println("2 - Вернуться назад")

	choice := readChoice()

	switch choice {
	case 1:
		describe("swamp_explore")
		describe("goblin_appear")
		goblin := newGoblin()
		victory := fight(player, goblin)
		if victory {
			describe("goblin_death")
			if player.hasKey {
				describe("key_obtained")
			}
		} else {
			describe("swamp_retreat")
		}
	case 2:
		describe("swamp_retreat")
	default:
		fmt.Println("Неизвестная команда")
	}
}

func oldCastle(player *Player) {
	renderLocation("castle", "🏰 СТАРЫЙ ЗАМОК")
	describe("castle_enter")

	fmt.Println("1 - Войти в замок")
	fmt.Println("2 - Вернуться назад")
	fmt.Print("\nВыбор: ")

	choice := readChoice()
	if choice != 1 {
		fmt.Println("Ты решаешь не рисковать и возвращаешься.")
		return
	}

	if !player.hasKey {
		describe("castle_no_key")
		return
	}

	describe("castle_hall")
	fmt.Println("\n1 - 🗡️  Тронный зал")
	fmt.Println("2 - 🛡️  Оружейная")
	fmt.Println("3 - 💰 Сокровищница")
	fmt.Println("4 - Вернуться назад")
	fmt.Print("\nВыбор: ")

	switch readChoice() {
	case 1:
		throneHall(player)
	case 2:
		armory(player)
	case 3:
		treasury(player)
	case 4:
		fmt.Println("Ты покидаешь замок.")
	default:
		fmt.Println("Неизвестная команда")
	}
}

func throneHall(player *Player) {
	describe("throne_hall_enter")
	describe("knight_appear")
	knight := newKnight()
	victory := fight(player, knight)
	if victory {
		describe("knight_death")
	}
}

func armory(player *Player) {
	describe("castle_armory_enter")
	describe("armory_guard_appear")
	guard := newArmoryGuard()
	victory := fight(player, guard)
	if victory {
		describe("armory_guard_death")
	}
}

func treasury(player *Player) {
	describe("castle_treasury_enter")
	if player.treasuryLooted {
		describe("castle_treasury_looted")
		return
	}
	player.treasuryLooted = true
	describe("castle_treasury_found")
	player.addGold(treasuryGoldReward)
	player.inventory.addItem(allItems["health_potion"])
}

func graveyard(player *Player) {
	renderLocation("graveyard", "⚰️  СТАРОЕ КЛАДБИЩЕ")
	describe("graveyard_enter")

	fmt.Println("1 - ⚔️  Сразиться с нежитью")
	fmt.Println("2 - 🏛️  Войти в старый мавзолей")
	fmt.Println("3 - 🪦 Войти в фамильный склеп")
	fmt.Println("4 - Вернуться назад")

	choice := readChoice()

	switch choice {
	case 1:
		describe("graveyard_fight")
		fmt.Println("1 - Атаковать скелета")
		fmt.Println("2 - Атаковать зомби")
		fightChoice := readChoice()
		switch fightChoice {
		case 1:
			describe("skeleton_appear")
			fight(player, newSkeleton())
		case 2:
			describe("zombie_appear")
			fight(player, newZombie())
		}
	case 2:
		mausoleum(player)
	case 3:
		familyCrypt(player)
	case 4:
		fmt.Println("Ты покидаешь кладбище.")
	default:
		fmt.Println("Неизвестная команда")
	}
}

func mausoleum(player *Player) {
	describe("mausoleum_enter")

	if player.gold < necromancerRecommendedGold {
		describe("mausoleum_warning")
		fmt.Println("1 - Войти несмотря на риск")
		fmt.Println("2 - Вернуться и подготовиться")
		choice := readChoice()
		if choice != 1 {
			fmt.Println("Мудрое решение. Вернись когда будешь готов.")
			return
		}
	}

	describe("necromancer_appear")
	victory := fightNecromancer(player)
	if victory {
		describe("necromancer_death")
	}
}

func familyCrypt(player *Player) {
	describe("crypt_enter")

	if player.inventory.hasItem("amulet") {
		describe("soul_amulet")
		soul := newSoul()
		victory := fight(player, soul)
		if victory {
			describe("soul_peace")
			player.inventory.removeItem("amulet")
			player.applySpellPower(spellPowerBonus)
		}
		return
	}

	describe("soul_no_amulet")
	fight(player, newSoul())
}

func koboldCave(player *Player) {
	renderLocation("cave", "⛏️  ПЕЩЕРА КОБОЛЬДОВ")
	describe("cave_enter")

	fmt.Println("1 - Войти в пещеру")
	fmt.Println("2 - Вернуться назад")

	choice := readChoice()

	if choice != 1 {
		fmt.Println("Ты решаешь не рисковать и уходишь.")
		return
	}

	describe("cave_level1")
	if !fight(player, newKobold()) || !player.isAlive() {
		return
	}
	fmt.Println("\n⚠️  Из темноты выбегает ещё один кобольд!")
	if !fight(player, newKobold()) || !player.isAlive() {
		return
	}
	describe("cave_level1_clear")

	fmt.Println("\n1 - Спуститься на второй уровень")
	fmt.Println("2 - Выйти из пещеры")
	choice = readChoice()
	if choice != 1 {
		fmt.Println("Ты выходишь из пещеры.")
		return
	}

	describe("cave_level2")
	for i := 1; i <= 3; i++ {
		fmt.Printf("\n⚠️  Кобольд %d из 3 атакует!\n", i)
		if !fight(player, newKobold()) || !player.isAlive() {
			return
		}
	}
	describe("cave_level2_clear")

	fmt.Println("\n1 - Спуститься в тронный зал короля")
	fmt.Println("2 - Выйти из пещеры")
	choice = readChoice()
	if choice != 1 {
		fmt.Println("Ты выходишь из пещеры.")
		return
	}

	if player.gold < koboldKingRecommendedGold {
		describe("cave_king_warning")
		fmt.Println("1 - Идти несмотря на риск")
		fmt.Println("2 - Выйти и подготовиться")
		choice = readChoice()
		if choice != 1 {
			fmt.Println("Мудрое решение. Вернись когда будешь готов.")
			return
		}
	}

	describe("cave_level3")
	victory := fightKoboldKing(player)

	if victory && player.isAlive() {
		describe("cave_king_throne")
		fmt.Println("\n1 - Осмотреть покои короля")
		fmt.Println("2 - Уйти")
		choice = readChoice()

		if choice == 1 {
			describe("cave_king_loot")
			player.learnLightRay()
			describe("light_ray_learned")
		}
	}
}

func shop(player *Player) {
	describe("shop_enter")
	fmt.Printf("\n💰 Твоё золото: %d\n", player.gold)
	fmt.Println()
	fmt.Printf("1 - 🧪 Зелье здоровья (+%d HP) — %d золота\n", healthPotionRestore, healthPotionCost)
	fmt.Printf("2 - 💙 Зелье маны (+%d маны) — %d золота\n", manaPotionRestore, manaPotionCost)
	fmt.Printf("3 - ⚔️  Железный меч (+%d урона) — %d золота\n", ironSwordDamageBonus, ironSwordCost)
	fmt.Printf("4 - ⚔️  Стальной меч (+%d урона) — %d золота\n", steelSwordDamageBonus, steelSwordCost)
	fmt.Printf("5 - 🛡️  Кожаная броня (-%d урона от атак) — %d золота\n", leatherArmorDefense, leatherArmorCost)
	fmt.Printf("6 - 🛡️  Кольчуга (-%d урона от атак) — %d золота\n", chainmailArmorDefense, chainmailArmorCost)
	fmt.Println("7 - Выйти")

	choice := readChoice()

	switch choice {
	case 1:
		if player.gold < healthPotionCost {
			describe("shop_no_gold")
			return
		}
		describe("shop_buy_health")
		player.gold -= healthPotionCost
		player.heal(healthPotionRestore)

	case 2:
		if player.gold < manaPotionCost {
			describe("shop_no_gold")
			return
		}
		describe("shop_buy_mana")
		player.gold -= manaPotionCost
		player.addMana(manaPotionRestore)

	case 3:
		buyGear(player, "iron_sword", ironSwordCost, "shop_buy_sword")

	case 4:
		buyGear(player, "steel_sword", steelSwordCost, "")

	case 5:
		buyGear(player, "leather_armor", leatherArmorCost, "")

	case 6:
		buyGear(player, "chainmail_armor", chainmailArmorCost, "")

	case 7:
		fmt.Println("До свидания, путник!")
	default:
		fmt.Println("Неизвестная команда")
	}
}

func buyGear(player *Player, itemID string, cost int, descKey string) {
	if player.gold < cost {
		describe("shop_no_gold")
		return
	}
	if descKey != "" {
		describe(descKey)
	}
	player.gold -= cost
	player.inventory.addItem(allItems[itemID])
	fmt.Println("   Экипируй его в меню «⚔️  Экипировка»")
}
