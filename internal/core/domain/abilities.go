package domain

// AbilityOption represents a selectable ability for a role/character.
type AbilityOption struct {
	Code        string `json:"code"`
	Label       string `json:"label"`
	Phase       string `json:"phase"`
	Side        string `json:"side"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// AbilityCatalog returns the curated list of available abilities for admins.
func AbilityCatalog() []AbilityOption {
	return []AbilityOption{
		{Code: "angel_of_salvation", Label: "Angel of Salvation", Phase: "both", Side: "town", Icon: "🪽", Description: "Saves the player with the most votes from elimination when used correctly."},
		{Code: "grave_digger", Label: "Grave Digger", Phase: "both", Side: "town", Icon: "⛏️", Description: "Investigates a dead player's role and alignment."},
		{Code: "gambler", Label: "Gambler", Phase: "both", Side: "town", Icon: "🎲", Description: "Boosts another player's ability effectiveness for the next night."},
		{Code: "fortune_teller", Label: "Fortune Teller", Phase: "both", Side: "town", Icon: "🔮", Description: "Communicates with a dead player to reveal their role and side."},
		{Code: "innocent", Label: "Innocent", Phase: "night", Side: "town", Icon: "🧒", Description: "If attacked by Mafia, exposes the killer the following day."},
		{Code: "lawyer", Label: "Lawyer", Phase: "day", Side: "town", Icon: "⚖️", Description: "Prevents a town player from being eliminated by votes."},
		{Code: "self_sacrificing", Label: "Self-Sacrificing", Phase: "day", Side: "town", Icon: "🩸", Description: "Sacrifices themselves to stop a Mafia shot against another player."},
		{Code: "sergeant", Label: "Sergeant", Phase: "day", Side: "town", Icon: "🛡️", Description: "Upgrades a player's weapon for a stronger shot."},
		{Code: "ranger", Label: "Ranger", Phase: "day", Side: "town", Icon: "🏹", Description: "Can make a shot and blocks the first Mafia shot against them."},
		{Code: "hunter", Label: "Hunter", Phase: "both", Side: "town", Icon: "🎯", Description: "Can eliminate threats such as Werewolves under special conditions."},
		{Code: "pope", Label: "Pope", Phase: "day", Side: "town", Icon: "✝️", Description: "Eliminates the Devil from the game when targeted."},
		{Code: "godfather", Label: "Godfather", Phase: "both", Side: "mafia", Icon: "👔", Description: "Commands priority to make the first shot."},
		{Code: "nato", Label: "Nato", Phase: "both", Side: "mafia", Icon: "🎖️", Description: "Executes a coordinated elimination with Mafia support."},
		{Code: "kidnapper", Label: "Kidnapper", Phase: "both", Side: "mafia", Icon: "🪤", Description: "Disables a selected player's abilities for the phase."},
		{Code: "simple_mafia", Label: "Simple Mafia", Phase: "night", Side: "mafia", Icon: "🕴️", Description: "Standard Mafia member who collaborates without extra powers."},
		{Code: "negotiator", Label: "Negotiator", Phase: "both", Side: "mafia", Icon: "🤝", Description: "Attempts to convert a town player to Mafia if chosen correctly."},
		{Code: "imposter", Label: "Imposter", Phase: "both", Side: "mafia", Icon: "🎭", Description: "Alters investigations to return false results."},
		{Code: "dr_lecter", Label: "Dr. Lecter", Phase: "both", Side: "mafia", Icon: "🩺", Description: "Saves a Mafia member from elimination by players."},
		{Code: "natasha", Label: "Natasha", Phase: "both", Side: "mafia", Icon: "🤫", Description: "Silences a player, removing their speech and vote for the day."},
		{Code: "terrorist", Label: "Terrorist", Phase: "both", Side: "mafia", Icon: "💣", Description: "Immediately eliminates a target player during the phase."},
		{Code: "bomb_maker", Label: "Bomb Maker", Phase: "both", Side: "mafia", Icon: "🧨", Description: "Plants a bomb that must be disarmed by guessing the correct number."},
		{Code: "poisoner", Label: "Poisoner", Phase: "both", Side: "mafia", Icon: "☠️", Description: "Poisons a player who will die the following day if not cured."},
		{Code: "mistress", Label: "Mistress", Phase: "both", Side: "mafia", Icon: "💋", Description: "On elimination grants Mafia an additional shot."},
		{Code: "dynamite_maker", Label: "Dynamite Maker", Phase: "both", Side: "mafia", Icon: "🧯", Description: "Sets dynamite with alternate disarm mechanics compared to bombs."},
		{Code: "thief", Label: "Thief", Phase: "both", Side: "mafia", Icon: "🗡️", Description: "Steals another player's abilities."},
		{Code: "informer", Label: "Informer", Phase: "both", Side: "mafia", Icon: "📡", Description: "Learns a town player's role."},
		{Code: "protector", Label: "Protector", Phase: "both", Side: "town", Icon: "🛡️", Description: "Shields a player from Mafia attacks."},
		{Code: "swayer", Label: "Swayer", Phase: "both", Side: "mafia", Icon: "🗳️", Description: "Manipulates votes or targeted abilities."},
		{Code: "hacker", Label: "Hacker", Phase: "both", Side: "mafia", Icon: "💻", Description: "Reveals town players and their roles."},
		{Code: "mercenary", Label: "Mercenary", Phase: "both", Side: "mafia", Icon: "🪖", Description: "Adjusts investigation results to favor their side."},
		{Code: "adopted_daughter", Label: "Adopted Daughter", Phase: "both", Side: "mafia", Icon: "👧", Description: "Converts a town player into Mafia when executed correctly."},
		{Code: "adopted_son", Label: "Adopted Son", Phase: "both", Side: "mafia", Icon: "👦", Description: "Assumes a Mafia role after two Mafia members fall."},
		{Code: "nostradamus", Label: "Nostradamus", Phase: "both", Side: "neutral", Icon: "🌙", Description: "Chooses a side based on the game's progression and predicts the winner."},
		{Code: "thousand_faces", Label: "Thousand Faces", Phase: "both", Side: "neutral", Icon: "🌀", Description: "Steals a player's role after they are eliminated."},
		{Code: "assassin", Label: "Assassin", Phase: "night", Side: "mafia", Icon: "🗡", Description: "Eliminates a player on even nights."},
		{Code: "sherlock_holmes", Label: "Sherlock Holmes", Phase: "night", Side: "town", Icon: "🕵️", Description: "Guesses a player's role and may swap sides if correct."},
		{Code: "saqi", Label: "Saqi", Phase: "night", Side: "neutral", Icon: "🍷", Description: "Intoxicates a player, suppressing their abilities temporarily."},
		{Code: "clumsy_hand", Label: "Clumsy Hand", Phase: "night", Side: "neutral", Icon: "✋", Description: "Disables a player's abilities for the remainder of the game."},
	}
}
