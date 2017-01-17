package d2s

// Item represents a base 111 bit item.
// Item represents an actual item
type Item struct {
	Identified         uint64
	Socketed           uint64
	New                uint64
	IsEar              uint64
	StarterItem        uint64
	SimpleItem         uint64
	Ethereal           uint64
	Personalized       uint64
	GivenRuneword      uint64
	LocationID         uint64
	EquippedID         uint64
	PositionY          uint64
	PositionX          uint64
	AltPositionID      uint64
	Type               string
	TypeID             uint64
	NrOfItemsInSockets uint64
	ID                 uint64
	Level              uint64
	Quality            uint64
	MultiplePictures   uint64
	PictureID          uint64
	ClassSpecific      uint64
	LowQualityID       uint64
	Timestamp          uint64
	DefenseRating      uint64
	MaxDurability      uint64
	CurrentDurability  uint64
	TotalNrOfSockets   uint64
	Quantity           uint64

	// Magical Item properties
	MagicPrefix uint64
	MagicSuffix uint64

	// Set item properties
	SetID        uint64
	SetItemLists uint64

	// All item types >= magicallyEnhanced
	MagicAttributes []magicAttribute
}

func (i Item) getTypeID() uint64 {

	if _, ok := armorCodes[i.Type]; ok {
		return armor
	}

	if _, ok := weaponCodes[i.Type]; ok {
		return weapon
	}

	return other
}

// Item types.
const (
	armor  = 0x01
	weapon = 0x02
	other  = 0x03
)

// Item location IDs.
const (
	stored   = 0x00
	equipped = 0x01
	belt     = 0x02
	cursor   = 0x03
	socketed = 0x04
)

// Rarity IDs.
const (
	lowQuality        = 0x01
	normal            = 0x02
	highQuality       = 0x03
	magicallyEnhanced = 0x04
	partOfSet         = 0x05
	rare              = 0x06
	unique            = 0x07
	crafted           = 0x08
)

// Jewelry and misc codes.
const (
	Ring = "rin"
)

