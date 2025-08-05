package commands

import (
	"fmt"
)

func commandParty(cfg *config) error {
	if len(cfg.Party) == 0 {
		return fmt.Errorf("You have no Pokemon in your party!")
	}
	for i := range cfg.PokeKeys{
		if i == 0{
			fmt.Printf("Lead: ")
		}
		mon := cfg.Party[cfg.PokeKeys[i]]
		if mon.NickName != mon.Name {
			fmt.Printf("%v the %v at level %v. HP: %v/%v\n", mon.NickName, mon.Name, mon.Level, mon.CurHp, mon.CurStats["hp"])
		} else {
			fmt.Printf("%v at level %v. HP: %v/%v\n", mon.Name, mon.Level, mon.CurHp, mon.CurStats["hp"])
		}
		fmt.Printf("EXP to next level: %v\n", expChart[mon.ExpGroup][mon.Level+1]-mon.TotalExp)
	}
	return nil
}