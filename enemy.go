package main

type SpellType string

const (
	SpellFire      SpellType = "fire"
	SpellIce       SpellType = "ice"
	SpellLightning SpellType = "lightning"
	SpellLight     SpellType = "light"
)

type MonsterClass string

const (
	ClassBeast    MonsterClass = "beast"
	ClassHumanoid MonsterClass = "humanoid"
	ClassUndead   MonsterClass = "undead"
	ClassBoss     MonsterClass = "boss"
)

type Resistance struct {
	spellType  SpellType
	multiplier float64
}

type Enemy struct {
	name         string
	health       int
	maxHealth    int
	damage       int
	drop         int
	expReward    int
	materialDrop string
	uniqueDrop   string
	class        MonsterClass
	dropKey      bool
	dropAmulet   bool
	dropSpell    string
	resistances  []Resistance
}

func undeadResistances() []Resistance {
	return []Resistance{
		{SpellFire, highResistMultiplier},
		{SpellIce, highResistMultiplier},
		{SpellLight, lightRayVulnMultiplier},
		{SpellLightning, vulnerableMultiplier},
	}
}

func newWolf() *Enemy {
	return &Enemy{
		name:         "Волк",
		health:       wolfHealth,
		maxHealth:    wolfHealth,
		damage:       wolfDamage,
		drop:         wolfDrop,
		expReward:    wolfExp,
		materialDrop: "wolf_fang",
		class:        ClassBeast,
		resistances: []Resistance{
			{SpellFire, vulnerableMultiplier},
			{SpellIce, resistMultiplier},
		},
	}
}

func newGoblin() *Enemy {
	return &Enemy{
		name:         "Гоблин",
		health:       goblinHealth,
		maxHealth:    goblinHealth,
		damage:       goblinDamage,
		drop:         goblinDrop,
		expReward:    goblinExp,
		materialDrop: "goblin_ear",
		class:        ClassHumanoid,
		dropKey:      true,
		resistances: []Resistance{
			{SpellFire, resistMultiplier},
			{SpellIce, vulnerableMultiplier},
		},
	}
}

func newKnight() *Enemy {
	return &Enemy{
		name:       "Тёмный рыцарь",
		health:     knightHealth,
		maxHealth:  knightHealth,
		damage:     knightDamage,
		drop:       knightDrop,
		expReward:  knightExp,
		class:      ClassHumanoid,
		dropAmulet: true,
		resistances: []Resistance{
			{SpellLightning, vulnerableMultiplier},
		},
	}
}

func newArmoryGuard() *Enemy {
	return &Enemy{
		name:         "Страж оружейной",
		health:       armoryGuardHealth,
		maxHealth:    armoryGuardHealth,
		damage:       armoryGuardDamage,
		drop:         armoryGuardDrop,
		expReward:    armoryGuardExp,
		materialDrop: "guard_insignia",
		class:        ClassHumanoid,
		resistances: []Resistance{
			{SpellLightning, vulnerableMultiplier},
		},
	}
}

func newSkeleton() *Enemy {
	return &Enemy{
		name:         "Скелет",
		health:       skeletonHealth,
		maxHealth:    skeletonHealth,
		damage:       skeletonDamage,
		drop:         skeletonDrop,
		expReward:    skeletonExp,
		materialDrop: "undead_ash",
		class:        ClassUndead,
		resistances:  undeadResistances(),
	}
}

func newZombie() *Enemy {
	return &Enemy{
		name:         "Зомби",
		health:       zombieHealth,
		maxHealth:    zombieHealth,
		damage:       zombieDamage,
		drop:         zombieDrop,
		expReward:    zombieExp,
		materialDrop: "undead_ash",
		class:        ClassUndead,
		resistances:  undeadResistances(),
	}
}

func newNecromancer() *Enemy {
	return &Enemy{
		name:       "Некромант",
		health:     necromancerHealth,
		maxHealth:  necromancerHealth,
		damage:     necromancerDamage,
		drop:       necromancerDrop,
		expReward:  necromancerExp,
		uniqueDrop: "necromancer_staff",
		class:      ClassHumanoid,
		resistances: []Resistance{
			{SpellFire, resistMultiplier},
			{SpellLightning, vulnerableMultiplier},
		},
	}
}

func newSoul() *Enemy {
	return &Enemy{
		name:         "Неупокоенная душа",
		health:       soulHealth,
		maxHealth:    soulHealth,
		damage:       soulDamage,
		drop:         soulDrop,
		expReward:    soulExp,
		materialDrop: "undead_ash",
		class:        ClassUndead,
		resistances:  undeadResistances(),
	}
}

func newKobold() *Enemy {
	return &Enemy{
		name:         "Кобольд",
		health:       koboldHealth,
		maxHealth:    koboldHealth,
		damage:       koboldDamage,
		drop:         koboldDrop,
		expReward:    koboldExp,
		materialDrop: "kobold_eye",
		class:        ClassHumanoid,
		resistances: []Resistance{
			{SpellLight, resistMultiplier},
			{SpellFire, vulnerableMultiplier},
		},
	}
}

func newKoboldKing() *Enemy {
	return &Enemy{
		name:       "Король кобольдов",
		health:     koboldKingHealth,
		maxHealth:  koboldKingHealth,
		damage:     koboldKingDamage,
		drop:       koboldKingDrop,
		expReward:  koboldKingExp,
		uniqueDrop: "kobold_king_crown",
		class:      ClassBoss,
		dropSpell:  "light_ray",
		resistances: []Resistance{
			{SpellFire, resistMultiplier},
			{SpellIce, resistMultiplier},
			{SpellLight, vulnerableMultiplier},
		},
	}
}

func (e *Enemy) scaleToLevel(level int) {
	if level <= 1 {
		return
	}
	multiplier := 1 + float64(level-1)*enemyLevelScaling

	e.maxHealth = int(float64(e.maxHealth) * multiplier)
	e.health = e.maxHealth
	e.damage = int(float64(e.damage) * multiplier)
	e.drop = int(float64(e.drop) * multiplier)
	e.expReward = int(float64(e.expReward) * multiplier)
}

func (e *Enemy) takeDamage(damage int) {
	e.health -= damage
	if e.health < 0 {
		e.health = 0
	}
}

func (e *Enemy) takeSpellDamage(damage int, spellType SpellType) int {
	multiplier := 1.0
	for _, r := range e.resistances {
		if r.spellType == spellType {
			multiplier = r.multiplier
			break
		}
	}
	finalDamage := int(float64(damage) * multiplier)
	e.takeDamage(finalDamage)
	return finalDamage
}

func (e *Enemy) isAlive() bool {
	return e.health > 0
}

func (e *Enemy) isUndead() bool {
	return e.class == ClassUndead
}
