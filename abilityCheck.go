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
	case attacker.Ability == "aura-break":
		if attacker.Movedata[move].Poketype == "dark" || attacker.Movedata[move].Poketype == "fairy"{
			multiplier = multiplier * 0.75
		}
	case defender.Ability == "aura-break":
		if attacker.Movedata[move].Poketype == "dark" || attacker.Movedata[move].Poketype == "fairy"{
			multiplier = multiplier * 0.75
		}
	case attacker.Ability == "blaze":
		if attacker.Movedata[move].Poketype == "fire"{
			if attacker.CurHp < attacker.CurStats["hp"]/3 {
				multiplier = multiplier * 1.5
			}
		}
	case attacker.Ability == "beads-of-ruin":
		if attacker.Movedata[move].Damagetype == "special"{
			multiplier = multiplier * 1.25
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
	case defender.Ability == "fluffy":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 0.5
		}
		if attacker.Movedata[move].Poketype == "fire"{
			multiplier = multiplier * 2
		}
	case defender.Ability == "fur-coat":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 0.5
		}
	case attacker.Ability == "gorilla-tactics":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 1.5
		}
	case defender.Ability == "heatproof":
		if attacker.Movedata[move].Poketype == "fire"{
			multiplier = multiplier * 0.5
		}
	case attacker.Ability == "huge-power":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 2.0
		}
	case attacker.Ability == "hustle":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 1.5
		}
	case defender.Ability == "ice-scales":
		if attacker.Movedata[move].Damagetype == "special"{
			multiplier = multiplier * 0.5
		}
	case attacker.Ability == "iron-fist":
		if fistMoves[move]{
			multiplier = multiplier * 1.2
		}
	case attacker.Ability == "mega-launcher":
		if launcherMoves[move] {
			multiplier = multiplier * 1.5
		}
	case defender.Ability == "multiscale" || defender.Ability == "shadow-shield":
		if defender.CurHp == defender.CurStats["hp"]{
			multiplier = multiplier * 0.5
		}
	case attacker.Ability == "neuroforce":
		mult := 1.0
		for _, Montype := range defender.Types {
			mult = mult * TypeChart[attacker.Movedata[move].Poketype][Montype]
		}
		if mult > 1.0 {
			multiplier = multiplier * 1.25
		}
	case attacker.Ability == "normalize":
		multiplier = multiplier * 1.2
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
		if soundMoves[move]{
			multiplier = multiplier * 1.3
		}
	case defender.Ability == "punk-rock":
		if soundMoves[move]{
			multiplier = multiplier * 0.5
		}
	case attacker.Ability == "pure-power":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 2.0
		}
	case defender.Ability == "prism-armor" || defender.Ability == "solid-rock":
		mult := 1.0
		for _, Montype := range defender.Types {
			mult = mult * TypeChart[attacker.Movedata[move].Poketype][Montype]
		}
		if mult > 1.0 {
			multiplier = multiplier * 0.75
		}
	case attacker.Ability == "reckless":
		if recoilMoves[move]{
			multiplier = multiplier * 1.2
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
		if sharpnessMoves[move]{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "steelworker":
		if attacker.Movedata[move].Poketype == "steel"{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "strong-jaw":
		if biteMoves[move]{
			multiplier = multiplier * 1.5
		}
	case attacker.Ability == "swarm":
		if attacker.Movedata[move].Poketype == "bug"{
			if attacker.CurHp < attacker.CurStats["hp"]/3 {
				multiplier = multiplier * 1.5
			}
		}
	case defender.Ability == "tablets-of-ruin":
		if attacker.Movedata[move].Damagetype == "physical"{
			multiplier = multiplier * 0.75
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
	case defender.Ability == "thick-fat":
		if attacker.Movedata[move].Poketype == "ice" || attacker.Movedata[move].Poketype == "fire"{
			multiplier = multiplier * 0.5
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
	case defender.Ability == "vessel-of-ruin":
		if attacker.Movedata[move].Damagetype == "special"{
			multiplier = multiplier * 0.75
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

func checkImmunityAbility(move string, defender Pokemon, attacker Pokemon) float64 {
	switch {
	case attacker.Ability == "neutralizing-gas":
		return 1.0
	case defender.Ability == "neutralizing-gas":
		return 1.0
	case attacker.Ability == "mold-breaker" || attacker.Ability == "teravolt":
		return 1.0
	case defender.Ability == "armor-tail" || defender.Ability == "dazzling" || defender.Ability == "queenly-majesty":
		if attacker.Movedata[move].Priority > 0 {
			return 0.0
		}
	case defender.Ability == "bulletproof":
		if ballAndBomb[move]{
			return 0.0
		}
	case defender.Ability == "damp":
		if move == "selfdestruct" || move == "explosion"{
			return 0.0
		}
	case defender.Ability == "earth-eater":
		if attacker.Movedata[move].Poketype == "ground"{
			return 0.25
		}
	case defender.Ability == "flash-fire":
		if attacker.Movedata[move].Poketype == "fire"{
			//Add logic to boost fire type moves in the future
			return 0.0
		}
	//Consider adding disguise and Ice Face here later
	case defender.Ability == "levitate":
		if attacker.Movedata[move].Poketype == "ground"{
			return 0.0
		}
	case defender.Ability == "motor-drive":
		if attacker.Movedata[move].Poketype == "electric"{
			//add speed raise to defender
			return 0.0
		}
	case defender.Ability == "sap-sipper":
		if attacker.Movedata[move].Poketype == "grass"{
			//Add attack to defender
			return 0.0
		}
	case defender.Ability == "storm-drain":
		if attacker.Movedata[move].Poketype == "water"{
			//Add sp. attack to defender
			return 0.0
		}
	case defender.Ability == "volt-absorb":
		if attacker.Movedata[move].Poketype == "electric"{
			return 0.25
		}
	case defender.Ability == "water-absorb":
		if attacker.Movedata[move].Poketype == "water"{
			return 0.25
		}
	case defender.Ability == "well-baked-body":
		if attacker.Movedata[move].Poketype == "fire"{
			//Raise defender's defense by 2 stages
			return 0.0
		}
	case defender.Ability == "wonder-guard":
		mult := 1.0
		for _, Montype := range defender.Types {
			mult = mult * TypeChart[attacker.Movedata[move].Poketype][Montype]
		}
		if mult > 1.0 {
			return 1.0
		} else {
			return 0.0
		}
	default:
		return 1.0
	}
	return 1.0
}

var ballAndBomb = map[string]bool{
    "acid-spray":   true,
    "aura-sphere":  true,
    "barrage":      true,
    "beak-blast":   true,
    "bullet-seed":  true,
    "egg-bomb":     true,
    "electro-ball": true,
    "energy-ball":  true,
    "focus-blast":  true,
    "gyro-ball":    true,
    "ice-ball":     true,
    "magnet-bomb":  true,
    "mist-ball":    true,
    "mud-bomb":     true,
    "octazooka":    true,
    "pollen-puff":  true,
    "rock-blast":   true,
    "rock-wrecker": true,
    "searing-shot": true,
    "seed-bomb":    true,
    "shadow-ball":  true,
    "sludge-bomb":  true,
    "weather-ball": true,
    "zap-cannon":   true,
}

var biteMoves = map[string]bool{
    "bite":            true,
    "crunch":          true,
    "fire-fang":       true,
    "fishious-rend":   true,
    "hyper-fang":      true,
    "ice-fang":        true,
    "jaw-lock":        true,
    "poison-fang":     true,
    "psychic-fangs":   true,
    "thunder-fang":    true,
}

var sharpnessMoves = map[string]bool{
    "aerial-ace":      true,
    "air-cutter":      true,
    "air-slash":       true,
    "aqua-cutter":     true,
    "behemoth-blade":  true,
    "bitter-blade":    true,
    "ceaseless-edge":  true,
    "cross-poison":    true,
    "cut":             true,
    "fury-cutter":     true,
    "kowtow-cleave":   true,
    "leaf-blade":      true,
    "mighty-cleave":   true,
    "night-slash":     true,
    "population-bomb": true,
    "psyblade":        true,
    "psycho-cut":      true,
    "razor-leaf":      true,
    "razor-shell":     true,
    "sacred-sword":    true,
    "secret-sword":    true,
    "slash":           true,
    "solar-blade":     true,
    "stone-axe":       true,
    "tachyon-cutter":  true,
    "x-scissor":       true,
}

var recoilMoves = map[string]bool{
    "brave-bird":       true,
    "double-edge":      true,
    "flare-blitz":      true,
    "head-charge":      true,
    "head-smash":       true,
    "high-jump-kick":   true,
    "jump-kick":        true,
    "light-of-ruin":    true,
    "submission":       true,
    "take-down":        true,
    "volt-tackle":      true,
    "wood-hammer":      true,
    "wild-charge":      true,
}

var soundMoves = map[string]bool{
    "boomburst":      true,
    "hyper-voice":    true,
    "overdrive":      true,
    "psychic-noise":  true,
    "round":          true,
    "snarl":          true,
    "snore":          true,
    "uproar":         true,
}

var launcherMoves = map[string]bool{
    "aura-sphere":   true,
    "dark-pulse":    true,
    "dragon-pulse":  true,
    "origin-pulse":  true,
    "terrain-pulse": true,
    "water-pulse":   true,
}

var fistMoves = map[string]bool{
    "bullet-punch":     true,
    "comet-punch":      true,
    "dizzy-punch":      true,
    "double-iron-bash": true,
    "drain-punch":      true,
    "dynamic-punch":    true,
    "fire-punch":       true,
    "focus-punch":      true,
    "hammer-arm":       true,
    "headlong-rush":    true,
    "ice-hammer":       true,
    "ice-punch":        true,
    "jet-punch":        true,
    "mach-punch":       true,
    "mega-punch":       true,
    "meteor-mash":      true,
    "plasma-fists":     true,
    "power-up-punch":   true,
    "rage-fist":        true,
    "shadow-punch":     true,
    "sky-uppercut":     true,
    "surging-strikes":  true,
    "thunder-punch":    true,
    "wicked-blow":      true,
}
