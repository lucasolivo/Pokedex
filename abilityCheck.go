package main

import (
	"strconv"
	"fmt"
)

func DamageEffectAttack(attacker Pokemon, defender Pokemon, move string, order int) float64{
	multiplier := 1.0
	switch {
	case attacker.Ability == "neutralizing-gas":
		break
	case defender.Ability == "neutralizing-gas":
		break
	case attacker.Ability == "adaptability":
		for _, Ptype := range attacker.Types{
			if attacker.Movedata[move].Poketype == Ptype {
				multiplier = multiplier * 1.333
			}
		}
	case attacker.Ability == "aerilate":
		if attacker.Movedata[move].Poketype == "normal"{
			multiplier = multiplier * 1.2
		}
	case attacker.Ability == "analytic":
		if order == 2{
			multiplier = multiplier * 1.3
		}
	case attacker.Ability == "blaze":
		if attacker.Movedata[move].Poketype == "fire"{
			if attacker.CurHp < attacker.CurStats["hp"]/3 {
				multiplier = multiplier * 1.5
			}
		}
	case attacker.Ability == "dark-aura":
		if attacker.Movedata[move].Poketype == "dark"{
			multiplier = multiplier * 1.33
		}
	case defender.Ability == "dark-aura":
		if attacker.Movedata[move].Poketype == "dark"{
			multiplier = multiplier * 1.33
		}
	case attacker.Ability == "dragon's-maw":
		if attacker.Movedata[move].Poketype == "dragon"{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "fairy-aura":
		if attacker.Movedata[move].Poketype == "fairy"{
			multiplier = multiplier * 1.33
		}
	case defender.Ability == "fairy-aura":
		if attacker.Movedata[move].Poketype == "fairy"{
			multiplier = multiplier * 1.33
		}
	case attacker.Ability == "gorilla-tactics":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "huge-power":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 2.0
		}
	case attacker.Ability == "hustle":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "iron-fist":
			switch {
			case move == "bullet-punch":
				multiplier = multiplier * 1.2
				break	
			case move == "comet-punch":
				multiplier = multiplier * 1.2
				break
			case move == "dizzy-punch":
				multiplier = multiplier * 1.2
				break
			case move == "double-iron-bash":
				multiplier = multiplier * 1.2
				break	
			case move == "drain-punch":
				multiplier = multiplier * 1.2
				break
			case move == "dynamic-punch":
				multiplier = multiplier * 1.2
				break
			case move == "fire-punch":
				multiplier = multiplier * 1.2
				break
			case move == "focus-punch":
				multiplier = multiplier * 1.2
				break
			case move == "hammer-arm":
				multiplier = multiplier * 1.2
				break
			case move == "headlong-rush":
				multiplier = multiplier * 1.2
				break
			case move == "ice-hammer":
				multiplier = multiplier * 1.2
				break
			case move == "ice-punch":
				multiplier = multiplier * 1.2
				break
			case move == "jet-punch":
				multiplier = multiplier * 1.2
				break
			case move == "mach-punch":
				multiplier = multiplier * 1.2
				break
			case move == "mega-punch":
				multiplier = multiplier * 1.2
				break
			case move == "meteor-mash":
				multiplier = multiplier * 1.2
				break
			case move == "plasma-fists":
				multiplier = multiplier * 1.2
				break
			case move == "power-up-punch":
				multiplier = multiplier * 1.2
				break
			case move == "rage-fist":
				multiplier = multiplier * 1.2
				break
			case move == "shadow-punch":
				multiplier = multiplier * 1.2
				break
			case move == "sky-uppercut":
				multiplier = multiplier * 1.2
				break
			case move == "surging-strikes":
				multiplier = multiplier * 1.2
				break
			case move == "thunder-punch":
				multiplier = multiplier * 1.2
				break
			case move == "wicked-blow":
				multiplier = multiplier * 1.2
				break

		default:
			break
	}

	case attacker.Ability == "mega-launcher":
		switch{
		case move == "aura-sphere":
			multiplier = multiplier * 1.5
			break
		case move == "dark-pulse":
			multiplier = multiplier * 1.5
			break
		case move == "dragon-pulse":
			multiplier = multiplier * 1.5
			break
		case move == "origin-pulse":
			multiplier = multiplier * 1.5
			break
		case move == "terrain-pulse":
			multiplier = multiplier * 1.5
			break
		case move == "water-pulse":
			multiplier = multiplier * 1.5
			break
		default:
			break
		}
	case attacker.Ability == "neuroforce":
		mult := 1.0
		for _, Montype := range defender.Types {
			mult = mult * TypeChart[attacker.Movedata[move].Poketype][Montype]
		}
		if mult > 1.0 {
			multiplier = multiplier * 1.25
		}
	case attacker.Ability == "overgrow":
		if attacker.Movedata[move].Poketype == "grass"{
			if attacker.CurHp < attacker.CurStats["hp"]/3 {
				multiplier = multiplier * 1.5
			}
		}
	case attacker.Ability == "pixilate":
		if attacker.Movedata[move].Poketype == "normal"{
			multiplier = multiplier * 1.2
		}
	case attacker.Ability == "punk-rock":
		switch {
		case move == "boomburst":
			multiplier = multiplier * 1.3
			break
		case move == "hyper-voice":
			multiplier = multiplier * 1.3
			break
		case move == "overdrive":
			multiplier = multiplier * 1.3
			break
		case move == "psychic-noise":
			multiplier = multiplier * 1.3
			break
		case move == "round":
			multiplier = multiplier * 1.3
			break
		case move == "snarl":
			multiplier = multiplier * 1.3
			break
		case move == "snore":
			multiplier = multiplier * 1.3
			break
		case move == "uproar":
			multiplier = multiplier * 1.3
			break
		default:
			break
		}
	case attacker.Ability == "pure-power":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 2.0
		}
	case attacker.Ability == "reckless":
		switch{
		case move == "brave-bird":
			multiplier = multiplier * 1.2
			break
		case move == "double-edge":
			multiplier = multiplier * 1.2
			break
		case move == "flare-blitz":
			multiplier = multiplier * 1.2
			break
		case move == "head-charge":
			multiplier = multiplier * 1.2
			break
		case move == "head-smash":
			multiplier = multiplier * 1.2
			break
		case move == "high-jump-kick":
			multiplier = multiplier * 1.2
			break
		case move == "jump-kick":
			multiplier = multiplier * 1.2
			break
		case move == "light-of-ruin":
			multiplier = multiplier * 1.2
			break
		case move == "submission":
			multiplier = multiplier * 1.2
			break
		case move == "take-down":
			multiplier = multiplier * 1.2
			break
		case move == "volt-tackle":
			multiplier = multiplier * 1.2
			break
		case move == "wood-hammer":
			multiplier = multiplier * 1.2
			break
		case move == "wild-charge":
			multiplier = multiplier * 1.2
			break
		default:
			break
		}
	case attacker.Ability == "refrigerate":
		if attacker.Movedata[move].Poketype == "normal"{
			multiplier = multiplier * 1.2
		}
	case attacker.Ability == "rocky-payload":
		if attacker.Movedata[move].Poketype == "rock"{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "sharpness":
		switch{
		case move == "aerial-ace":
			multiplier = multiplier * 1.5
			break
		case move == "air-cutter":
			multiplier = multiplier * 1.5
			break
		case move == "air-slash":
			multiplier = multiplier * 1.5
			break
		case move == "aqua-cutter":
			multiplier = multiplier * 1.5
			break
		case move == "behemoth-blade":
			multiplier = multiplier * 1.5
			break
		case move == "bitter-blade":
			multiplier = multiplier * 1.5
			break
		case move == "ceaseless-edge":
			multiplier = multiplier * 1.5
			break
		case move == "cross-poison":
			multiplier = multiplier * 1.5
			break
		case move == "cut":
			multiplier = multiplier * 1.5
			break
		case move == "fury-cutter":
			multiplier = multiplier * 1.5
			break
		case move == "kowtow-cleave":
			multiplier = multiplier * 1.5
			break
		case move == "leaf-blade":
			multiplier = multiplier * 1.5
			break
		case move == "mighty-cleave":
			multiplier = multiplier * 1.5
			break
		case move == "night-slash":
			multiplier = multiplier * 1.5
			break
		case move == "population-bomb":
			multiplier = multiplier * 1.5
			break
		case move == "psyblade":
			multiplier = multiplier * 1.5
			break
		case move == "psycho-cut":
			multiplier = multiplier * 1.5
			break
		case move == "razor-leaf":
			multiplier = multiplier * 1.5
			break
		case move == "razor-shell":
			multiplier = multiplier * 1.5
			break
		case move == "sacred-sword":
			multiplier = multiplier * 1.5
			break
		case move == "secret-sword":
			multiplier = multiplier * 1.5
			break
		case move == "slash":
			multiplier = multiplier * 1.5
			break
		case move == "solar-blade":
			multiplier = multiplier * 1.5
			break
		case move == "stone-axe":
			multiplier = multiplier * 1.5
			break
		case move == "tachyon-cutter":
			multiplier = multiplier * 1.5
			break
		case move == "x-scissor":
			multiplier = multiplier * 1.5
			break
		default:
			break
		}
	case attacker.Ability == "steelworker":
		if attacker.Movedata[move].Poketype == "steel"{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "strong-jaw":
		switch{
		case move == "bite":
			multiplier = multiplier * 1.6
			break
		case move == "crunch":
			multiplier = multiplier * 1.5
			break
		case move == "fire-fang":
			multiplier = multiplier * 1.5
			break
		case move == "fishious-rend":
			multiplier = multiplier * 1.5
			break
		case move == "hyper-fang":
			multiplier = multiplier * 1.5
			break
		case move == "ice-fang":
			multiplier = multiplier * 1.5
			break
		case move == "jaw-lock":
			multiplier = multiplier * 1.5
			break
		case move == "poison-fang":
			multiplier = multiplier * 1.5
			break
		case move == "psychic-fangs":
			multiplier = multiplier * 1.5
			break
		case move == "thunder-fang":
			multiplier = multiplier * 1.5
			break
		default:
			break
		}
	case attacker.Ability == "swarm":
		if attacker.Movedata[move].Poketype == "bug"{
			if attacker.CurHp < attacker.CurStats["hp"]/3 {
				multiplier = multiplier * 1.5
			}
		}
	case attacker.Ability == "technician":
		p, err := strconv.Atoi(attacker.Movedata[move].Power)
		if err != nil {
			fmt.Println("error with conversion")
			return multiplier
		}
		if p <= 60 {
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "tinted-lens":
		mult := 1.0
		for _, Montype := range defender.Types {
			mult = mult * TypeChart[attacker.Movedata[move].Poketype][Montype]
		}
		if mult < 1.0 {
			multiplier = multiplier * 2
		}
	case attacker.Ability == "torrent":
		if attacker.Movedata[move].Poketype == "water"{
			if attacker.CurHp < attacker.CurStats["hp"]/3 {
				multiplier = multiplier * 1.5
			}
		}
	case attacker.Ability == "transistor":
		if attacker.Movedata[move].Poketype == "electric"{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "water-bubble":
		if attacker.Movedata[move].Poketype == "water"{
			multiplier = multiplier * 2.0
		}
	default:
		break
	}
	return multiplier
}