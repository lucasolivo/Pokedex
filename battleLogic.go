package main 

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func battle(mon1 Pokemon, move1 string, mon2 Pokemon, move2 string) (Pokemon, Pokemon) {
	// Decide turn order
	leftFirst := true
	if mon1.CurStats["speed"] > mon2.CurStats["speed"] {
		leftFirst = true
	} else if mon1.CurStats["speed"] < mon2.CurStats["speed"] {
		leftFirst = false
	} else {
		CoinToss := rand.Intn(2)
		if CoinToss == 0 {
			leftFirst = true
		} else {
			leftFirst = false
		}
	}
	// use helper function to calculate damage
	if leftFirst{
		mon2 = damage(mon1, move1, mon2, 1)
		if mon2.CurHp == 0 {
			return mon1, mon2
		} else {
			mon1 = damage(mon2, move2, mon1, 2)
			return mon1, mon2
		}
	} else {
		mon1 = damage(mon2, move2, mon1, 1) 
		if mon1.CurHp == 0 {
			return mon1, mon2
		} else {
			mon2 = damage(mon1, move1, mon2, 2)
			return mon1, mon2
		}
	}
	
	
}

func switchMons(cfg *config, switchOut Pokemon, switchIn Pokemon) error{
	var first int
	var second int 
	for i, name := range cfg.PokeKeys {
		if name == switchOut.Name {
			first = i
		}
		if name == switchIn.Name {
			second = i
		}
	}
	if second == 0 {
		return fmt.Errorf("%v is not in your party.", switchIn.Name)
	}
	cfg.PokeKeys[first] = switchIn.Name
	cfg.PokeKeys[second] = switchOut.Name
	fmt.Printf("Go, %v!\n", switchIn.Name)
	return nil
}

func damage(attacker Pokemon, move string, defender Pokemon, order int) Pokemon {
	fmt.Printf("%v used %v!\n", attacker.Name, move)
	power := attacker.Movedata[move].Power
	accuracy := attacker.Movedata[move].Accuracy
	damageType := attacker.Movedata[move].Damagetype
	pokeType := attacker.Movedata[move].Poketype
	STABDam := 1.0
	// Check if attacking move type matches pokemon type
	for _, attackType := range attacker.Types{
		if attackType == pokeType {
			STABDam = 1.5
		}
	}
	STABDam = STABDam * DamageEffectAttack(attacker, defender, move, order)
	if power == "N/A" {
		fmt.Println("To be implemented")
		return defender
	} else {
		intPower, err := strconv.Atoi(power)
		if err != nil {
			fmt.Println("Problem with power conversion")
			return defender
		}
		// if accuracy is N/A then it'll always hit, accounting for -6 I made this like crazy high
		var floatAccuracy float64
		if accuracy == "N/A" {
			floatAccuracy = 10000.0
		} else {
			acc, err := strconv.Atoi(accuracy)
			if err != nil {
				fmt.Println("Problem with accuracy conversion")
				return defender
			}
			floatAccuracy = float64(acc)/100
		}
		hitChance := rand.Float64()
		if hitChance > floatAccuracy {
			fmt.Printf("%v avoided the attack!\n", defender.Name)
			return defender
		}
		multiplier := 1.0
		for _, Montype := range defender.Types {
			multiplier = multiplier * TypeChart[pokeType][Montype]
		}
		if multiplier == 0.0 {
			fmt.Printf("It doesn't affect %v...\n", defender.Name)
		} else if multiplier < 1.0 {
			fmt.Println("It's not very effective...")
		} else if multiplier > 1.0 {
			fmt.Println("It's super effective!")
		}
		// determine if attack is physical or special
		var damStat int
		var defStat int
		if damageType == "physical" {
			damStat = attacker.CurStats["attack"]
			defStat = defender.CurStats["defense"]
		} else {
			damStat = attacker.CurStats["special-attack"]
			defStat = defender.CurStats["special-defense"]
		}
		damageRoll := float64(rand.Intn(16)+85) / 100.0
		critChance := 1.0/24.0
		critDam := 1.0
		if critChance > rand.Float64() {
			fmt.Println("A critical hit!")
			critDam = 1.5
		}
		base := float64((2*attacker.Level)/5 + 2)
		scaling := float64(intPower) * float64(damStat) / float64(defStat)
		baseDamage := ((base * scaling) / 50.0 + 2) * critDam * STABDam * multiplier * damageRoll
		finalDamage := int(math.Ceil(baseDamage))

		
		defender.CurHp -= finalDamage

		fmt.Printf("%v took %v damage!, it now has %v health remaining.\n", defender.Name, finalDamage, max(0, defender.CurHp))

		if defender.CurHp < 0 {
			defender.CurHp = 0
			fmt.Printf("%v Fainted!\n", defender.Name)
		}

		return defender
	}
}

func expGain(winner Pokemon, loser Pokemon, isTrainer bool, cfg *config){
	a := 1.0
	if isTrainer{
		a = 1.5
	}
	b := float64(loser.BaseExperience)
	l := float64(loser.Level)
	lp := float64(winner.Level)
	gain := ((b * l) / 5.0 * math.Pow((2*l+10)/(l+lp+10), 2.5) + 1) * a
	gainInt := int(math.Floor(gain))
	winner.TotalExp = winner.TotalExp + gainInt
	fmt.Printf("You gained %v experience!\n", gainInt)
	for {
		// Check the EXP required for the next level
		if winner.Level >= 100 {
			// Already max level
			break
		}
		nextLevel := expChart[winner.ExpGroup][winner.Level+1]
		if winner.TotalExp < nextLevel {
			// Not enough EXP to level up further
			cfg.Pokedex[winner.Name] = winner 
			cfg.Party[winner.Name] = winner
			break
		}

		// Level up the Pokémon
		name, updatedWinner := level(cfg, winner, winner.Name)
		winner = updatedWinner

		// Update Pokedex & Party
		cfg.Pokedex[name] = winner
		cfg.Party[name] = winner
	}
}