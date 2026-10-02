package main

import "fmt"

func manageEquipment(p *Player) {
	for {
		fmt.Println("\n╔═══════════════════════════════╗")
		fmt.Println("║       ⚔️  ЭКИПИРОВКА          ║")
		fmt.Println("╚═══════════════════════════════╝")
		printEquipped(p)

		gear := p.inventory.equippableItems()
		if len(gear) == 0 {
			fmt.Println("\nВ инвентаре нет предметов для экипировки.")
			fmt.Println("0 - Назад")
			fmt.Print("\nВыбор: ")
			readChoice()
			return
		}

		fmt.Println("\nДоступно для экипировки:")
		for i, item := range gear {
			fmt.Printf("%d - %s (%s)\n", i+1, item.name, item.description)
		}
		fmt.Println("0 - Назад")
		fmt.Print("\nВыбор: ")

		choice := readChoice()
		if choice == 0 || choice > len(gear) {
			return
		}
		p.equipItem(gear[choice-1])
	}
}

func printEquipped(p *Player) {
	weapon := "пусто"
	if p.weapon != nil {
		weapon = fmt.Sprintf("%s (+%d урона)", p.weapon.name, p.weapon.value)
	}
	armor := "пусто"
	if p.armorItem != nil {
		armor = fmt.Sprintf("%s (-%d урона от атак)", p.armorItem.name, p.armorItem.value)
	}
	fmt.Printf("⚔️  Оружие: %s\n", weapon)
	fmt.Printf("🛡️  Броня:  %s\n", armor)
}
