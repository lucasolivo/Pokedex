package main

import (
	"fmt"
	"math/rand"
)

func checkStatEffectMove(move string, attacker Pokemon, defender Pokemon) (Pokemon, Pokemon, bool){
	hit := false
	switch{
	//lower accuracy
	case move == "flash" || move == "kinesis" || move == "sand-attack" || move == "smokescreen":
		hit = true
		if defender.StatEffects["accuracy"] == -6{
			fmt.Printf("%v's accuracy can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's accuracy fell!\n", defender.Name)
		defender.StatEffects["accuracy"] -= 1
		break

	//lower attack
	case move == "baby-doll-eyes" || move == "growl" || move == "play-nice":
		hit = true
		if defender.StatEffects["attack"] == -6{
			fmt.Printf("%v's attack can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's attack fell!\n", defender.Name)
		defender.StatEffects["attack"] -= 1
		break

	//harshly lower attack
	case move == "charm" || move == "feather-dance":
		hit = true
		if defender.StatEffects["attack"] == -6{
			fmt.Printf("%v's attack can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's attack harshly fell!\n", defender.Name)
		defender.StatEffects["attack"] = max(-6, defender.StatEffects["attack"] - 2)
		break

	//lower attack and special attack
	case move == "noble-roar" || move == "parting-shot" || move == "tearful-look":
		hit = true
		if defender.StatEffects["attack"] == -6{
			fmt.Printf("%v's attack can't go lower!\n", defender.Name)
		} else {
			fmt.Printf("%v's attack fell!\n", defender.Name)
			defender.StatEffects["attack"] -= 1
		}
		if defender.StatEffects["special-attack"] == -6{
			fmt.Printf("%v's special attack can't go lower!\n", defender.Name)
			break
		} else {
			fmt.Printf("%v's special attack fell!\n", defender.Name)
			defender.StatEffects["special-attack"] -= 1
			break
		}

	//Lower Defense and Special Defense
	case move == "octolock":
		hit = true
		if defender.StatEffects["defense"] == -6{
			fmt.Printf("%v's defense can't go lower!\n", defender.Name)
		} else {
			fmt.Printf("%v's defense fell!\n", defender.Name)
			defender.StatEffects["defense"] -= 1
		}
		if defender.StatEffects["special-defense"] == -6{
			fmt.Printf("%v's special defense can't go lower!\n", defender.Name)
			break
		} else {
			fmt.Printf("%v's special defense fell!\n", defender.Name)
			defender.StatEffects["special-defense"] -= 1
			break
		}

	//lower defense
	case move == "leer" || move == "tail-whip":
		hit = true
		if defender.StatEffects["defense"] == -6{
			fmt.Printf("%v's defense can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's defense fell!\n", defender.Name)
		defender.StatEffects["defense"] -= 1
		break

	//harshly lower defense
	case move == "screech":
		hit = true
		if defender.StatEffects["defense"] == -6{
			fmt.Printf("%v's defense can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's defense harshly fell!\n", defender.Name)
		defender.StatEffects["defense"] = max(-6, defender.StatEffects["attack"] - 2)
		break

	// harshly lower defense, sharply raise attack
	case move == "spicy-extract":
		hit = true
		if defender.StatEffects["defense"] == -6{
			fmt.Printf("%v's defense can't go lower!\n", defender.Name)
		} else {
			fmt.Printf("%v's defense harshly fell!\n", defender.Name)
			defender.StatEffects["defense"] = max(-6, defender.StatEffects["defense"] - 2)
		}
		if defender.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", defender.Name)
			break
		}
		fmt.Printf("%v's attack sharply rose!\n", defender.Name)
		defender.StatEffects["attack"] = min(6, defender.StatEffects["attack"] + 2)
		break

	//Lower Evasion
	case move == "sweet-scent" || move == "defog":
		hit = true
		if defender.StatEffects["evasion"] == -6{
			fmt.Printf("%v's evasion can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's evasion fell!\n", defender.Name)
		defender.StatEffects["evasion"] -= 1
		break

	//Lower special attack harshly
	case move == "captivate" || move == "eerie-impulse":
		hit = true
		if defender.StatEffects["special-attack"] == -6{
			fmt.Printf("%v's special attack can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's special-attack harshly fell!\n", defender.Name)
		defender.StatEffects["special-attack"] = max(-6, defender.StatEffects["special-attack"] - 2)
		break

	//Lower special attack
	case move == "confide":
		hit = true
		if defender.StatEffects["special-attack"] == -6{
			fmt.Printf("%v's special attack can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's special attack fell!\n", defender.Name)
		defender.StatEffects["special-attack"] -= 1
		break

	//Lower attack and special attack harshly
	case move == "memento":
		hit = true
		if defender.StatEffects["special-attack"] == -6{
			fmt.Printf("%v's special attack can't go lower!\n", defender.Name)
		} else {
			fmt.Printf("%v's special-attack harshly fell!\n", defender.Name)
			defender.StatEffects["special-attack"] = max(-6, defender.StatEffects["special-attack"] - 2)
		}
		if defender.StatEffects["attack"] == -6{
			fmt.Printf("%v's attack can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's attack harshly fell!\n", defender.Name)
		defender.StatEffects["attack"] = max(-6, defender.StatEffects["attack"] - 2)
		break

	//Harshly lower special defense
	case move == "fake-tears" || move == "metal-sound":
		hit = true
		if defender.StatEffects["special-defense"] == -6{
			fmt.Printf("%v's special defense can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's special-defense harshly fell!\n", defender.Name)
		defender.StatEffects["special-defense"] = max(-6, defender.StatEffects["special-defense"] - 2)
		break

	//harshly lower speed
	case move == "cotton-spore" || move == "scary-face" || move == "string-shot":
		hit = true
		if defender.StatEffects["speed"] == -6{
			fmt.Printf("%v's speed can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's speed harshly fell!\n", defender.Name)
		defender.StatEffects["speed"] = max(-6, defender.StatEffects["speed"] - 2)
		break

	//Lower speed
	case move == "tar-shot" || move == "toxic-thread":
		hit = true
		if defender.StatEffects["speed"] == -6{
			fmt.Printf("%v's speed can't go lower!\n", defender.Name)
			break
		}
		fmt.Printf("%v's speed fell!\n", defender.Name)
		defender.StatEffects["speed"] -= 1
		break

	//affects user's stats

	// Shell smash
	case move == "shell-smash":
		hit = true
		if attacker.StatEffects["special-defense"] == -6{
			fmt.Printf("%v's special defense can't go lower!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special-defense harshly fell!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = max(-6, attacker.StatEffects["special-defense"] - 2)
		}
		if attacker.StatEffects["defense"] == -6{
			fmt.Printf("%v's defense can't go lower!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense harshly fell!\n", attacker.Name)
			attacker.StatEffects["defense"] = max(-6, attacker.StatEffects["defense"] - 2)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 2)
		}
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 2)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed sharply rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 2)
		}
		break

	//Curse
	case move == "curse":
		hit = true
		if attacker.StatEffects["speed"] == -6{
			fmt.Printf("%v's speed can't go lower!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed fell!\n", attacker.Name)
			attacker.StatEffects["speed"] -= 1
		}
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] += 1
		}
		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
			break
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] += 1
			break
		}

	case move == "acupressure":
		hit = true
		toChange := chooseStatAccupressure(attacker)
		if len(toChange) == 0 {
			fmt.Println("The move failed!")
			break
		}
		dex := rand.Intn(len(toChange)) - 1
		changing := toChange[dex]
		fmt.Printf("%v's %v sharply rose!\n", attacker.Name, changing)
		attacker.StatEffects[changing] = min(6, attacker.StatEffects[changing] + 2)
		break
	
	case move == "decorate":
		hit = true
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 2)
		}
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 2)
		}
		break

	case move == "rototiller":
		hit = true
		connected := false
		eligible := false
		for _, typ := range attacker.Types{
			if typ == "flying" {
				eligible = false
				break
			}
			if typ == "grass" {
				eligible = true
			}
		}

		if eligible {
			connected = true
			if attacker.StatEffects["attack"] == 6{
				fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
			} else {
				fmt.Printf("%v's attack rose!\n", attacker.Name)
				attacker.StatEffects["attack"] += 1
			}
			if attacker.StatEffects["special-attack"] == 6{
				fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
			} else {
				fmt.Printf("%v's special attack rose!\n", attacker.Name)
				attacker.StatEffects["special-attack"] += 1
			}
		}

		for _, typ := range defender.Types{
			if typ == "flying" {
				eligible = false
				break
			}
			if typ == "grass" {
				eligible = true
			}
		}

		if eligible{
			connected = true
			if eligible {
			connected = true
			if defender.StatEffects["attack"] == 6{
				fmt.Printf("%v's attack can't go higher!\n", defender.Name)
			} else {
				fmt.Printf("%v's attack rose!\n", defender.Name)
				defender.StatEffects["attack"] += 1
			}
			if defender.StatEffects["special-attack"] == 6{
				fmt.Printf("%v's special attack can't go higher!\n", defender.Name)
			} else {
				fmt.Printf("%v's special attack rose!\n", defender.Name)
				defender.StatEffects["special-attack"] += 1
			}
		}
		}

		if !connected {
			fmt.Println("The move failed!")
		}

	case move == "flower-shield":
		hit = true
		for _, typ := range attacker.Types{
			if typ == "grass"{
				if attacker.StatEffects["attack"] == 6{
					fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
					break
				} else {
					fmt.Printf("%v's attack rose!\n", attacker.Name)
					attacker.StatEffects["attack"] += 1
					break
				}
			}
		}
		for _, typ := range defender.Types{
			if typ == "grass"{
				if defender.StatEffects["defense"] == 6{
				fmt.Printf("%v's defense can't go higher!\n", defender.Name)
				break
			} else {
				fmt.Printf("%v's defense rose!\n", defender.Name)
				defender.StatEffects["defense"] += 1
				break
			}
			}
		}

	case move == "coil":
		hit = true
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}

		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}

		if attacker.StatEffects["accuracy"] == 6{
			fmt.Printf("%v's accuracy can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's accuracy rose!\n", attacker.Name)
			attacker.StatEffects["accuracy"] = min(6, attacker.StatEffects["accuracy"] + 1)
		}
		break
	
	case move == "hone-claws":
		hit = true
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["accuracy"] == 6{
			fmt.Printf("%v's accuracy can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's accuracy rose!\n", attacker.Name)
			attacker.StatEffects["accuracy"] = min(6, attacker.StatEffects["accuracy"] + 1)
		}
		break

	case move == "belly-drum":
		hit = true
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			if attacker.CurHp < attacker.CurStats["hp"] / 2{
				fmt.Printf("The move failed.")
				break
			}
			attacker.CurHp -= attacker.CurStats["hp"] / 2
			fmt.Printf("%v's attack was maxed out!\n", attacker.Name)
			attacker.StatEffects["attack"] = 6
			break
		}

	case move == "bulk-up":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}
		break

	case move == "clangorous-soul":
		hit = true
		if attacker.CurHp < attacker.CurStats["hp"] / 3 || 
		(attacker.StatEffects["attack"] == 6 && attacker.StatEffects["special-attack"] == 6 && 
		attacker.StatEffects["defense"] == 6 &&
		attacker.StatEffects["special-defense"] == 6 && attacker.StatEffects["speed"] == 6){
			fmt.Printf("The move failed.")
			break
		}
		attacker.CurHp -= attacker.CurStats["hp"] / 3
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 1)
		}
		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 1)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 1)
		}
		break

	case move == "dragon-dance":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 1)
		}
		break

	case move == "fillet-away":
		hit = true

		if attacker.CurHp < attacker.CurStats["hp"] / 2 || 
		(attacker.StatEffects["attack"] == 6 && attacker.StatEffects["special-attack"] == 6 &&
		attacker.StatEffects["speed"] == 6){
			fmt.Printf("The move failed.")
			break
		}

		attacker.CurHp -= attacker.CurStats["hp"] / 2

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 2)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 2)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed sharply rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 2)
		}

	case move == "growth":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 1)
		}
		break

	case move == "howl":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		break
		
	case move == "meditate":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		break

	case move == "no-retreat":
		hit = true

		// add no switch logic

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 1)
		}
		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 1)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 1)
		}
		break

	case move == "sharpen":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		break

	case move == "shift-gear":
		hit = true
		
		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed sharply rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 2)
		}

		break
		
	case move == "swords-dance":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 2)
		}

		break

	case move == "tidy-up":
		hit = true

		//logic to remove hazards

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 1)
		}
		break

	case move == "victory-dance":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 1)
		}
		break

	case move == "work-up":
		hit = true

		if attacker.StatEffects["attack"] == 6{
			fmt.Printf("%v's attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's attack rose!\n", attacker.Name)
			attacker.StatEffects["attack"] = min(6, attacker.StatEffects["attack"] + 1)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 1)
		}
		break

		
	case move == "acid-armor" || move == "barrier" || move == "iron-defense" || move == "shelter":
		hit = true

		if defender.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
			break
		} else {
			fmt.Printf("%v's defense sharply rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 2)
			break
		}
		
	case move == "cosmic-power" || move == "defend-order":
		hit = true

		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}
		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 1)
		}

		break

	case move == "cotton-guard":
		hit = true

		if defender.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
			break
		} else {
			fmt.Printf("%v's defense drastically rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 3)
			break
		}

	case move == "defense-curl" || move == "harden" || move == "withdraw":
		hit = true

		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}

		break

	case move == "stockpile": // add stockpile/swallow/spit-up logic?
		hit = true

		if attacker.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's defense rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 1)
		}
		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 1)
		}

		break

	case move == "stuff-cheeks": //implement berry eating logic
		hit = true

		if defender.StatEffects["defense"] == 6{
			fmt.Printf("%v's defense can't go higher!\n", attacker.Name)
			break
		} else {
			fmt.Printf("%v's defense sharply rose!\n", attacker.Name)
			attacker.StatEffects["defense"] = min(6, attacker.StatEffects["defense"] + 2)
			break
		}

	case move == "double-team" || move == "minimize":
		hit = true

		if attacker.StatEffects["evasion"] == 6{
			fmt.Printf("%v's evasion can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's evasion sharply rose!\n", attacker.Name)
			attacker.StatEffects["evasion"] = min(6, attacker.StatEffects["evasion"] + 2)
		}

		break

	case move == "calm-mind":
		hit = true

		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 1)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 1)
		}
		break

	case move == "geomancy":
		hit = true

		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense sharply rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 2)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 2)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed sharply rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 2)
		}
		break

	case move == "quiver-dance":
		hit = true

		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 1)
		}
		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 1)
		}
		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 1)
		}
		break

	case move == "nasty-plot":
		hit = true

		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack sharply rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 2)
		}

		break

	case move == "tail-glow":
		hit = true

		if attacker.StatEffects["special-attack"] == 6{
			fmt.Printf("%v's special attack can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special attack drastically rose!\n", attacker.Name)
			attacker.StatEffects["special-attack"] = min(6, attacker.StatEffects["special-attack"] + 3)
		}

		break

	case move == "amnesia":
		hit = true

		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense sharply rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 2)
		}

		break

	case move == "charge": //add logic for next electric move to double
		hit = true

		if attacker.StatEffects["special-defense"] == 6{
			fmt.Printf("%v's special defense can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's special defense rose!\n", attacker.Name)
			attacker.StatEffects["special-defense"] = min(6, attacker.StatEffects["special-defense"] + 1)
		}

		break

	case move == "agility" || move == "rock-polish":
		hit = true

		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed sharply rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 2)
		}

		break

	case move == "autotomize": //add logic to reduce pokemon weight
		hit = true

		if attacker.StatEffects["speed"] == 6{
			fmt.Printf("%v's speed can't go higher!\n", attacker.Name)
		} else {
			fmt.Printf("%v's speed sharply rose!\n", attacker.Name)
			attacker.StatEffects["speed"] = min(6, attacker.StatEffects["speed"] + 2)
		}

		break
	
	default:
		break
	}
	return attacker, defender, hit
}

//func checkStatusMove()

func chooseStatAccupressure(mon Pokemon) []string {
	canChange := []string{}
	for name, val := range mon.StatEffects{
		if val < 6 {
			canChange = append(canChange, name)
		}
	}
	return canChange
}

func resetStatEffects() map[string]int {
    return map[string]int{
        "attack":    0,
        "defense":   0,
        "sp_attack": 0,
        "sp_defense": 0,
        "speed":     0,
        "accuracy":  0,
        "evasion":   0,
    }
}