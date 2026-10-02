package main

import "fmt"

type Player struct {
	name      string
	level     int
	exp       int
	expToNext int
	health    int
	maxHealth int
	mana      int
	maxMana   int
	gold      int

	baseDamage   int
	damage       int
	armorDefense int
	defense      bool
	hasKey       bool
	spellPower   int
	hasLightRay  bool
	inventory    *Inventory
	weapon       *Item
	armorItem    *Item

	questElderDone    bool
	questHealerDone   bool
	questHuntressDone bool

	treasuryLooted   bool
	forestGladeFound bool
}

func newPlayer(name string) *Player {
	return &Player{
		name:        name,
		level:       playerStartLevel,
		exp:         0,
		expToNext:   levelExpBase,
		health:      playerStartHealth,
		maxHealth:   playerStartHealth,
		mana:        playerStartMana,
		maxMana:     playerMaxMana,
		gold:        playerStartGold,
		baseDamage:  playerStartDamage,
		damage:      playerStartDamage,
		spellPower:  0,
		hasLightRay: false,
		inventory:   newInventory(),
	}
}

func (p *Player) takeDamage(damage int) {
	if p.armorDefense > 0 {
		damage -= p.armorDefense
		if damage < 0 {
			damage = 0
		}
	}
	if p.defense {
		damage = int(float64(damage) * defenseMultiplier)
		fmt.Printf("🛡 Ты защищаешься! Урон снижен до %d\n", damage)
		p.defense = false
	}
	p.health -= damage
	if p.health < 0 {
		p.health = 0
	}
}

func (p *Player) addGold(gold int) {
	p.gold += gold
	fmt.Printf("💰 Ты получил %d золота! Всего: %d\n", gold, p.gold)
}

func (p *Player) addExp(amount int) {
	if amount <= 0 {
		return
	}
	p.exp += amount
	fmt.Printf("⭐ Получено опыта: %d (%d/%d)\n", amount, p.exp, p.expToNext)

	for p.exp >= p.expToNext {
		p.exp -= p.expToNext
		p.levelUp()
	}
}

func (p *Player) levelUp() {
	p.level++
	p.maxHealth += levelHealthBonus
	p.health = p.maxHealth
	p.maxMana += levelManaBonus
	p.mana = p.maxMana
	p.baseDamage += levelDamageBonus
	p.recomputeDamage()
	p.expToNext = int(float64(p.expToNext) * levelExpGrowth)

	fmt.Println("\n✨═══════════════════════════✨")
	fmt.Printf("🎉 НОВЫЙ УРОВЕНЬ: %d!\n", p.level)
	fmt.Printf("❤️  Макс. HP: %d (+%d), восстановлено полностью\n", p.maxHealth, levelHealthBonus)
	fmt.Printf("💙 Макс. мана: %d (+%d), восстановлена полностью\n", p.maxMana, levelManaBonus)
	fmt.Printf("⚔️  Урон: %d (+%d)\n", p.damage, levelDamageBonus)
	fmt.Println("✨═══════════════════════════✨")
}

func (p *Player) addMana(mana int) {
	p.mana += mana
	if p.mana > p.maxMana {
		p.mana = p.maxMana
	}
	fmt.Printf("💙 Мана восстановлена! Мана: %d/%d\n", p.mana, p.maxMana)
}

func (p *Player) heal(amount int) {
	p.health += amount
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
	fmt.Printf("❤️  Здоровье восстановлено! HP: %d/%d\n", p.health, p.maxHealth)
}

func (p *Player) learnLightRay() {
	if p.hasLightRay {
		fmt.Println("Ты уже знаешь это заклинание!")
		return
	}
	p.hasLightRay = true
	fmt.Println("✨ Ты изучил заклинание ☀️  Луч света!")
	fmt.Println("   Наносит повышенный урон нежити!")
}

func (p *Player) castSpell(enemy *Enemy, spellType SpellType, baseDamage int, manaCost int) bool {
	if p.mana < manaCost {
		describe("no_mana")
		fmt.Printf("Нужно %d маны, у тебя %d\n", manaCost, p.mana)
		return false
	}
	p.mana -= manaCost

	damage := baseDamage
	if p.spellPower > 0 {
		damage = int(float64(baseDamage) * (1.0 + float64(p.spellPower)/100.0))
	}

	finalDamage := enemy.takeSpellDamage(damage, spellType)

	if enemy.isUndead() && spellType == SpellLight {
		fmt.Printf("☀️  СВЯЩЕННЫЙ УРОН ПО НЕЖИТИ! Урон: %d!\n", finalDamage)
	} else {
		switch {
		case float64(finalDamage) > float64(damage)*1.1:
			fmt.Printf("💥 УЯЗВИМОСТЬ! Урон: %d!\n", finalDamage)
		case float64(finalDamage) < float64(damage)*0.9:
			fmt.Printf("🛡 СОПРОТИВЛЕНИЕ! Урон: %d!\n", finalDamage)
		default:
			fmt.Printf("💫 Урон: %d\n", finalDamage)
		}
	}
	fmt.Printf("💙 Мана: %d/%d\n", p.mana, p.maxMana)
	return true
}

