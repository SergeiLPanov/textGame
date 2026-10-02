package main

import "fmt"

func elfVillage(player *Player) {
	renderLocation("elf_village", "🧝 ДЕРЕВНЯ ЭЛЬФОВ")
	describe("elf_village_enter")

	fmt.Println("1 - 🏠 Хижина старосты")
	fmt.Println("2 - ⛺ Палатка лекаря")
	fmt.Println("3 - 🏚️  Лачуга главной охотницы")
	fmt.Println("4 - Вернуться назад")
	fmt.Print("\nВыбор: ")

	switch readChoice() {
	case 1:
		elderHut(player)
	case 2:
		healerTent(player)
	case 3:
		huntressHut(player)
	case 4:
		fmt.Println("Ты покидаешь деревню.")
	default:
		fmt.Println("Неизвестная команда")
	}
}

func elderHut(player *Player) {
	describe("elder_hut_enter")
	if player.questElderDone {
		describe("elder_hut_done")
		return
	}
	offerVillageQuest(player, "goblin_ear", elderQuestGoblinEarCount, func() {
		player.questElderDone = true
		describe("elder_hut_reward")
		player.addGold(elderQuestRewardGold)
		player.addExp(elderQuestRewardExp)
	})
}

func healerTent(player *Player) {
	describe("healer_tent_enter")
	if player.questHealerDone {
		describe("healer_tent_done")
		return
	}
	offerVillageQuest(player, "undead_ash", healerQuestUndeadAshCount, func() {
		player.questHealerDone = true
		describe("healer_tent_reward")
		for i := 0; i < healerQuestRewardHealthPotions; i++ {
			player.inventory.addItem(allItems["health_potion"])
		}
		for i := 0; i < healerQuestRewardManaPotions; i++ {
			player.inventory.addItem(allItems["mana_potion"])
		}
		player.addExp(healerQuestRewardExp)
	})
}

func huntressHut(player *Player) {
	describe("huntress_hut_enter")
	if player.questHuntressDone {
		describe("huntress_hut_done")
		return
	}
	offerVillageQuest(player, "wolf_fang", huntressQuestWolfFangCount, func() {
		player.questHuntressDone = true
		describe("huntress_hut_reward")
		player.addGold(huntressQuestRewardGold)
		player.addExp(huntressQuestRewardExp)
		for i := 0; i < huntressQuestRewardHealthPotions; i++ {
			player.inventory.addItem(allItems["health_potion"])
		}
	})
}

func offerVillageQuest(player *Player, itemID string, itemCount int, onReward func()) {
	have := player.inventory.countItem(itemID)
	fmt.Printf("\nНужно: %s x%d (у тебя: %d)\n", allItems[itemID].name, itemCount, have)
	fmt.Println("1 - Отдать трофеи")
	fmt.Println("2 - Уйти")
	fmt.Print("\nВыбор: ")

	if readChoice() != 1 {
		return
	}
	if have < itemCount {
		fmt.Println("❌ У тебя пока недостаточно трофеев.")
		return
	}
	for i := 0; i < itemCount; i++ {
		player.inventory.removeItem(itemID)
	}
	onReward()
}