// Armor codes.
var armorCodes = map[string]string{
	"uow": "Aegis",
	"pa4": "Aerin Shield",
	"pa7": "Akaran Rondache",
	"pa6": "Akaran Targe",
	"dr6": "Alpha Helm",
	"aar": "Ancient Armor",
	"pad": "Ancient Shield",
	"xts": "Ancient Shield",
	"dr3": "Antlers",
	"utp": "Archon Plate",
	"ulm": "Armet",
	"ba4": "Assault Helmet",
	"ba5": "Avenger Guard",
	"upl": "Balrog Skin",
	"xpk": "Barbed Shield",
	"xhl": "Basinet",
	"ztb": "Battle Belt",
	"xtb": "Battle Boots",
	"xtg": "Battle Gauntlets",
	"mbl": "Belt",
	"upk": "Blade Barrier",
	"drb": "Blood Spirit",
	"nef": "Bloodlord Skull",
	"bhm": "Bone Helm",
	"bsh": "Bone Shield",
	"uh9": "Bone Visage",
	"uhn": "Boneweave",
	"umb": "Boneweave Boots",
	"mgl": "Bracers",
	"ulg": "Bramble Mitts",
	"brs": "Breast Plate",
	"buc": "Buckler",
	"ne9": "Cantor Trophy",
	"cap": "Cap",
	"bab": "Carnage Helm",
	"xlm": "Casque",
	"mbt": "Chain Boots",
	"chn": "Chain Mail",
	"xul": "Chaos Armor",
	"ci0": "Circlet",
	"uhc": "Colossus Girdle",
	"bae": "Conquerer Crown",
	"urn": "Corona",
	"ci1": "Coronet",
	"crn": "Crown",
	"pa5": "Crown Shield",
	"utg": "Crusader Gauntlets",
	"xrs": "Cuirass",
	"xsk": "Death Mask",
	"xuc": "Defender",
	"ne5": "Demon Head",
	"usk": "DemonHead",
	"xla": "Demonhide Armor",
	"xlb": "Demonhide Boots",
	"xlg": "Demonhide Gloves",
	"zlb": "Demonhide Sash",
	"bad": "Destroyer Helm",
	"ci3": "Diadem",
	"ung": "Diamond Mail",
	"xit": "Dragon Shield",
	"drf": "Dream Spirit",
	"uui": "Dusk Shroud",
	"drd": "Earth Spirit",
	"xth": "Embossed Plate",
	"dr4": "Falcon Mask",
	"ba2": "Fanged Helm",
	"ne7": "Fetish Trophy",
	"fld": "Field Plate",
	"fhl": "Full Helm",
	"ful": "Full Plate Mail",
	"bac": "Fury Visor",
	"ne4": "Gargoyle Head",
	"hgl": "Gauntlets",
	"xui": "Ghost Armor",
	"uhl": "Giant Conch",
	"hbl": "Girdle",
	"lgl": "Gloves",
	"gth": "Gothic Plate",
	"gts": "Gothic Shield",
	"xrn": "Grand Crown",
	"urs": "Great Hauberk",
	"ghm": "Great Helm",
	"dr7": "Griffon Headress",
	"xh9": "Grim Helm",
	"xsh": "Grim Shield",
	"baf": "Guardian Crown",
	"pa9": "Gilded Shield",
	"hla": "Hard Leather Armor",
	"dr2": "Hawk Helm",
	"uuc": "Heater",
	"tbl": "Heavy Belt",
	"vbt": "Heavy Boots",
	"xmg": "Heavy Bracers",
	"vgl": "Heavy Gloves",
	"nea": "Heirophant Trophy",
	"ult": "Hellforged Plate",
	"neg": "Hellspawn Skull",
	"hlm": "Helm",
	"pa3": "Heraldic Shield",
	"ba3": "Horned Helm",
	"dr8": "Hunter's Guise",
	"ukp": "Hydraskull",
	"urg": "Hyperion",
	"ba1": "Jawbone Cap",
	"ba6": "Jawbone Visor",
	"kit": "Kite Shield",
	"uld": "Kraken Shell",
	"uth": "Lacquered Plate",
	"lrg": "Large Shield",
	"lea": "Leather Armor",
	"lbt": "Leather Boots",
	"vbl": "Light Belt",
	"tgl": "Light Gauntlets",
	"ltp": "Light Plate",
	"tbt": "Light Plated Boots",
	"xng": "Linked Mail",
	"ba7": "Lion Helm",
	"ucl": "Loricated Mail",
	"uml": "Luna",
	"xtp": "Mage Plate",
	"msk": "Mask",
	"xhn": "Mesh Armor",
	"zmb": "Mesh Belt",
	"xmb": "Mesh Boots",
	"neb": "Minion Skull",
	"utb": "Mirrored Boots",
	"umc": "Mithril Coil",
	"uit": "Monarch",
	"ne6": "Mummified Trophy",
	"uhb": "Myrmidon Greaves",
	"uhg": "Ogre Gauntlets",
	"xar": "Ornate Plate",
	"ned": "Overseer Skull",
	"xow": "Pasive",
	"hbt": "Plate Boots",
	"ne1": "Preserved Head",
	"pa8": "Protector Shield",
	"qui": "Quilted Armor",
	"ba8": "Rage Mask",
	"rng": "Ring Mail",
	"pa2": "Rondache",
	"xml": "Round Shield",
	"paa": "Royal Shield",
	"xpl": "Russet Armor",
	"uar": "Sacred Armor",
	"dr9": "Sacred Feathers",
	"pac": "Sacred Rondache",
	"pab": "Sacred Targe",
	"xkp": "Sallet",
	"lbl": "Sash",
	"ba9": "Savage Helmet",
	"scl": "Scale Mail",
	"ula": "Scarab Husk",
	"uvb": "Scarabshell Boots",
	"xrg": "Scutum",
	"xea": "Serpentskin Armor",
	"ne8": "Sexton Trophy",
	"uul": "Shadow Plate",
	"uap": "Shako",
	"zvb": "Sharkskin Belt",
	"xvb": "Sharkskin Boots",
	"xvg": "Sharkskin Gloves",
	"xld": "Sharktooth Armor",
	"skp": "Skull Cap",
	"dre": "Sky Spirit",
	"baa": "Slayer Guard",
	"sml": "Small Shield",
	"ulc": "Spiderweb Sash",
	"spk": "Spiked Shield",
	"uhm": "Spired Helm",
	"dr5": "Spirit Mask",
	"spl": "Splint Mail",
	"stu": "Studded Leather",
	"nee": "Succubus Skull",
	"drc": "Sun Spirit",
	"pa1": "Targe",
	"xlt": "Templar Coat",
	"ci2": "Tiara",
	"xcl": "Tigulated Mail",
	"dra": "Totemic Mask",
	"tow": "Tower Shield",
	"xtu": "Trellised Armor",
	"utc": "Troll Belt",
	"ush": "Troll Nest",
	"ne3": "Unraveller Head",
	"umg": "Vambraces",
	"uvg": "Vampirebone Gloves",
	"uvc": "Vampirefang Belt",
	"paf": "Vortex Shield",
	"zhb": "War Belt",
	"xhb": "War Boots",
	"xhg": "War Gauntlets",
	"xap": "War Hat",
	"uts": "Ward",
	"xhm": "Winged Helm",
	"utu": "Wire Fleece",
	"dr1": "Wolf Head",
	"uea": "Wyrmhide",
	"ulb": "Wyrmhide Boots",
	"pae": "Zakarum Shield",
	"ne2": "Zombie Head",
}

