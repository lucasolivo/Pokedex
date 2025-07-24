Welcome to the Command Line Pokedex! To start up the Pokedex, run "go run ." and the Pokedex will fire up!
You can call the 'help' command to view all the possible commands in the Pokedex. 
More features are being added soon...

Data is sourced through PokeAPI.

Update 5/23: Added the ability (x) command. generates 1-10 random abilities to use on a pokemon! Use this to randomize pokemon on pokehax for fun or just to see some abilities and their effects.

Also added levels when catching a pokemon and added a command to level a pokemon by 1 ('candy')

Update 6/11: Added a 'party' command that lets the user see their current party of 6 and their leading pokemon. Will add a box and switch and switchlead command later.

Update 6/12: Added a save and load feature along with a reset option, so that the Pokedex and party persist after the program is exited. 
Added a makeCommands function to clean up the repl.go file.
Added the encounter command.

Update 6/16: Added the 'stats' command to show the current stats of a pokemon and logic to calculate the current stats of a pokemon based on level (natures and evs/ivs not implemented). Logic is based on current gen calculations.

Update 6/17: Added abilities to Pokemon that you catch, a random one is assigned, even hidden abilities!
Added nickname support to get closer to your Pokemon.
Added command to box pokemon.
Added logic to not allow duplicate Pokemon to be caught.
Added the add command to add pokemon from your box into your party
Added the switch command to switch Pokemon in your party

Update 6/19: Added level up movesets and the move command to check a Pokemon's moves.
Implemented teach command and logic to prevent duplicate moves

Update 7/22: Added MoveData to Pokemon, allowing the power, accuracy, type, and damage type of pokemon moves to be seen.
Changed the library used to search for input in repl.go to support up and down arrows for previous commands.

Update 7/24: Added a CurHP value to the Pokemon struct, along with printing out HP in a "Cur/Max" format with the stats command