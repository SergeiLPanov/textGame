package assets

const Hero = `  ⠀⢀⡀⠀
  ⣿⣿⣿
 ⢰⣿⣿⣿⡆
  ⠀⣿⣿
  ⢸⠀⠸`

const Wolf = `  ⣀⠀⠀⣀
  ⣿⣷⣾⣿
  ⠸⣿⣿⠇
  ⠀⣿⣿⠀
  ⠀⠛⠛⠀`

const Goblin = `  ⢀⣀⡀
  ⣾⣿⣷
  ⣿⣿⣿
  ⡇⠀⢸
  ⠓⠀⠚`

const Knight = `  ⣶⣶⣶
  ⣿⣿⣿
  ⣿⣿⣿
  ⣿⣿⣿
  ⠿⠀⠿`

const Skeleton = `  ⢀⣀⡀
  ⣿⠛⣿
  ⠿⣿⠿
  ⠀⣿⠀
  ⣀⠁⢀`

const Zombie = `  ⣶⣶⣶
  ⣿⣿⣿
  ⣿⣿⣿
  ⢸⣿⡇
  ⣿⠀⣿`

const Soul = `  ⢀⣀⡀
  ⣿⣿⣿
  ⣿⣿⣿
  ⠻⣿⠟
  ⠀⠛⠀`

const Necromancer = `  ⣀⣀⣀
  ⣿⣿⣿
  ⣿⣿⣿
  ⢸⣿⡇
  ⣾⠀⣷`

const Kobold = `  ⠀⣀⠀
  ⣾⣿⣷
  ⣿⣿⣿
  ⡇⠀⢸
  ⠓⠀⠚`

const KoboldKing = `  ⣶⣦⣶
  ⣿⣿⣿
  ⣿⣿⣿
  ⣿⠀⣿
  ⠿⠀⠿`

const Forest = `⢀⡀⠀⢀⡀⢀⡀
⣿⣿⣿⣿⣿⣿⣿
⠻⣿⣿⣿⣿⣿⠟
⠀⢸⣿⣿⣿⡇⠀
⠀⠘⠛⠛⠛⠃⠀`

const Swamp = `⠶⠶⠶⠶⠶⠶⠶
⣴⣦⠀⣴⣦⠀⠀
⣿⣿⣿⣿⣿⠀⠀
⠛⠛⠛⠛⠛⠀⠀
⠀⠀⠛⠛⠀⠀⠀`

const Castle = `⠀⣿⠀⣿⠀⣿
⣾⣿⣿⣿⣿⣷
⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿
⠛⠛⠛⠛⠛⠛`

const Graveyard = `⠀⣿⠀⣿⠀⠀
⠀⣿⠀⣿⠀⠀
⣶⣿⣶⣿⣶⠀
⠀⠀⠀⠀⠀⠀
⠛⠛⠛⠛⠛⠛`

const Cave = `⣤⣤⣤⣤⣤⣤
⣿⠀⠀⠀⠀⣿
⣿⠀⠀⠀⠀⣿
⣿⣤⠀⣤⣿⠀
⠀⠀⠀⠀⠀⠀`

const ElfVillage = `⠀⢀⡀⢀⡀⠀
⢠⣾⣿⣷⣦⡄
⠀⣿⣿⣿⣿⠀
⠀⠛⣿⣿⠛⠀
⠀⠀⠛⠛⠀⠀`

func HealthBar(current, max int) string {
	barLen := 10
	filled := current * barLen / max
	if filled < 0 {
		filled = 0
	}
	bar := "❤️  ["
	for i := 0; i < barLen; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += "]"
	return bar
}

func ManaBar(current, max int) string {
	barLen := 10
	filled := current * barLen / max
	if filled < 0 {
		filled = 0
	}
	bar := "💙 ["
	for i := 0; i < barLen; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += "]"
	return bar
}

func BattleScreen(heroSprite, enemySprite, enemyName string) string {
	heroLines := splitLines(heroSprite)
	enemyLines := splitLines(enemySprite)

	maxLines := len(heroLines)
	if len(enemyLines) > maxLines {
		maxLines = len(enemyLines)
	}

	screen := "\n╔══════════════════════════════════════╗\n"
	screen += "║          ⚔️  БОЙ  ⚔️                 ║\n"
	screen += "╠══════════════════════════════════════╣\n"

	for i := 0; i < maxLines; i++ {
		heroLine := "     "
		if i < len(heroLines) {
			heroLine = heroLines[i]
		}
		enemyLine := "     "
		if i < len(enemyLines) {
			enemyLine = enemyLines[i]
		}
		screen += "║  " + padRight(heroLine, 14) + "  vs  " + padRight(enemyLine, 14) + "║\n"
	}

	screen += "╠══════════════════════════════════════╣\n"
	screen += "║  🧙 ГЕРОЙ" + spaces(28) + "║\n"
	screen += "║  👹 " + padRight(enemyName, 33) + "║\n"
	screen += "╚══════════════════════════════════════╝"
	return screen
}

func LocationScreen(sprite, name string) string {
	lines := splitLines(sprite)
	screen := "\n╔══════════════════════════╗\n"
	screen += "║  " + padRight(name, 24) + "║\n"
	screen += "╠══════════════════════════╣\n"
	for _, line := range lines {
		screen += "║  " + padRight(line, 24) + "║\n"
	}
	screen += "╚══════════════════════════╝"
	return screen
}

func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, ch := range s {
		if ch == '\n' {
			if current != "" {
				lines = append(lines, current)
			}
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func padRight(s string, length int) string {
	runeCount := 0
	for range []rune(s) {
		runeCount++
	}
	for runeCount < length {
		s += " "
		runeCount++
	}
	return s
}

func spaces(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}
