package main

import (
	"fmt"
	"textGame/assets"
)

func getEnemySprite(name string) string {
	switch name {
	case "Волк":
		return assets.Wolf
	case "Гоблин":
		return assets.Goblin
	case "Тёмный рыцарь":
		return assets.Knight
	case "Страж оружейной":
		return assets.Knight
	case "Скелет":
		return assets.Skeleton
	case "Зомби":
		return assets.Zombie
	case "Неупокоенная душа":
		return assets.Soul
	case "Некромант":
		return assets.Necromancer
	case "Кобольд":
		return assets.Kobold
	case "Король кобольдов":
		return assets.KoboldKing
	default:
		return `  ⠀⣀⡀
  ⣿⣿⣿
  ⠀⣿⠀
  ⠀⣿⠀`
	}
}

func getLocationSprite(location string) string {
	switch location {
	case "forest":
		return assets.Forest
	case "swamp":
		return assets.Swamp
	case "castle":
		return assets.Castle
	case "graveyard":
		return assets.Graveyard
	case "cave":
		return assets.Cave
	case "elf_village":
		return assets.ElfVillage
	default:
		return ""
	}
}

func renderBattle(enemy *Enemy) {
	sprite := getEnemySprite(enemy.name)
	screen := assets.BattleScreen(assets.Hero, sprite, enemy.name)
	fmt.Println(screen)
}

func renderLocation(key, name string) {
	sprite := getLocationSprite(key)
	if sprite != "" {
		screen := assets.LocationScreen(sprite, name)
		fmt.Println(screen)
	}
}

func renderBars(player *Player, enemy *Enemy) {
	fmt.Println()
	fmt.Printf("🧙 %s\n", player.name)
	fmt.Printf("%s %d/%d\n", assets.HealthBar(player.health, player.maxHealth), player.health, player.maxHealth)
	fmt.Printf("%s %d/%d\n", assets.ManaBar(player.mana, player.maxMana), player.mana, player.maxMana)
	fmt.Println()
	fmt.Printf("👹 %s\n", enemy.name)
	fmt.Printf("%s %d/%d\n", assets.HealthBar(enemy.health, enemy.maxHealth), enemy.health, enemy.maxHealth)
}