func (p *Player) useItemInBattle() bool {
	consumables := p.inventory.getConsumables()
	if len(consumables) == 0 {
		fmt.Println("❌ Нет расходуемых предметов!")
		return false
	}

	fmt.Println("\n🎒 Выбери предмет:")
	for i, item := range consumables {
		fmt.Printf("%d - %s (%s)\n", i+1, item.name, item.description)
	}
	fmt.Println("0 - Отмена")

	choice := readChoice()

	if choice == 0 || choice > len(consumables) {
		return false
	}

	item := consumables[choice-1]
	switch item.id {
	case "health_potion":
		p.heal(item.value)
		p.inventory.removeItem(item.id)
	case "mana_potion":
		p.addMana(item.value)
		p.inventory.removeItem(item.id)
	}
	return true
}

func (p *Player) applySpellPower(bonus int) {
	p.spellPower += bonus
	fmt.Printf("✨ Сила заклинаний увеличена на %d%%! Итого: +%d%%\n", bonus, p.spellPower)
}

func (p *Player) recomputeDamage() {
	p.damage = p.baseDamage
	if p.weapon != nil {
		p.damage += p.weapon.value
	}
}

func (p *Player) equipItem(item Item) {
	switch item.itemType {
	case ItemWeapon:
		if p.weapon != nil {
			p.inventory.items = append(p.inventory.items, *p.weapon)
		}
		equipped := item
		p.weapon = &equipped
		p.inventory.removeItem(item.id)
		p.recomputeDamage()
		fmt.Printf("⚔️  Экипировано: %s. Урон: %d\n", item.name, p.damage)
	case ItemArmor:
		if p.armorItem != nil {
			p.inventory.items = append(p.inventory.items, *p.armorItem)
		}
		equipped := item
		p.armorItem = &equipped
		p.armorDefense = item.value
		p.inventory.removeItem(item.id)
		fmt.Printf("🛡️  Надето: %s. Снижение урона: %d\n", item.name, p.armorDefense)
	default:
		fmt.Println("❌ Этот предмет нельзя экипировать.")
	}
}

func (p *Player) printStatus() {
	fmt.Println("\n╔═══════════════════════════════╗")
	fmt.Printf("║  ⚔️  Герой: %-18s║\n", p.name)
	fmt.Println("╠═══════════════════════════════╣")
	fmt.Printf("║  🎖️  Уровень: %d\n", p.level)
	fmt.Printf("║  ⭐ Опыт: %d/%d\n", p.exp, p.expToNext)
	fmt.Printf("║  ❤️  HP:   %d/%d\n", p.health, p.maxHealth)
	fmt.Printf("║  💙 Мана:  %d/%d\n", p.mana, p.maxMana)
	fmt.Printf("║  💰 Золото: %d\n", p.gold)
	fmt.Printf("║  ⚔️  Урон:  %d\n", p.damage)
	if p.weapon != nil {
		fmt.Printf("║      └─ оружие: %s (+%d)\n", p.weapon.name, p.weapon.value)
	}
	if p.armorItem != nil {
		fmt.Printf("║  🛡️  Броня: %s (-%d)\n", p.armorItem.name, p.armorItem.value)
	}
	if p.spellPower > 0 {
		fmt.Printf("║  ✨ Сила заклинаний: +%d%%\n", p.spellPower)
	}
	if p.hasKey {
		fmt.Println("║  🗝️  Ключ от замка: есть")
	}
	if p.hasLightRay {
		fmt.Println("║  ☀️  Луч света: изучен")
	}
	questsDone := 0
	for _, done := range []bool{p.questElderDone, p.questHealerDone, p.questHuntressDone} {
		if done {
			questsDone++
		}
	}
	if questsDone > 0 {
		fmt.Printf("║  🧝 Квесты деревни эльфов: %d/3\n", questsDone)
	}
	fmt.Println("╚═══════════════════════════════╝")
	p.inventory.print()
}

func (p *Player) isAlive() bool {
	return p.health > 0
}
