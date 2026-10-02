package main

import (
	"fmt"
)

func main() {
	describe("welcome")

	player := startPlayer()

	for player.isAlive() {
		fmt.Println("\n══════════════════════════════")
		fmt.Println("📍 Куда направишься?")
		fmt.Println("══════════════════════════════")
		fmt.Println("1 - 🌲 Тёмный лес")
		fmt.Println("2 - 🌿 Гнилое болото")
		fmt.Println("3 - 🏰 Старый замок")
		fmt.Println("4 - ⚰️  Старое кладбище")
		fmt.Println("5 - ⛏️  Пещера кобольдов")
		fmt.Println("6 - 🧝 Деревня эльфов")
		fmt.Println("7 - 🏪 Магазин")
		fmt.Println("8 - 📊 Статус героя")
		fmt.Println("9 - 💾 Сохранить игру")
		fmt.Println("10 - ⚔️  Экипировка")
		fmt.Println("0 - 🚪 Выйти из игры")
		fmt.Print("\nВыбор: ")

		choice := readChoice()

		switch choice {
		case 1:
			darkForest(player)
		case 2:
			swamp(player)
		case 3:
			oldCastle(player)
		case 4:
			graveyard(player)
		case 5:
			koboldCave(player)
		case 6:
			elfVillage(player)
		case 7:
			shop(player)
		case 8:
			player.printStatus()
		case 9:
			printSaveResult(saveGame(player))
		case 10:
			manageEquipment(player)
		case 0:
			if confirmSaveOnExit() {
				printSaveResult(saveGame(player))
			}
			fmt.Printf("\n👋 До свидания, %s!\n", player.name)
			return
		default:
			fmt.Println("❓ Такого варианта нет")
		}
	}

	describe("player_death")
	fmt.Printf("\nГерой %s пал в бою.\n", player.name)
	fmt.Printf("Собрано золота: %d\n", player.gold)
}

func startPlayer() *Player {
	if saveExists() {
		fmt.Println("\n💾 Найдено сохранение.")
		fmt.Println("1 - Продолжить игру")
		fmt.Println("2 - Начать новую игру")
		fmt.Print("\nВыбор: ")

		if readChoice() == 1 {
			player, err := loadGame()
			if err != nil {
				fmt.Println("⚠️  Не удалось загрузить сохранение, начинаем новую игру.")
			} else {
				fmt.Printf("\n⚔️  С возвращением, %s!\n", player.name)
				return player
			}
		}
	}

	fmt.Println("\nПредставься, авантюрист:")
	fmt.Print("> ")
	name := readLine()
	if name == "" {
		name = "Безымянный герой"
	}

	player := newPlayer(name)
	fmt.Printf("\n⚔️  Добро пожаловать, %s!\n", player.name)
	describe("intro")
	return player
}

func confirmSaveOnExit() bool {
	fmt.Println("\n1 - 💾 Сохранить и выйти")
	fmt.Println("2 - 🚪 Выйти без сохранения")
	fmt.Print("\nВыбор: ")
	return readChoice() == 1
}

func printSaveResult(err error) {
	if err != nil {
		fmt.Println("⚠️  Не удалось сохранить игру:", err)
		return
	}
	fmt.Println("💾 Игра сохранена.")
}