// Weapon Codes.
var weaponCodes = map[string]string{
	"9gi": "Ancient Axe",
	"9wd": "Ancient Sword",
	"8lx": "Arbalest",
	"6ws": "Archon Staff",
	"am6": "Ashwood Bow",
	"7sm": "Ataghan",
	"axe": "Axe",
	"bal": "Balanced Axe",
	"bkf": "Balanced Knife",
	"8hx": "Balista",
	"7gs": "Balrog Blade",
	"7s7": "Balrog Spear",
	"9sp": "Barbed Club",
	"bar": "Bardiche",
	"bsw": "Bastard Sword",
	"btx": "Battle Axe",
	"7cs": "Battle Cestus",
	"9tk": "Battle Dart",
	"9wh": "Battle Hammer",
	"9s8": "Battle Scythe",
	"bst": "Battle Staff",
	"9bs": "Battle Sword",
	"9ba": "Bearded Axe",
	"9h9": "Bec De Corbin",
	"7wa": "Berserker Axe",
	"9vo": "Bill",
	"bld": "Blade",
	"6hb": "Blade Bow",
	"btl": "Blade Talons",
	"7dg": "Bone Knife",
	"bwn": "Bone Wand",
	"brn": "Brandistock",
	"bax": "Broad Axe",
	"bsd": "Broad Sword",
	"9wn": "Burnt Wand",
	"7ws": "Caduceus",
	"8lb": "Cedar Bow",
	"8cs": "Cedar Staff",
	"am7": "Ceremonial Bow",
	"ama": "Ceremonial Javelin",
	"am9": "Ceremonial Pike",
	"am8": "Ceremonial Spear",
	"ces": "Cestus",
	"7ga": "Champion Axe",
	"7b7": "Champion Sword",
	"gpm": "Choking Gas Potion",
	"8rx": "Chu-Ko-Nu",
	"9kr": "Cinquedeas",
	"ob4": "Clasped Orb",
	"clw": "Claws",
	"clm": "Claymore",
	"9ax": "Cleaver",
	"ob8": "Cloudy Sphere",
	"clb": "Club",
	"7fb": "Colossal Sword",
	"7gd": "Colossal Blade",
	"6hx": "Colossus Crossbow",
	"7vo": "Colossus Voulge",
	"cbw": "Composite Bow",
	"7bs": "Conquest Sword",
	"mxb": "Crossbow",
	"9mp": "Crowbill",
	"6l7": "Crusader Bow",
	"7pa": "Cryptic Axe",
	"7ls": "Cryptic Sword",
	"crs": "Crystal Sword",
	"ob7": "Crystalline Globe",
	"9cl": "Cudgel",
	"9sm": "Cutlass",
	"9cm": "Dacian Falx",
	"dgr": "Dagger",
	"7bt": "Decapitator",
	"d33": "Decoy Dagger",
	"6rx": "Demon Crossbow",
	"obd": "Demon Heart",
	"7mt": "Devil Star",
	"6s7": "Diamond Bow",
	"9cr": "Dimensional Blade",
	"obf": "Dimensional Shard",
	"dir": "Dirk",
	"9ws": "Divine Scepter",
	"2ax": "Double Axe",
	"8cb": "Double Bow",
	"ob5": "Dragon Stone",
	"ob1": "Eable Orb",
	"8sb": "Edge Bow",
	"6cs": "Elder Staff",
	"obc": "Eldritch Orb",
	"7sb": "Elegant Blade",
	"92h": "Espadon",
	"72a": "Ettin Axe",
	"9gd": "Executioner Sword",
	"opm": "Exploding Potion",
	"7ss": "Falcata",
	"flc": "Falchion",
	"7kr": "Fanged Knife",
	"9xf": "Fascia",
	"7la": "Feral Axe",
	"7lw": "Feral Claws",
	"fla": "Flail",
	"flb": "Flamberge",
	"9ma": "Flanged Mace",
	"7ta": "Flying Axe",
	"7tk": "Flying Knife",
	"9ta": "Francisca",
	"opl": "Fulmating Potion",
	"9tr": "Fuscina",
	"7gl": "Ghost Glaive",
	"7st": "Ghost Spear",
	"7yw": "Ghost Wand",
	"gix": "Giant Axe",
	"gis": "Giant Sword",
	"7wc": "Giant Thresher",
	"g33": "Gidbinn",
	"9ss": "Gladius",
	"glv": "Glaive",
	"7gi": "Glorious Axe",
	"ob6": "Glowing Orb",
	"cst": "Gnarled Staff",
	"6mx": "Gorgon Crossbow",
	"9ga": "Gothic Axe",
	"8lw": "Gothic Bow",
	"8bs": "Gothic Staff",
	"9b9": "Gothic Sword",
	"amc": "Grand Matron Bow",
	"gsc": "Grand Scepter",
	"9gw": "Grave Wand",
	"gax": "Great Axe",
	"6cb": "Great Bow",
	"gma": "Great Maul",
	"9pi": "Great Pilum",
	"7h7": "Great Poleaxe",
	"gsd": "Great Sword",
	"9lw": "Greater Claws",
	"9tw": "Greater Talons",
	"9wc": "Grim Scythe",
	"gwn": "Grim Wand",
	"hal": "Halberd",
	"hax": "Hand Axe",
	"9cs": "Hand Scythe",
	"9ts": "Harpoon",
	"9ha": "Hatchet",
	"axf": "HatchetHands",
	"obb": "Heavenly Stone",
	"hxb": "Heavy Crossbow",
	"hfh": "Hellforge Hammer",
	"7cm": "Highland Blade",
	"9qs": "Holy Water Sprinkler",
	"hdm": "Horadric Malus",
	"hst": "Horadric Staff",
	"hbw": "Hunters Bow",
	"9b8": "Hurlbat",
	"6lw": "Hydra Bow",
	"7fc": "Hydra Edge",
	"7ja": "Hyperion Javelin",
	"7sr": "Hyperion Spear",
	"9mt": "Jagged Star",
	"jav": "Javelin",
	"8ss": "Jo Staff",
	"ktr": "Katar",
	"qf1": "Khalim Flail",
	"9fl": "Knout",
	"kri": "Kriss",
	"9p9": "Lance",
	"lax": "Large Axe",
	"7bl": "Legend Spike",
	"72h": "Legend Sword",
	"7wh": "Legendary Mallet",
	"7bw": "Lich Wand",
	"lxb": "Light Crossbow",
	"9b7": "Lochaber Axe",
	"lbb": "Long Battle Bow",
	"lbw": "Long Bow",
	"8l8": "Long Siege Bow",
	"lst": "Long Staff",
	"lsd": "Long Sword",
	"lwb": "Long War Bow",
	"mac": "Mace",
	"am5": "Maiden Javelin",
	"am4": "Maiden Pike",
	"am3": "Maiden Spear",
	"7br": "Mancatcher",
	"9gm": "Martel de Fer",
	"amb": "Matriarchal Bow",
	"ame": "Matriarchal Pike",
	"amd": "Matriarchal Spear",
	"amf": "Matriarchal Javelin",
	"mau": "Maul",
	"7sc": "Mighty Scepter",
	"9la": "Military Axe",
	"mpi": "Military Pick",
	"7di": "Mithral Point",
	"mst": "Morning Star",
	"7wd": "Mythical Sword",
	"9wa": "Naga",
	"7o7": "Ogre Axe",
	"7m7": "Ogre Maul",
	"ops": "Oil Potion",
	"9pa": "Partizan",
	"6lx": "Pellet Bow",
	"9yw": "Petrified Wand",
	"7cr": "Phase Blade",
	"pik": "Pike",
	"pil": "Pilum",
	"9dg": "Poignard",
	"pax": "Poleaxe",
	"7wn": "Polished Wand",
	"8ls": "Quarter Staff",
	"9ar": "Quhab",
	"gps": "Rancid Gas Potion",
	"8hb": "Razor Bow",
	"am2": "Reflex Bow",
	"7ma": "Reinforced Mace",
	"rxb": "Repeating Crossbow",
	"9di": "Rondel",
	"8sw": "Rune Bow",
	"9sc": "Rune Scepter",
	"8ws": "Rune Staff",
	"9ls": "Rune Sword",
	"7tw": "Runic Talons",
	"sbr": "Sabre",
	"ob2": "Sacred Globe",
	"scp": "Scepter",
	"scm": "Scimitar",
	"skr": "Scissors Katar",
	"9qr": "Scissors Quhab",
	"7qr": "Scissors Suwayyah",
	"7fl": "Scourge",
	"scy": "Scythe",
	"7qs": "Seraph Rod",
	"6lb": "Shadow Bow",
	"9sb": "Shamshir",
	"6bs": "Shillelah",
	"sbb": "Short Battle Bow",
	"sbw": "Short Bow",
	"8s8": "Short Siege Bow",
	"ssp": "Short Spear",
	"sst": "Short Staff",
	"ssd": "Short Sword",
	"swb": "Short War Bow",
	"8mx": "Siege Crossbow",
	"7ba": "Silver Edged Axe",
	"9s9": "Simbilan",
	"7ax": "Small Crescent",
	"ob3": "Smoked Sphere",
	"ob9": "Sparkling Ball",
	"spr": "Spear",
	"spt": "Spetum",
	"9gl": "Spiculum",
	"6sb": "Spider Bow",
	"spc": "Spiked Club",
	"msf": "Staff Of The Kings",
	"am1": "Stag Bow",
	"6ls": "Stalagmite",
	"9bl": "Stilleto",
	"gpl": "Strangling Gas Potion",
	"7tr": "Stygian Pike",
	"7pi": "Stygian Pilum",
	"qf2": "Super Khalim Flail",
	"7ar": "Suwayyah",
	"oba": "Swirling Crystal",
	"9bt": "Tabar",
	"7s8": "Thresher",
	"tax": "Throwing Axe",
	"tkf": "Throwing Knife",
	"tsp": "Throwing Spear",
	"7gm": "Thunder Maul",
	"7ha": "Tomahawk",
	"9bw": "Tomb Wand",
	"tri": "Trident",
	"7cl": "Truncheon",
	"9fc": "Tulwar",
	"9gs": "Tusk Sword",
	"92a": "Twin Axe",
	"2hs": "Two Handed Axe",
	"7sp": "Tyrant Club",
	"7gw": "Unearthed Wand",
	"obe": "Vortex Orb",
	"vou": "Voulge",
	"6ss": "Walking Stick",
	"wnd": "Wand",
	"wax": "War Axe",
	"9m9": "War Club",
	"9bk": "War Dart",
	"7xf": "War Fist",
	"9br": "War Fork",
	"whm": "War Hammer",
	"9ja": "War Javelin",
	"7p7": "War Pike",
	"wsp": "War Scepter",
	"wsc": "War Scythe",
	"9sr": "War Spear",
	"7mp": "War Spike",
	"wst": "War Staff",
	"wsd": "War Sword",
	"6sw": "Ward Bow",
	"7b8": "Winged Axe",
	"7ts": "Winged Harpoon",
	"7bk": "Winged Knife",
	"leg": "Wirts Leg",
	"wrb": "Wrist Blade",
	"9wb": "Wrist Spike",
	"7wb": "Wrist Sword",
	"9st": "Yari",
	"ywn": "Yew Wand",
	"9fb": "Zweihander",
}

