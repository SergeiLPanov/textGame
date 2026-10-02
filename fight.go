package main

import (
	"fmt"
	"math/rand"
)

func rollDrop(chance float64, roll float64) bool {
	return roll < chance
}

func printFightMenu(player *Player, enemy *Enemy) {
	renderBars(player, enemy)
	fmt.Println("─────────────────────────────────")
	fmt.Println("1 - ⚔️  Атаковать")
	fmt.Println("2 - 🛡  Защищаться")
	fmt.Println("3 - 🏃 Отступить")
	fmt.Printf("4 - 🔥 Огненный шар (%d ур, %d MP)\n", fireballDamage, fireballManaCost)
	fmt.Printf("5 - ❄️  Ледяная стрела (%d ур, %d MP)\n", iceArrowDamage, iceArrowManaCost)
	fmt.Printf("6 - ⚡ Удар молнии (%d ур, %d MP)\n", lightningDamage, lightningManaCost)
	if player.hasLightRay {
		if enemy.isUndead() {
			fmt.Printf("7 - ☀️  Луч света (%d ур x%.1f по нежити, %d MP)\n",
				lightRayDamage, lightRayVulnMultiplier, lightRayManaCost)
		} else {
			fmt.Printf("7 - ☀️  Луч света (%d ур, %d MP)\n", lightRayDamage, lightRayManaCost)
		}
	}
	fmt.Println("8 - 🎒 Использовать предмет")
	fmt.Print("\nВыбор: ")
}

func handleSpell(player *Player, enemy *Enemy, spellType SpellType, damage int, manaCost int, descKey string) bool {
	describe(descKey)
	if !player.castSpell(enemy, spellType, damage, manaCost) {
		return false
	}
	return true
}

func fight(player *Player, enemy *Enemy) bool {
	return runBattle(player, enemy, nil)
}

func fightBoss(player *Player, boss *Enemy, onHalfHP func()) bool {
	return runBattle(player, boss, onHalfHP)
}

func runBattle(player *Player, enemy *Enemy, onHalfHP func()) bool {
	enemy.scaleToLevel(player.level)
	renderBattle(enemy)
	isBoss := onHalfHP != nil
	if !isBoss {
		fmt.Printf("\n⚠️  На тебя напал: %s!\n", enemy.name)
	}

	halfHPTriggered := false

	for player.isAlive() && enemy.isAlive() {
		if isBoss && !halfHPTriggered && enemy.health <= enemy.maxHealth/2 {
			halfHPTriggered = true
			onHalfHP()
			if !player.isAlive() {
				return false
			}
		}

		printFightMenu(player, enemy)
		choice := readChoice()

		enemyAttacks := true

		switch choice {
		case 1:
			enemy.takeDamage(player.damage)
			fmt.Printf("⚔️  Ты нанёс %d урона!\n", player.damage)
		case 2:
			fmt.Println("🛡 Ты принимаешь защитную стойку!")
			player.defense = true
		case 3:
			fmt.Println("🏃 Ты отступаешь!")
			player.takeDamage(enemy.damage / 2)
			fmt.Printf("💢 %s задевает тебя!\n", enemy.name)
			return false
		case 4:
			if !handleSpell(player, enemy, SpellFire, fireballDamage, fireballManaCost, "fireball_cast") {
				enemyAttacks = false
			}
		case 5:
			if !handleSpell(player, enemy, SpellIce, iceArrowDamage, iceArrowManaCost, "ice_arrow_cast") {
				enemyAttacks = false
			}
		case 6:
			if !handleSpell(player, enemy, SpellLightning, lightningDamage, lightningManaCost, "lightning_cast") {
				enemyAttacks = false
			}
		case 7:
			if !player.hasLightRay {
				fmt.Println("❓ Неизвестная команда")
				enemyAttacks = false
			} else if !handleSpell(player, enemy, SpellLight, lightRayDamage, lightRayManaCost, "light_ray_cast") {
				enemyAttacks = false
			}
		case 8:
			player.useItemInBattle()
			enemyAttacks = false
		default:
			fmt.Println("❓ Неизвестная команда")
			enemyAttacks = false
		}

		if !enemy.isAlive() {
			if isBoss {
				fmt.Printf("\n🏆 %s повержен!\n", enemy.name)
			} else {
				fmt.Printf("\n🏆 Ты победил %s!\n", enemy.name)
			}
			if enemy.drop > 0 {
				player.addGold(enemy.drop)
			}
			if enemy.expReward > 0 {
				player.addExp(enemy.expReward)
			}
			if enemy.dropAmulet {
				player.inventory.addItem(allItems["amulet"])
				describe("amulet_drop")
			}
			if enemy.materialDrop != "" && rollDrop(materialDropChance, rand.Float64()) {
				player.inventory.addItem(allItems[enemy.materialDrop])
			}
			if enemy.uniqueDrop != "" {
				fmt.Println("💎 С поверженного врага падает нечто особенное...")
				player.inventory.addItem(allItems[enemy.uniqueDrop])
			}

			if !isBoss {
				if enemy.dropKey {
					player.hasKey = true
					fmt.Println("🗝️  Ты подбираешь ключ от замка!")
				}
				if enemy.dropSpell == "light_ray" {
					describe("light_ray_found")
					player.learnLightRay()
				}
			}
			return true
		}

		if enemyAttacks {
			player.takeDamage(enemy.damage)
			fmt.Printf("💢 %s наносит тебе %d урона!\n", enemy.name, enemy.damage)
			if !player.isAlive() {
				describe("player_death")
				return false
			}
		}
	}
	return false
}

func fightNecromancer(player *Player) bool {
	necromancer := newNecromancer()
	return fightBoss(player, necromancer, func() {
		describe("necromancer_summon")
		skeleton := newSkeleton()
		fmt.Println("⚠️  Сначала разберись со скелетом!")
		fight(player, skeleton)
	})
}

func fightKoboldKing(player *Player) bool {
	king := newKoboldKing()
	describe("kobold_king_appear")

	victory := fightBoss(player, king, func() {
		describe("kobold_king_phase2")
		for i := 1; i <= 2; i++ {
			fmt.Printf("\n⚠️  Кобольд-страж %d атакует!\n", i)
			fight(player, newKobold())
		}
		describe("kobold_king_enraged")
	})

	if victory {
		describe("kobold_king_death")
	}
	return victory
}
