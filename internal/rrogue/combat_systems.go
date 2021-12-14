package rrogue

import (
	"fmt"

	"github.com/bytearena/ecs"
)

func AttackSystem(g *Game, attackerPosition, defenderPosition *Position) {
	var attacker *ecs.QueryResult = nil
	var defender *ecs.QueryResult = nil

	for _, playerCombatant := range g.World.Query(g.WorldTags["players"]) {
		pos := playerCombatant.Components[position].(*Position)

		if pos.isEqual(attackerPosition) {
			attacker = playerCombatant
		} else if pos.isEqual(defenderPosition) {
			defender = playerCombatant
		}
	}

	for _, cbt := range g.World.Query(g.WorldTags["monsters"]) {
		pos := cbt.Components[position].(*Position)
		if pos.isEqual(attackerPosition) {
			attacker = cbt
		} else if pos.isEqual(defenderPosition) {
			defender = cbt
		}
	}

	if attacker == nil || defender == nil {
		return
	}

	defenderArmor := defender.Components[armor].(*Armor)
	defenderHealth := defender.Components[health].(*Health)
	defenderName := defender.Components[name].(*Name).Label

	attackerWeapon := attacker.Components[meleeWeapon].(*MeleeWeapon)
	attackerName := attacker.Components[name].(*Name).Label

	toHitRoll := GetDiceRoll(10)

	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
		damageRoll := GetRandomBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)
		damageDone := damageRoll - defenderArmor.Defense

		if damageDone < 0 {
			damageDone = 0
		}

		defenderHealth.CurrentHealth -= damageDone
		fmt.Printf(
			"%s swings %s at %s and hits for %d health\n",
			attackerName,
			attackerWeapon.Name,
			defenderName,
			damageDone,
		)

		if defenderHealth.CurrentHealth <= 0 {
			fmt.Printf("%s has died!\n", defenderName)
			if defenderName == "Player" {
				fmt.Printf("Game Over!\n")
				g.Turn = GameOVer
			}
			g.World.DisposeEntities(defender.Entity)
		}
	} else {
		fmt.Printf(
			"%s swings %s at %s and misses\n",
			attackerName,
			attackerWeapon.Name,
			defenderName,
		)
	}
}