// Misc codes, like jewelry, gems, potions and runes.
var miscCodes = map[string]string{
	"gsv": "Amethyst",
	"r11": "Amn Rune",
	"amu": "Amulet",
	"yps": "Antidote Potion",
	"aqv": "Arrows",
	"bey": "Baal's Eye",
	"bks": "Bark Scroll",
	"r30": "Ber Rune",
	"cqv": "Bolts",
	"ass": "Book of Skill",
	"bet": "Burning Essence of Terror",
	"r32": "Cham Rune",
	"ceh": "Charged Essence of Hatred",
	"cm3": "Grand Charm",
	"cm2": "Large Charm",
	"cm1": "Small Charm",
	"gcv": "Chipped Amethyst",
	"gcw": "Chipped Diamond",
	"gcg": "Chipped Emerald",
	"gcr": "Chipped Ruby",
	"gcb": "Chipped Sapphire",
	"skc": "Chipped Skull",
	"gcy": "Chipped Topaz",
	"bkd": "Deciphered Bark Scroll",
	"dhn": "Diablo's Horn",
	"gsw": "Diamond",
	"r14": "Dol Rune",
	"r01": "El Rune",
	"r02": "Eld Rune",
	"elx": "Elixir",
	"gsg": "Emerald",
	"r05": "Eth Rune",
	"r19": "Fal Rune",
	"fed": "Festering Essence of Destruction",
	"gfv": "Flawed Amethyst",
	"gfw": "Flawed Diamond",
	"gfg": "Flawed Emerald",
	"gfr": "Flawed Ruby",
	"gfb": "Flawed Sapphire",
	"skf": "Flawed Skull",
	"gfy": "Flawed Topaz",
	"gzv": "Flawless Amethyst",
	"glw": "Flawless Diamond",
	"glg": "Flawless Emerald",
	"glr": "Flawless Ruby",
	"glb": "Flawless Sapphire",
	"skl": "Flawless Skull",
	"gly": "Flawless Topaz",
	"hpf": "Full Healing Potion",
	"mpf": "Full Mana Potion",
	"rvl": "Full Rejuvenation Potion",
	"gld": "Gold",
	"g34": "Gold Bird",
	"hp5": "Greater Healing Potion",
	"mp5": "Greater Mana Potion",
	"r25": "Gul Rune",
	"hp3": "Healing Potion",
	"hpo": "Healing Potion",
	"r15": "Hel Rune",
	"hrb": "Herb",
	"box": "Horadric Cube",
	"ibk": "Identify Book",
	"isc": "Identify Scroll",
	"r16": "Io Rune",
	"r24": "Ist Rune",
	"r06": "Ith Rune",
	"j34": "Jade Figurine",
	"r31": "Jah Rune",
	"jew": "Jewel",
	"pk3": "Key of Destruction",
	"pk2": "Key of Hate",
	"pk1": "Key of Terror",
	"qbr": "Khalim's Brain",
	"qey": "Khalim's Eye",
	"qhr": "Khalim's Heart",
	"r18": "Ko Rune",
	"bbb": "Lam Esens Tome",
	"bpl": "Large Blue Potion",
	"rpl": "Large Red Potion",
	"r20": "Lem Rune",
	"hp1": "Lesser Healing Potion",
	"mp1": "Lesser Mana Potion",
	"hp2": "Light Healing Potion",
	"mp2": "Light Mana Potion",
	"r28": "Lo Rune",
	"r17": "Lum Rune",
	"ice": "Maguffin",
	"r23": "Mal Rune",
	"mp3": "Mana Potion",
	"mpo": "Mana Potion",
	"mbr": "Mephisto's Brain",
	"luv": "Mephisto Key",
	"mss": "Mephisto Soul Stone",
	"r04": "Nef Rune",
	"r27": "Ohm Rune",
	"r09": "Ort Rune",
	"gpv": "Perfect Amethyst",
	"gpw": "Perfect Diamond",
	"gpg": "Perfect Emerald",
	"gpr": "Perfect Ruby",
	"gpb": "Perfect Sapphire",
	"skz": "Perfect Skull",
	"gpy": "Perfect Topaz",
	"ear": "Player Ear",
	"r21": "Pul Rune",
	"r08": "Ral Rune",
	"rvs": "Rejuvenation Potion",
	"rin": "Ring",
	"gsr": "Ruby",
	"gsb": "Sapphire",
	"0sc": "Scroll",
	"tr1": "Scroll of Horadric",
	"tr2": "Scroll of Malah",
	"xyz": "Potion of Life",
	"r13": "Shael Rune",
	"key": "Skeleton Key",
	"sku": "Skull",
	"bps": "Small Blue Potion",
	"rps": "Small Red Potion",
	"r12": "Sol Rune",
	"vps": "Stamina Potion",
	"std": "Standard of Heroes",
	"hp4": "Strong Healing Potion",
	"mp4": "Strong Mana Potion",
	"r29": "Sur Rune",
	"r07": "Tal Rune",
	"wms": "Thawing Potion",
	"r10": "Thul Rune",
	"r03": "Tir Rune",
	"toa": "Token of Absolution",
	"gsy": "Topaz",
	"tch": "Torch",
	"tbk": "Town Portal Book",
	"tsc": "Town Portal Scroll",
	"tes": "Twisted Essence of Suffering",
	"r22": "Um Rune",
	"r26": "Vex Rune",
	"vip": "Viper Amulet",
	"r33": "Zod Rune",
}

