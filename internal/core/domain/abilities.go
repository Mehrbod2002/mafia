package domain

// AbilityOption represents a selectable ability for a role/character.
type AbilityOption struct {
	Code               string `json:"code"`
	Label              string `json:"label"`
	Phase              string `json:"phase"`
	Side               string `json:"side"`
	PersianLabel       string `json:"persian_label"`
	PersianDescription string `json:"persian_description"`
	Description        string `json:"description"`
	Icon               string `json:"icon"`
}

// AbilityCatalog returns the curated list of available abilities for admins.
func AbilityCatalog() []AbilityOption {
	return []AbilityOption{
		{Code: "angel_of_salvation", Label: "Angel of Salvation", PersianLabel: "فرشته نجات", Phase: "both", Side: "town", Icon: "🪽", Description: "Saves the player with the most votes from elimination when used correctly.", PersianDescription: "بازیکنی که بیشترین رای را برای حذف شدن دارد نجات می‌دهد در صورتی که به درستی استفاده شود."},
		{Code: "grave_digger", Label: "Grave Digger", PersianLabel: "گورکن", Phase: "both", Side: "town", Icon: "⛏️", Description: "Investigates a dead player's role and alignment.", PersianDescription: "نقش و سمت بازیکن مرده را بررسی می‌کند."},
		{Code: "gambler", Label: "Gambler", PersianLabel: "قمارباز", Phase: "both", Side: "town", Icon: "🎲", Description: "Boosts another player's ability effectiveness for the next night.", PersianDescription: "اثربخشی توانایی یک بازیکن دیگر را برای شب بعد تقویت می‌کند."},
		{Code: "fortune_teller", Label: "Fortune Teller", PersianLabel: "فالگیر", Phase: "both", Side: "town", Icon: "🔮", Description: "Communicates with a dead player to reveal their role and side.", PersianDescription: "با یک بازیکن مرده ارتباط برقرار کرده و نقش و سمت آن را فاش می‌کند."},
		{Code: "innocent", Label: "Innocent", PersianLabel: "بی‌گناه", Phase: "night", Side: "town", Icon: "🧒", Description: "If attacked by Mafia, exposes the killer the following day.", PersianDescription: "اگر توسط مافیا حمله شود، قاتل را روز بعد فاش می‌کند."},
		{Code: "lawyer", Label: "Lawyer", PersianLabel: "وکیل", Phase: "day", Side: "town", Icon: "⚖️", Description: "Prevents a town player from being eliminated by votes.", PersianDescription: "از حذف یک بازیکن شهروندی توسط رای جلوگیری می‌کند."},
		{Code: "self_sacrificing", Label: "Self-Sacrificing", PersianLabel: "فداکار", Phase: "day", Side: "town", Icon: "🩸", Description: "Sacrifices themselves to stop a Mafia shot against another player.", PersianDescription: "برای متوقف کردن تیر مافیا به بازیکن دیگر، خود را فدای او می‌کند."},
		{Code: "sergeant", Label: "Sergeant", PersianLabel: "سرهنگ", Phase: "day", Side: "town", Icon: "🛡️", Description: "Upgrades a player's weapon for a stronger shot.", PersianDescription: "سلاح یک بازیکن را ارتقا می‌دهد تا تیر قوی‌تری بزند."},
		{Code: "ranger", Label: "Ranger", PersianLabel: "رِینجر", Phase: "day", Side: "town", Icon: "🏹", Description: "Can make a shot and blocks the first Mafia shot against them.", PersianDescription: "می‌تواند شلیک کند و اولین تیر مافیا را در برابر خود مسدود می‌کند."},
		{Code: "hunter", Label: "Hunter", PersianLabel: "شکارچی", Phase: "both", Side: "town", Icon: "🎯", Description: "Can eliminate threats such as Werewolves under special conditions.", PersianDescription: "می‌تواند تهدیدهایی مانند گرگینه‌ها را تحت شرایط خاص از بین ببرد."},
		{Code: "pope", Label: "Pope", PersianLabel: "پاپ", Phase: "day", Side: "town", Icon: "✝️", Description: "Eliminates the Devil from the game when targeted.", PersianDescription: "شیطان را از بازی حذف می‌کند وقتی هدف قرار گیرد."},
		{Code: "godfather", Label: "Godfather", PersianLabel: "کاپو", Phase: "both", Side: "mafia", Icon: "👔", Description: "Commands priority to make the first shot.", PersianDescription: "اولویت برای شلیک اول را دارد."},
		{Code: "nato", Label: "Nato", PersianLabel: "ناتو", Phase: "both", Side: "mafia", Icon: "🎖️", Description: "Executes a coordinated elimination with Mafia support.", PersianDescription: "حذف هماهنگ‌شده با حمایت مافیا انجام می‌دهد."},
		{Code: "kidnapper", Label: "Kidnapper", PersianLabel: "قاچاقچی", Phase: "both", Side: "mafia", Icon: "🪤", Description: "Disables a selected player's abilities for the phase.", PersianDescription: "توانایی‌های یک بازیکن انتخابی را برای فاز غیرفعال می‌کند."},
		{Code: "simple_mafia", Label: "Simple Mafia", PersianLabel: "مافیا ساده", Phase: "night", Side: "mafia", Icon: "🕴️", Description: "Standard Mafia member who collaborates without extra powers.", PersianDescription: "عضو معمولی مافیا که بدون قدرت اضافی همکاری می‌کند."},
		{Code: "negotiator", Label: "Negotiator", PersianLabel: "مذاکره‌کننده", Phase: "both", Side: "mafia", Icon: "🤝", Description: "Attempts to convert a town player to Mafia if chosen correctly.", PersianDescription: "اگر به درستی انتخاب شود، سعی می‌کند یک بازیکن شهروندی را به مافیا تبدیل کند."},
		{Code: "imposter", Label: "Imposter", PersianLabel: "جعل‌کار", Phase: "both", Side: "mafia", Icon: "🎭", Description: "Alters investigations to return false results.", PersianDescription: "تحقیقات را تغییر می‌دهد تا نتایج نادرست برگرداند."},
		{Code: "dr_lecter", Label: "Dr. Lecter", PersianLabel: "دکتر لِکتِر", Phase: "both", Side: "mafia", Icon: "🩺", Description: "Saves a Mafia member from elimination by players.", PersianDescription: "یک عضو مافیا را از حذف توسط بازیکنان نجات می‌دهد."},
		{Code: "natasha", Label: "Natasha", PersianLabel: "ناتاشا", Phase: "both", Side: "mafia", Icon: "🤫", Description: "Silences a player, removing their speech and vote for the day.", PersianDescription: "یک بازیکن را ساکت کرده و صحبت و رای آن‌ها را برای روز حذف می‌کند."},
		{Code: "terrorist", Label: "Terrorist", PersianLabel: "تروریست", Phase: "both", Side: "mafia", Icon: "💣", Description: "Immediately eliminates a target player during the phase.", PersianDescription: "یک بازیکن هدف را به‌طور فوری در طول فاز حذف می‌کند."},
		{Code: "bomb_maker", Label: "Bomb Maker", PersianLabel: "بمب‌ساز", Phase: "both", Side: "mafia", Icon: "🧨", Description: "Plants a bomb that must be disarmed by guessing the correct number.", PersianDescription: "یک بمب می‌کارد که باید با حدس زدن شماره صحیح خنثی شود."},
		{Code: "poisoner", Label: "Poisoner", PersianLabel: "مسمم‌کننده", Phase: "both", Side: "mafia", Icon: "☠️", Description: "Poisons a player who will die the following day if not cured.", PersianDescription: "یک بازیکن را مسموم می‌کند که اگر درمان نشود، روز بعد می‌میرد."},
		{Code: "mistress", Label: "Mistress", PersianLabel: "معشوقه", Phase: "both", Side: "mafia", Icon: "💋", Description: "On elimination grants Mafia an additional shot.", PersianDescription: "هنگام حذف، یک شلیک اضافی به مافیا می‌دهد."},
		{Code: "dynamite_maker", Label: "Dynamite Maker", PersianLabel: "دینامیت‌ساز", Phase: "both", Side: "mafia", Icon: "🧯", Description: "Sets dynamite with alternate disarm mechanics compared to bombs.", PersianDescription: "دینامیت را با مکانیزم‌های خنثی‌سازی متفاوت از بمب‌ها تنظیم می‌کند."},
		{Code: "thief", Label: "Thief", PersianLabel: "دزد", Phase: "both", Side: "mafia", Icon: "🗡️", Description: "Steals another player's abilities.", PersianDescription: "توانایی‌های یک بازیکن دیگر را می‌دزدد."},
		{Code: "informer", Label: "Informer", PersianLabel: "اطلاع‌رسان", Phase: "both", Side: "mafia", Icon: "📡", Description: "Learns a town player's role.", PersianDescription: "نقش یک بازیکن شهروندی را می‌آموزد."},
		{Code: "protector", Label: "Protector", PersianLabel: "حفاظت‌کننده", Phase: "both", Side: "town", Icon: "🛡️", Description: "Shields a player from Mafia attacks.", PersianDescription: "یک بازیکن را از حملات مافیا محافظت می‌کند."},
		{Code: "swayer", Label: "Swayer", PersianLabel: "سایر", Phase: "both", Side: "mafia", Icon: "🗳️", Description: "Manipulates votes or targeted abilities.", PersianDescription: "رای‌ها یا توانایی‌های هدف‌گذاری شده را دستکاری می‌کند."},
		{Code: "hacker", Label: "Hacker", PersianLabel: "هکر", Phase: "both", Side: "mafia", Icon: "💻", Description: "Reveals town players and their roles.", PersianDescription: "بازیکنان شهروندی و نقش‌های آن‌ها را فاش می‌کند."},
		{Code: "mercenary", Label: "Mercenary", PersianLabel: "مزدور", Phase: "both", Side: "mafia", Icon: "🪖", Description: "Adjusts investigation results to favor their side.", PersianDescription: "نتایج تحقیقات را به نفع خود تغییر می‌دهد."},
		{Code: "adopted_daughter", Label: "Adopted Daughter", PersianLabel: "دختر خوانده", Phase: "both", Side: "mafia", Icon: "👧", Description: "Converts a town player into Mafia when executed correctly.", PersianDescription: "یک بازیکن شهروندی را به مافیا تبدیل می‌کند وقتی به درستی اجرا شود."},
		{Code: "adopted_son", Label: "Adopted Son", PersianLabel: "پسر خوانده", Phase: "both", Side: "mafia", Icon: "👦", Description: "Assumes a Mafia role after two Mafia members fall.", PersianDescription: "پس از سقوط دو عضو مافیا، نقش مافیا را به عهده می‌گیرد."},
		{Code: "nostradamus", Label: "Nostradamus", PersianLabel: "نستراداموس", Phase: "both", Side: "neutral", Icon: "🌙", Description: "Chooses a side based on the game's progression and predicts the winner.", PersianDescription: "یک سمت را بر اساس پیشرفت بازی انتخاب کرده و برنده را پیش‌بینی می‌کند."},
		{Code: "thousand_faces", Label: "Thousand Faces", PersianLabel: "هزار چهره", Phase: "both", Side: "neutral", Icon: "🌀", Description: "Steals a player's role after they are eliminated.", PersianDescription: "نقش یک بازیکن را پس از حذف آن‌ها می‌دزدد."},
		{Code: "assassin", Label: "Assassin", PersianLabel: "کشنده", Phase: "night", Side: "mafia", Icon: "🗡", Description: "Eliminates a player on even nights.", PersianDescription: "یک بازیکن را در شب‌های زوج حذف می‌کند."},
		{Code: "sherlock_holmes", Label: "Sherlock Holmes", PersianLabel: "شرلوك هولمز", Phase: "night", Side: "town", Icon: "🕵️", Description: "Guesses a player's role and may swap sides if correct.", PersianDescription: "نقش یک بازیکن را حدس زده و در صورت درست بودن، ممکن است سمت‌ها را عوض کند."},
		{Code: "saqi", Label: "Saqi", PersianLabel: "ساقی", Phase: "night", Side: "neutral", Icon: "🍷", Description: "Intoxicates a player, suppressing their abilities temporarily.", PersianDescription: "یک بازیکن را مست کرده و توانایی‌های آن را موقتاً سرکوب می‌کند."},
		{Code: "clumsy_hand", Label: "Clumsy Hand", PersianLabel: "دست اشتباهی", Phase: "night", Side: "neutral", Icon: "✋", Description: "Disables a player's abilities for the remainder of the game.", PersianDescription: "توانایی‌های یک بازیکن را برای باقی‌مانده بازی غیرفعال می‌کند."},
	}
}

// AbilityIndex maps ability codes to their full option definitions for quick lookup.
func AbilityIndex() map[string]AbilityOption {
	idx := make(map[string]AbilityOption)
	for _, ability := range AbilityCatalog() {
		idx[ability.Code] = ability
	}
	return idx
}

// FindAbility fetches the ability option for a code if it exists.
func FindAbility(code string) (AbilityOption, bool) {
	ability, ok := AbilityIndex()[code]
	return ability, ok
}