type magicAttribute struct {
	ID     uint64
	Name   string
	Values []uint64
}

type magicalProperty struct {
	Bits []uint
	Bias uint64
	Name string
}

var magicalProperties = map[uint64]magicalProperty{
	0:  {Bits: []uint{7}, Bias: 32, Name: "+X to Strength"},
	1:  {Bits: []uint{7}, Bias: 32, Name: "+X to Energy"},
	2:  {Bits: []uint{7}, Bias: 32, Name: "+X to Dexterity"},
	3:  {Bits: []uint{7}, Bias: 32, Name: "+X to Vitality"},
	7:  {Bits: []uint{9}, Bias: 32, Name: "+X to Life"},
	9:  {Bits: []uint{8}, Bias: 32, Name: "+X to Mana"},
	11: {Bits: []uint{8}, Bias: 32, Name: "+X to Maximum Stamina"},
	16: {Bits: []uint{9}, Name: "+X% Enhanced Defense"},
	17: {Bits: []uint{9, 9}, Name: "+X% Enhanced Damage"},
	19: {Bits: []uint{10}, Name: "+X to Attack rating"},
	20: {Bits: []uint{6}, Name: "+X% Increased chance of blocking"},
	21: {Bits: []uint{6}, Name: "+X to Minimum damage"},
	22: {Bits: []uint{7}, Name: "+X to Maximum damage"},
	// Possible duplicate of id: 21, usually seen together
	23: {Bits: []uint{6}, Name: "+X to Minimum damage"},
	// Possible duplicate of id: 22, usually seen together
	24: {Bits: []uint{7}, Name: "+X to Maximum damage"},
	27: {Bits: []uint{8}, Name: "Regenerate Mana X%"},
	28: {Bits: []uint{8}, Name: "Heal Stamina X%"},
	31: {Bits: []uint{10}, Bias: 10, Name: "+X Defense"},
	32: {Bits: []uint{10}, Bias: 10, Name: "+X vs. Missile"},
	33: {Bits: []uint{10}, Bias: 10, Name: "+X vs. Melee"},
	34: {Bits: []uint{6}, Name: "Damage Reduced by X"},
	35: {Bits: []uint{6}, Name: "Magic Damage Reduced by X"},
	36: {Bits: []uint{8}, Name: "Damage Reduced by X%"},
	37: {Bits: []uint{8}, Name: "Magic Resist +X%"},
	38: {Bits: []uint{8}, Name: "+X% to Maximum Magic Resist"},
	39: {Bits: []uint{8}, Name: "Fire Resist +X%"},
	40: {Bits: []uint{5}, Name: "+X% to Maximum Fire Resist"},
	41: {Bits: []uint{8}, Name: "Lightning Resist +X%"},
	42: {Bits: []uint{5}, Name: "+X% to Maximum Lightning Resist"},
	43: {Bits: []uint{8}, Name: "Cold Resist +X%"},
	44: {Bits: []uint{5}, Name: "+X% to Maximum Cold Resist"},
	45: {Bits: []uint{8}, Name: "Poison Resist +X%"},
	46: {Bits: []uint{5}, Name: "+X% to Maximum Poison Resist"},
	48: {Bits: []uint{8, 8}, Name: "Adds X-Y Fire Damage"},
	50: {Bits: []uint{6, 9}, Name: "Adds X-Y Lightning Damage"},
	52: {Bits: []uint{6, 7}, Name: "Adds X-Y Magic Damage"},
	54: {Bits: []uint{6, 8, 8}, Name: "Adds X-Y Cold Damage"},
	57: {Bits: []uint{10, 10, 9}, Name: "Adds X-Y Poison Damage over Z Seconds"},
	60: {Bits: []uint{7}, Name: "X% Life stolen per hit"},
	73: {Bits: []uint{8}, Bias: 30, Name: "+X Maximum Durability"},
	74: {Bits: []uint{6}, Bias: 30, Name: "Replenish Life +X"},
	75: {Bits: []uint{7}, Bias: 20, Name: "Increase Maximum Durability X%"},
	76: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Life X%"},
	77: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Mana X%"},
	78: {Bits: []uint{7}, Name: "Attacker Takes Damage of X"},
	79: {Bits: []uint{9}, Bias: 100, Name: "X% Extra Gold from Monsters"},
	80: {Bits: []uint{8}, Bias: 100, Name: "X% Better Chance of Getting Magic Items"},
	81: {Bits: []uint{7}, Name: "Knockback"},
	83: {Bits: []uint{3}, Name: "+X to Amazon Skill Levels"},
	84: {Bits: []uint{3}, Name: "+X to Paladin Skill Levels"},
	85: {Bits: []uint{3}, Name: "+X to Necromancer Skill Levels"},
	86: {Bits: []uint{3}, Name: "+X to Sorceress Skill Levels"},
	87: {Bits: []uint{3}, Name: "+X to Barbarian Skill Levels"},
	89: {Bits: []uint{4}, Bias: 4, Name: "+X to Light Radius"},
	// This property is not displayed on the item, but its effect is to alter
	// the color of the ambient light.
	90: {Bits: []uint{5}, Name: "Ambient light"},
	// After subtracting the bias, this is usually a negative number.
	91:  {Bits: []uint{8}, Bias: 100, Name: "Requirements -X%"},
	93:  {Bits: []uint{7}, Bias: 20, Name: "X% Increased Attack Speed"},
	96:  {Bits: []uint{7}, Bias: 20, Name: "X% Faster Run/Walk"},
	99:  {Bits: []uint{7}, Bias: 20, Name: "X% Faster Hit Recovery"},
	102: {Bits: []uint{7}, Bias: 20, Name: "X% Faster Block Rate"},
	105: {Bits: []uint{7}, Bias: 20, Name: "X% Faster Cast Rate"},

	// These properties usually applied to class specific items,
	// first value selects the skill, the second determines how many
	// additional skill points are given.
	107: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	108: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	109: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	181: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	182: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	183: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	184: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	185: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	186: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	187: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},

	110: {Bits: []uint{8}, Bias: 20, Name: "Poison Length Reduced by X%"},
	111: {Bits: []uint{7}, Bias: 20, Name: "Damage +X"},
	112: {Bits: []uint{7}, Name: "Hit Causes Monsters to Flee X%"},
	113: {Bits: []uint{7}, Name: "Hit Blinds Target +X"},
	114: {Bits: []uint{6}, Name: "X% Damage Taken Goes to Mana"},
	115: {Bits: []uint{1}, Name: "Ignore Target Defense"},
	// The value of the data field is always 1.
	116: {Bits: []uint{7}, Name: "X% Target Defense"},
	// The value of the data field is always 1.
	117: {Bits: []uint{7}, Name: "Prevent Monster Heal"},
	// The value of the data field is always 1.
	118: {Bits: []uint{1}, Name: "Half Freeze Duration"},
	119: {Bits: []uint{9}, Bias: 20, Name: "X% Bonus to Attack Rating"},
	120: {Bits: []uint{7}, Bias: 128, Name: "X to Monster Defense Per Hit"},
	121: {Bits: []uint{9}, Bias: 20, Name: "+X% Damage to Demons"},
	122: {Bits: []uint{9}, Bias: 20, Name: "+X% Damage to Undead"},
	123: {Bits: []uint{10}, Bias: 128, Name: "+X to Attack Rating against Demons"},
	124: {Bits: []uint{10}, Bias: 128, Name: "+X to Attack Rating against Undead"},
	126: {Bits: []uint{4}, Name: "+X to Fire Skills"},
	127: {Bits: []uint{3}, Name: "+X to All Skill Levels"},
	128: {Bits: []uint{5}, Name: "Attacker Takes Lightning Damage of X"},
	134: {Bits: []uint{5}, Name: "Freezes Target"},
	135: {Bits: []uint{7}, Name: "X% Chance of Open Wounds"},
	136: {Bits: []uint{7}, Name: "X% Change of Crushing Blow"},
	137: {Bits: []uint{7}, Name: "+X Kick Damage"},
	138: {Bits: []uint{7}, Name: "+X to Mana After Each Kill"},
	139: {Bits: []uint{7}, Name: "+X Life after each Demon Kill"},
	// Unknown property, shows up on Swordback Hold Spiked Shield.
	140: {Bits: []uint{7}, Name: "Unknown"},
	141: {Bits: []uint{7}, Name: "X% Deadly Strike"},
	142: {Bits: []uint{7}, Name: "Fire Absorb X%"},
	143: {Bits: []uint{7}, Name: "X Fire Absorb"},
	144: {Bits: []uint{7}, Name: "Lightning Absorb X%"},
	145: {Bits: []uint{7}, Name: "X Lightning Absorb"},
	146: {Bits: []uint{7}, Name: "Magic Absorb X%"},
	147: {Bits: []uint{7}, Name: "X Magic Absorb"},
	148: {Bits: []uint{7}, Name: "Cold Absorb X%"},
	149: {Bits: []uint{7}, Name: "X Cold Absorb"},
	150: {Bits: []uint{7}, Name: "Slows Target by X%"},
	151: {Bits: []uint{7}, Name: "Blessed Aim"},
	152: {Bits: []uint{7}, Name: "Defiance"},
	153: {Bits: []uint{1}, Name: "Cannot Be Frozen"},
	154: {Bits: []uint{7}, Name: "X% Slower Stamina Drain"},
	155: {Bits: []uint{7}, Name: "X% Chance to Reanimate Target"},
	156: {Bits: []uint{7}, Name: "Piercing Attack"},
	157: {Bits: []uint{7}, Name: "Fires Magic Arrows"},
	158: {Bits: []uint{7}, Name: "Fires Explosive Arrows or Bolts"},
	// Never seen this by itself, always accompanied by properties 21 and 23
	159: {Bits: []uint{6}, Name: "+X to Minimum Damage"},
	// Never seen this by itself, always accompanied by properties 22 and 24
	160: {Bits: []uint{7}, Name: "+X to Maximum Damage"},
	179: {Bits: []uint{3}, Name: "+X to Druid Skill Levels"},
	180: {Bits: []uint{3}, Name: "+X to Assassin Skill Levels"},

	// A skill set is a class specific skill tree id, e.g bow and crossbow skills,
	// traps or war cries. ID's are described below.
	188: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	189: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	190: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	191: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	192: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	193: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},

	194: {Bits: []uint{4}, Name: "Adds X extra sockets to the item"},

	// First order is spell id, % chance and level.
	195: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on attack"},
	196: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on attack"},
	197: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on attack"},

	// First order is spell id, % chance and level.
	198: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on striking"},
	199: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on striking"},
	200: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on striking"},

	// First order is spell id, % chance and level.
	201: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id when struck"},
	202: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id when struck"},
	203: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id when struck"},

	// First value selects the spell id, second value is level, third is remaining charges
	// and the last is the total number of charges.
	204: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	205: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	206: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	207: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	208: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	209: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	210: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	211: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	212: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	213: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},

	// All values based on character level are stored in eights, so take
	// the number divide by 8 and multiply by the character level and round down.
	214: {Bits: []uint{6}, Name: "+X to Defense (Based on Character Level)"},
	215: {Bits: []uint{6}, Name: "X% Enhanced Defense (Based on Character Level)"},
	216: {Bits: []uint{6}, Name: "+X to Life (Based on Character Level)"},
	217: {Bits: []uint{6}, Name: "+X to Mana (Based on Character Level)"},
	218: {Bits: []uint{6}, Name: "+X to Maximum Damage (Based on Character Level)"},
	219: {Bits: []uint{6}, Name: "X% Enhanced Maximum Damage (Based on Character Level)"},
	220: {Bits: []uint{6}, Name: "+X to Strength (Based on Character Level)"},
	221: {Bits: []uint{6}, Name: "+X to Dexterity (Based on Character Level)"},
	222: {Bits: []uint{6}, Name: "+X to Energy (Based on Character Level)"},
	223: {Bits: []uint{6}, Name: "+X to Vitality (Based on Character Level)"},
	224: {Bits: []uint{6}, Name: "+X to Attack Rating (Based on Character Level)"},
	225: {Bits: []uint{6}, Name: "X% Bonus to Attack Rating (Based on Character Level)"},
	226: {Bits: []uint{6}, Name: "+X Cold Damage (Based on Character Level)"},
	227: {Bits: []uint{6}, Name: "+X Fire Damage (Based on Character Level)"},
	228: {Bits: []uint{6}, Name: "+X Lightning Damage (Based on Character Level)"},
	229: {Bits: []uint{6}, Name: "+X Poison Damage (Based on Character Level)"},
	230: {Bits: []uint{6}, Name: "Cold Resist +X% (Based on Character Level)"},
	231: {Bits: []uint{6}, Name: "Fire Resist +X% (Based on Character Level)"},
	232: {Bits: []uint{6}, Name: "Lightning Resist +X% (Based on Character Level)"},
	233: {Bits: []uint{6}, Name: "Poison Resist +X% (Based on Character Level)"},
	234: {Bits: []uint{6}, Name: "+X Cold Absorb (Based on Character Level)"},
	235: {Bits: []uint{6}, Name: "+X Fire Absorb (Based on Character Level)"},
	236: {Bits: []uint{6}, Name: "+X Lightning Absorb (Based on Character Level)"},
	237: {Bits: []uint{6}, Name: "X Poison Absorb (Based on Character Level)"},
	238: {Bits: []uint{6}, Name: "Attacker Takes Damage of X (Based on Character Level)"},
	239: {Bits: []uint{6}, Name: "X% Extra Gold from Monsters (Based on Character Level)"},
	240: {Bits: []uint{6}, Name: "X% Better Chance of Getting Magic Items (Based on Character Level)"},
	241: {Bits: []uint{6}, Name: "Heal Stamina Plus X% (Based on Character Level)"},
	243: {Bits: []uint{6}, Name: "X% Damage to Demons (Based on Character Level)"},
	244: {Bits: []uint{6}, Name: "X% Damage to Undead (Based on Character Level)"},
	245: {Bits: []uint{6}, Name: "+X to Attack Rating against Demons (Based on Character Level)"},
	246: {Bits: []uint{6}, Name: "+X to Attack Rating against Undead (Based on Character Level)"},
	247: {Bits: []uint{6}, Name: "X% Chance of Crushing Blow (Based on Character Level)"},
	248: {Bits: []uint{6}, Name: "X% Chance of Open Wounds (Based on Character Level)"},
	249: {Bits: []uint{6}, Name: "+X Kick Damage (Based on Character Level)"},
	250: {Bits: []uint{6}, Name: "X% to Deadly Strike (Based on Character Level)"},

	// The value of the data field is not actually a time period, but a frequency in terms
	// of the number of times durability is repaired over a period of 100 seconds.
	// For example, if the value is 5, then this property repairs 1 durability in 100 / 5 = 20 seconds.
	252: {Bits: []uint{5}, Name: "Repairs 1 Durability in X Seconds"},

	// As in the previous property, the value of the data field is a frequency in terms of the number
	// replenished over a period of 100 seconds. For example if the value is 4, then this property
	// replenishes 1 item in 100 / 4 = 25 seconds.
	253: {Bits: []uint{5}, Name: "Replenishes Quantity"},

	// Number of additional items beyond the base limit, for example if the base
	// is 50 and additional is 30, then the total is 50 + 30.
	254: {Bits: []uint{8}, Name: "Increased Stack Size"},

	// These are some weird values, no idea if they are used in game or not.
	// But these values change depending on the time of day in the game.
	// The format of the bit fields are the same in all cases, the first 2 bits
	// specifies the the of time when the value is at its maximum.
	//
	// The second and third are respectively the minimum and maximum values of the property.
	// The maximum value at the time specified and the minimum at the opposite.

	// TODO: Add ids 268 - 303 if they prove to exist.
}

// All item types that contain the quantity bits will exist in here,
// we'll use this when reading items to make sure we only read quantity bits
// when they exist, or we'll ruin the rest of the bit offsets for the item.
var quantityMap = map[string]bool{}
