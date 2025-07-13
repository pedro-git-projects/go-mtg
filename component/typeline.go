package component

import (
	"fmt"
	"slices"
	"strings"
)

type Supertype uint8

const (
	Legendary Supertype = iota
	Basic
	Snow
	World
	Ongoing
	supertypeCount
)

var supertypeNames = [...]string{
	Legendary: "Legendary",
	Basic:     "Basic",
	Snow:      "Snow",
	World:     "World",
	Ongoing:   "Ongoing",
}

var supertypeLookup = map[string]Supertype{
	"Legendary": Legendary,
	"Basic":     Basic,
	"Snow":      Snow,
	"World":     World,
	"Ongoing":   Ongoing,
}

func (s Supertype) String() string {
	if int(s) < len(supertypeNames) {
		return supertypeNames[s]
	}
	return fmt.Sprintf("Supertype(%d)", s)
}

func (s Supertype) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *Supertype) UnmarshalText(text []byte) error {
	str := string(text)
	if v, ok := supertypeLookup[str]; ok {
		*s = v
		return nil
	}
	for k, v := range supertypeLookup {
		if strings.EqualFold(k, str) {
			*s = v
			return nil
		}
	}
	return fmt.Errorf("unknown Supertype %q", str)
}

type TypeId uint8

const (
	Artifact TypeId = iota
	Battle
	Conspiracy
	Creature
	Dungeon
	Enchantment
	Instant
	Kindred
	Land
	Phenomenon
	Plane
	Planeswalker
	Scheme
	Sorcery
	Vanguard
	typeCount
)

var typeNames = [...]string{
	Artifact:     "Artifact",
	Battle:       "Battle",
	Conspiracy:   "Conspiracy",
	Creature:     "Creature",
	Dungeon:      "Dungeon",
	Enchantment:  "Enchantment",
	Instant:      "Instant",
	Kindred:      "Kindred",
	Land:         "Land",
	Phenomenon:   "Phenomenon",
	Plane:        "Plane",
	Planeswalker: "Planeswalker",
	Scheme:       "Scheme",
	Sorcery:      "Sorcery",
	Vanguard:     "Vanguard",
}

var typeLookup = map[string]TypeId{
	"Artifact":     Artifact,
	"Battle":       Battle,
	"Conspiracy":   Conspiracy,
	"Creature":     Creature,
	"Dungeon":      Dungeon,
	"Enchantment":  Enchantment,
	"Instant":      Instant,
	"Kindred":      Kindred,
	"Land":         Land,
	"Phenomenon":   Phenomenon,
	"Plane":        Plane,
	"Planeswalker": Planeswalker,
	"Scheme":       Scheme,
	"Sorcery":      Sorcery,
	"Vanguard":     Vanguard,
}

func (t TypeId) String() string {
	if int(t) < len(typeNames) {
		return typeNames[t]
	}
	return fmt.Sprintf("TypeId(%d)", t)
}

func (t TypeId) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *TypeId) UnmarshalText(text []byte) error {
	str := string(text)
	if v, ok := typeLookup[str]; ok {
		*t = v
		return nil
	}
	for k, v := range typeLookup {
		if strings.EqualFold(k, str) {
			*t = v
			return nil
		}
	}
	return fmt.Errorf("unknown TypeId %q", str)
}

type SubtypeId uint16

const (
	Sub_Ajani SubtypeId = iota
	Sub_Aminatou
	Sub_Angrath
	Sub_Arlinn
	Sub_Ashiok
	Sub_Bahamut
	Sub_Basri
	Sub_Bolas
	Sub_Calix
	Sub_Chandra
	Sub_Comet
	Sub_Dack
	Sub_Dakkon
	Sub_Daretti
	Sub_Davriel
	Sub_Dihada
	Sub_Domri
	Sub_Dovin
	Sub_Ellywick
	Sub_Elminster
	Sub_Elspeth
	Sub_Estrid
	Sub_Freyalise
	Sub_Garruk
	Sub_Gideon
	Sub_Grist
	Sub_Guff
	Sub_Huatli
	Sub_Jace
	Sub_Jared
	Sub_Jaya
	Sub_Jeska
	Sub_Kaito
	Sub_Karn
	Sub_Kasmina
	Sub_Kaya
	Sub_Kiora
	Sub_Koth
	Sub_Liliana
	Sub_Lolth
	Sub_Lukka
	Sub_Minsc
	Sub_Mordenkainen
	Sub_Nahiri
	Sub_Narset
	Sub_Niko
	Sub_Nissa
	Sub_Nixilis
	Sub_Oko
	Sub_Quintorius
	Sub_Ral
	Sub_Rowan
	Sub_Saheeli
	Sub_Samut
	Sub_Sarkhan
	Sub_Serra
	Sub_Sivitri
	Sub_Sorin
	Sub_Szat
	Sub_Tamiyo
	Sub_Tasha
	Sub_Teferi
	Sub_Teyo
	Sub_Tezzeret
	Sub_Tibalt
	Sub_Tyvar
	Sub_Ugin
	Sub_Urza
	Sub_Venser
	Sub_Vivien
	Sub_Vraska
	Sub_Vronos
	Sub_Will
	Sub_Windgrace
	Sub_Wrenn
	Sub_Xenagos
	Sub_Yanggu
	Sub_Yanling
	Sub_Zariel

	// Special spell subtypes
	Sub_Adventure
	Sub_Arcane
	Sub_Lesson
	Sub_Omen
	Sub_Trap

	// Creature subtype
	Sub_TimeLord

	// Creature and other subtypes (A–Z)
	Sub_Advisor
	Sub_Aetherborn
	Sub_Alien
	Sub_Ally
	Sub_Angel
	Sub_Antelope
	Sub_Ape
	Sub_Archer
	Sub_Archon
	Sub_Armadillo
	Sub_Army
	Sub_Artificer
	Sub_Assassin
	Sub_AssemblyWorker
	Sub_Astartes
	Sub_Atog
	Sub_Aurochs
	Sub_Avatar
	Sub_Azra
	Sub_Badger
	Sub_Balloon
	Sub_Barbarian
	Sub_Bard
	Sub_Basilisk
	Sub_Bat
	Sub_Bear
	Sub_Beast
	Sub_Beaver
	Sub_Beeble
	Sub_Beholder
	Sub_Berserker
	Sub_Bird
	Sub_Blinkmoth
	Sub_Boar
	Sub_Bringer
	Sub_Brushwagg
	Sub_Camarid
	Sub_Camel
	Sub_Capybara
	Sub_Caribou
	Sub_Carrier
	Sub_Cat
	Sub_Centaur
	Sub_Child
	Sub_Chimera
	Sub_Citizen
	Sub_Cleric
	Sub_Clown
	Sub_Cockatrice
	Sub_Construct
	Sub_Coward
	Sub_Coyote
	Sub_Crab
	Sub_Crocodile
	Sub_Ctan
	Sub_Custodes
	Sub_Cyberman
	Sub_Cyclops
	Sub_Dalek
	Sub_Dauthi
	Sub_Demigod
	Sub_Demon
	Sub_Deserter
	Sub_Detective
	Sub_Devil
	Sub_Dinosaur
	Sub_Djinn
	Sub_Doctor
	Sub_Dog
	Sub_Dragon
	Sub_Drake
	Sub_Dreadnought
	Sub_Drone
	Sub_Druid
	Sub_Dryad
	Sub_Dwarf
	Sub_Efreet
	Sub_Egg
	Sub_Elder
	Sub_Eldrazi
	Sub_Elemental
	Sub_Elephant
	Sub_Elf
	Sub_Elk
	Sub_Employee
	Sub_Eye
	Sub_Faerie
	Sub_Ferret
	Sub_Fish
	Sub_Flagbearer
	Sub_Fox
	Sub_Fractal
	Sub_Frog
	Sub_Fungus
	Sub_Gamer
	Sub_Gargoyle
	Sub_Germ
	Sub_Giant
	Sub_Gith
	Sub_Glimmer
	Sub_Gnoll
	Sub_Gnome
	Sub_Goat
	Sub_Goblin
	Sub_God
	Sub_Golem
	Sub_Gorgon
	Sub_Graveborn
	Sub_Gremlin
	Sub_Griffin
	Sub_Guest
	Sub_Hag
	Sub_Halfling
	Sub_Hamster
	Sub_Harpy
	Sub_Hellion
	Sub_Hero
	Sub_Hippo
	Sub_Hippogriff
	Sub_Homarid
	Sub_Homunculus
	Sub_Horror
	Sub_Horse
	Sub_Human
	Sub_Hydra
	Sub_Hyena
	Sub_Illusion
	Sub_Imp
	Sub_Incarnation
	Sub_Inkling
	Sub_Inquisitor
	Sub_Insect
	Sub_Jackal
	Sub_Jellyfish
	Sub_Juggernaut
	Sub_Kavu
	Sub_Kirin
	Sub_Kithkin
	Sub_Knight
	Sub_Kobold
	Sub_Kor
	Sub_Kraken
	Sub_Llama
	Sub_Lamia
	Sub_Lammasu
	Sub_Leech
	Sub_Leviathan
	Sub_Lhurgoyf
	Sub_Licid
	Sub_Lizard
	Sub_Manticore
	Sub_Masticore
	Sub_Mercenary
	Sub_Merfolk
	Sub_Metathran
	Sub_Minion
	Sub_Minotaur
	Sub_Mite
	Sub_Mole
	Sub_Monger
	Sub_Mongoose
	Sub_Monk
	Sub_Monkey
	Sub_Moogle
	Sub_Moonfolk
	Sub_Mount
	Sub_Mouse
	Sub_Mutant
	Sub_Myr
	Sub_Mystic
	Sub_Nautilus
	Sub_Necron
	Sub_Nephilim
	Sub_Nightmare
	Sub_Nightstalker
	Sub_Ninja
	Sub_Noble
	Sub_Noggle
	Sub_Nomad
	Sub_Nymph
	Sub_Octopus
	Sub_Ogre
	Sub_Ooze
	Sub_Orb
	Sub_Orc
	Sub_Orgg
	Sub_Otter
	Sub_Ouphe
	Sub_Ox
	Sub_Oyster
	Sub_Pangolin
	Sub_Peasant
	Sub_Pegasus
	Sub_Pentavite
	Sub_Performer
	Sub_Pest
	Sub_Phelddagrif
	Sub_Phoenix
	Sub_Phyrexian
	Sub_Pilot
	Sub_Pincher
	Sub_Pirate
	Sub_Plant
	Sub_Porcupine
	Sub_Possum
	Sub_Praetor
	Sub_Primarch
	Sub_Prism
	Sub_Processor
	Sub_Qu
	Sub_Rabbit
	Sub_Raccoon
	Sub_Ranger
	Sub_Rat
	Sub_Rebel
	Sub_Reflection
	Sub_Rhino
	Sub_Rigger
	Sub_Robot
	Sub_Rogue
	Sub_Sable
	Sub_Salamander
	Sub_Samurai
	Sub_Sand
	Sub_Saproling
	Sub_Satyr
	Sub_Scarecrow
	Sub_Scientist
	Sub_Scion
	Sub_Scorpion
	Sub_Scout
	Sub_Sculpture
	Sub_Seal
	Sub_Serf
	Sub_Serpent
	Sub_Servo
	Sub_Shade
	Sub_Shaman
	Sub_Shapeshifter
	Sub_Shark
	Sub_Sheep
	Sub_Siren
	Sub_Skeleton
	Sub_Skunk
	Sub_Slith
	Sub_Sliver
	Sub_Sloth
	Sub_Slug
	Sub_Snail
	Sub_Snake
	Sub_Soldier
	Sub_Soltari
	Sub_Spawn
	Sub_Specter
	Sub_Spellshaper
	Sub_Sphinx
	Sub_Spider
	Sub_Spike
	Sub_Spirit
	Sub_Splinter
	Sub_Sponge
	Sub_Squid
	Sub_Squirrel
	Sub_Starfish
	Sub_Surrakar
	Sub_Survivor
	Sub_Synth
	Sub_Tentacle
	Sub_Tetravite
	Sub_Thalakos
	Sub_Thopter
	Sub_Thrull
	Sub_Tiefling
	Sub_Toy
	Sub_Treefolk
	Sub_Trilobite
	Sub_Triskelavite
	Sub_Troll
	Sub_Turtle
	Sub_Tyranid
	Sub_Unicorn
	Sub_Vampire
	Sub_Varmint
	Sub_Vedalken
	Sub_Volver
	Sub_Wall
	Sub_Walrus
	Sub_Warlock
	Sub_Warrior
	Sub_Weasel
	Sub_Weird
	Sub_Werewolf
	Sub_Whale
	Sub_Wizard
	Sub_Wolf
	Sub_Wolverine
	Sub_Wombat
	Sub_Worm
	Sub_Wraith
	Sub_Wurm
	Sub_Yeti
	Sub_Zombie
	Sub_Zubera

	// Plane subtypes
	Sub_TheAbyss
	Sub_Alara
	Sub_AlfavaMetraxis
	Sub_Amonkhet
	Sub_AndrozaniMinor
	Sub_Antausia
	Sub_Apalapucia
	Sub_Arcavios
	Sub_Arkhos
	Sub_Avishkar
	Sub_Azgol
	Sub_Belenon
	Sub_BolasMeditationRealm
	Sub_Capenna
	Sub_Cridhe
	Sub_TheDalekAsylum
	Sub_Darillium
	Sub_Dominaria
	Sub_Earth
	Sub_Echoir
	Sub_Eldraine
	Sub_Equilor
	Sub_Ergamon
	Sub_Fabacin
	Sub_Fiora
	Sub_Gallifrey
	Sub_Gargantikar
	Sub_Gobakhan
	Sub_HorseheadNebula
	Sub_Ikoria
	Sub_Innistrad
	Sub_Iquatana
	Sub_Ir
	Sub_Ixalan
	Sub_Kaldheim
	Sub_Kamigawa
	Sub_Kandoka
	Sub_Karsus
	Sub_Kephalai
	Sub_Kinshala
	Sub_Kolbahan
	Sub_Kylem
	Sub_Kyneth
	Sub_TheLibrary
	Sub_Lorwyn
	Sub_Luvion
	Sub_Mars
	Sub_Mercadia
	Sub_Mirrodin
	Sub_Moag
	Sub_Mongseng
	Sub_Moon
	Sub_Muraganda
	Sub_Necros
	Sub_NewEarth
	Sub_NewPhyrexia
	Sub_OutsideMuttersSpiral
	Sub_Phyrexia
	Sub_Pyrulea
	Sub_Rabiah
	Sub_Rath
	Sub_Ravnica
	Sub_Regatha
	Sub_Segovia
	Sub_SerrasRealm
	Sub_Shadowmoor
	Sub_Shandalar
	Sub_Shenmeng
	Sub_Skaro
	Sub_Spacecraft
	Sub_Tarkir
	Sub_Theros
	Sub_Time
	Sub_Trenzalore
	Sub_Ulgrotha
	Sub_UnknownPlanet
	Sub_Valla
	Sub_Vryn
	Sub_Wildfire
	Sub_Xerex
	Sub_Zendikar
	Sub_Zhalfir

	// Dungeon subtype
	Sub_UndercitySiege

	// Battle subtype
	Sub_Siege
)

var subtypeNames = [...]string{
	// Planeswalkers
	"Ajani", "Aminatou", "Angrath", "Arlinn", "Ashiok", "Bahamut", "Basri",
	"Bolas", "Calix", "Chandra", "Comet", "Dack", "Dakkon", "Daretti",
	"Davriel", "Dihada", "Domri", "Dovin", "Ellywick", "Elminster",
	"Elspeth", "Estrid", "Freyalise", "Garruk", "Gideon", "Grist", "Guff",
	"Huatli", "Jace", "Jared", "Jaya", "Jeska", "Kaito", "Karn", "Kasmina",
	"Kaya", "Kiora", "Koth", "Liliana", "Lolth", "Lukka", "Minsc",
	"Mordenkainen", "Nahiri", "Narset", "Niko", "Nissa", "Nixilis", "Oko",
	"Quintorius", "Ral", "Rowan", "Saheeli", "Samut", "Sarkhan", "Serra",
	"Sivitri", "Sorin", "Szat", "Tamiyo", "Tasha", "Teferi", "Teyo",
	"Tezzeret", "Tibalt", "Tyvar", "Ugin", "Urza", "Venser", "Vivien",
	"Vraska", "Vronos", "Will", "Windgrace", "Wrenn", "Xenagos", "Yanggu",
	"Yanling", "Zariel",

	// Special spell subtypes
	"Adventure", "Arcane", "Lesson", "Omen", "Trap",

	// Other special subtype
	"Time Lord",

	// Creature/other subtypes A–Z
	"Advisor", "Aetherborn", "Alien", "Ally", "Angel", "Antelope", "Ape",
	"Archer", "Archon", "Armadillo", "Army", "Artificer", "Assassin",
	"Assembly-Worker", "Astartes", "Atog", "Aurochs", "Avatar", "Azra",
	"Badger", "Balloon", "Barbarian", "Bard", "Basilisk", "Bat", "Bear",
	"Beast", "Beaver", "Beeble", "Beholder", "Berserker", "Bird", "Blinkmoth",
	"Boar", "Bringer", "Brushwagg", "Camarid", "Camel", "Capybara", "Caribou",
	"Carrier", "Cat", "Centaur", "Child", "Chimera", "Citizen", "Cleric",
	"Clown", "Cockatrice", "Construct", "Coward", "Coyote", "Crab",
	"Crocodile", "Ctan", "Custodes", "Cyberman", "Cyclops", "Dalek", "Dauthi",
	"Demigod", "Demon", "Deserter", "Detective", "Devil", "Dinosaur", "Djinn",
	"Doctor", "Dog", "Dragon", "Drake", "Dreadnought", "Drone", "Druid",
	"Dryad", "Dwarf", "Efreet", "Egg", "Elder", "Eldrazi", "Elemental",
	"Elephant", "Elf", "Elk", "Employee", "Eye", "Faerie", "Ferret", "Fish",
	"Flagbearer", "Fox", "Fractal", "Frog", "Fungus", "Gamer", "Gargoyle",
	"Germ", "Giant", "Gith", "Glimmer", "Gnoll", "Gnome", "Goat", "Goblin",
	"God", "Golem", "Gorgon", "Graveborn", "Gremlin", "Griffin", "Guest",
	"Hag", "Halfling", "Hamster", "Harpy", "Hellion", "Hero", "Hippo",
	"Hippogriff", "Homarid", "Homunculus", "Horror", "Horse", "Human",
	"Hydra", "Hyena", "Illusion", "Imp", "Incarnation", "Inkling",
	"Inquisitor", "Insect", "Jackal", "Jellyfish", "Juggernaut", "Kavu",
	"Kirin", "Kithkin", "Knight", "Kobold", "Kor", "Kraken", "Llama",
	"Lamia", "Lammasu", "Leech", "Leviathan", "Lhurgoyf", "Licid", "Lizard",
	"Manticore", "Masticore", "Mercenary", "Merfolk", "Metathran", "Minion",
	"Minotaur", "Mite", "Mole", "Monger", "Mongoose", "Monk", "Monkey",
	"Moogle", "Moonfolk", "Mount", "Mouse", "Mutant", "Myr", "Mystic",
	"Nautilus", "Necron", "Nephilim", "Nightmare", "Nightstalker", "Ninja",
	"Noble", "Noggle", "Nomad", "Nymph", "Octopus", "Ogre", "Ooze", "Orb",
	"Orc", "Orgg", "Otter", "Ouphe", "Ox", "Oyster", "Pangolin", "Peasant",
	"Pegasus", "Pentavite", "Performer", "Pest", "Phelddagrif", "Phoenix",
	"Phyrexian", "Pilot", "Pincher", "Pirate", "Plant", "Porcupine",
	"Possum", "Praetor", "Primarch", "Prism", "Processor", "Qu", "Rabbit",
	"Raccoon", "Ranger", "Rat", "Rebel", "Reflection", "Rhino", "Rigger",
	"Robot", "Rogue", "Sable", "Salamander", "Samurai", "Sand", "Saproling",
	"Satyr", "Scarecrow", "Scientist", "Scion", "Scorpion", "Scout",
	"Sculpture", "Seal", "Serf", "Serpent", "Servo", "Shade", "Shaman",
	"Shapeshifter", "Shark", "Sheep", "Siren", "Skeleton", "Skunk", "Slith",
	"Sliver", "Sloth", "Slug", "Snail", "Snake", "Soldier", "Soltari",
	"Spawn", "Specter", "Spellshaper", "Sphinx", "Spider", "Spike", "Spirit",
	"Splinter", "Sponge", "Squid", "Squirrel", "Starfish", "Surrakar",
	"Survivor", "Synth", "Tentacle", "Tetravite", "Thalakos", "Thopter",
	"Thrull", "Tiefling", "Toy", "Treefolk", "Trilobite", "Triskelavite",
	"Troll", "Turtle", "Tyranid", "Unicorn", "Vampire", "Varmint",
	"Vedalken", "Volver", "Wall", "Walrus", "Warlock", "Warrior", "Weasel",
	"Weird", "Werewolf", "Whale", "Wizard", "Wolf", "Wolverine", "Wombat",
	"Worm", "Wraith", "Wurm", "Yeti", "Zombie", "Zubera",

	// Plane subtypes
	"The Abyss", "Alara", "Alfava Metraxis", "Amonkhet", "Androzani Minor",
	"Antausia", "Apalapucia", "Arcavios", "Arkhos", "Avishkar", "Azgol",
	"Belenon", "Bolas's Meditation Realm", "Capenna", "Cridhe",
	"The Dalek Asylum", "Darillium", "Dominaria", "Earth", "Echoir",
	"Eldraine", "Equilor", "Ergamon", "Fabacin", "Fiora", "Gallifrey",
	"Gargantikar", "Gobakhan", "Horsehead Nebula", "Ikoria", "Innistrad",
	"Iquatana", "Ir", "Ixalan", "Kaldheim", "Kamigawa", "Kandoka", "Karsus",
	"Kephalai", "Kinshala", "Kolbahan", "Kylem", "Kyneth", "The Library",
	"Lorwyn", "Luvion", "Mars", "Mercadia", "Mirrodin", "Moag", "Mongseng",
	"Moon", "Muraganda", "Necros", "New Earth", "New Phyrexia",
	"Outside Mutter's Spiral", "Phyrexia", "Pyrulea", "Rabiah", "Rath",
	"Ravnica", "Regatha", "Segovia", "Serra's Realm", "Shadowmoor",
	"Shandalar", "Shenmeng", "Skaro", "Spacecraft", "Tarkir", "Theros",
	"Time", "Trenzalore", "Ulgrotha", "Unknown Planet", "Valla", "Vryn",
	"Wildfire", "Xerex", "Zendikar", "Zhalfir",

	// Dungeon & Battle subtypes
	"Undercity Siege", "Siege",
}

var subtypeLookup = func() map[string]SubtypeId {
	m := make(map[string]SubtypeId, len(subtypeNames))
	for i, name := range subtypeNames {
		m[name] = SubtypeId(i)
	}
	return m
}()

func (s SubtypeId) String() string {
	if int(s) < len(subtypeNames) {
		return subtypeNames[s]
	}
	return fmt.Sprintf("SubtypeId(%d)", s)
}

func (s SubtypeId) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *SubtypeId) UnmarshalText(text []byte) error {
	str := string(text)
	if v, ok := subtypeLookup[str]; ok {
		*s = v
		return nil
	}
	for k, v := range subtypeLookup {
		if strings.EqualFold(k, str) {
			*s = v
			return nil
		}
	}
	return fmt.Errorf("unknown SubtypeId %q", str)
}

type TypeLineComponent struct {
	SuperMask uint8       // up to 8 supertypes
	TypeMask  uint16      // up to 16 types
	Subtypes  []SubtypeId // usually 0–3 entries
}

func (tl *TypeLineComponent) HasSuper(s Supertype) bool {
	return tl.SuperMask&(1<<s) != 0
}

func (tl *TypeLineComponent) HasType(t TypeId) bool {
	return tl.TypeMask&(1<<t) != 0
}

func (tl *TypeLineComponent) HasSubtype(st SubtypeId) bool {
	return slices.Contains(tl.Subtypes, st)
}

type TypeLineConfig struct {
	Supertypes []Supertype `toml:"supertypes,omitempty"`
	Types      []TypeId    `toml:"types"`
	Subtypes   []SubtypeId `toml:"subtypes,omitempty"`
}

func (cfg TypeLineConfig) ToComponent() TypeLineComponent {
	var tl TypeLineComponent
	for _, s := range cfg.Supertypes {
		tl.SuperMask |= 1 << s
	}
	for _, t := range cfg.Types {
		tl.TypeMask |= 1 << t
	}
	tl.Subtypes = append(tl.Subtypes, cfg.Subtypes...)
	return tl
}

// ParseTypeLine splits a Scryfall type_line (ie. "Creature — Griffin")
// into supertypes, types and subtypes.
func ParseTypeLine(s string) TypeLineConfig {
	// s is like "Legendary Artifact Creature — Human Wizard"
	parts := strings.Split(s, "—")
	left := strings.Fields(strings.TrimSpace(parts[0]))
	var sups []Supertype
	var tys []TypeId
	for _, tok := range left {
		// try supertype
		var st Supertype
		if err := st.UnmarshalText([]byte(tok)); err == nil {
			sups = append(sups, st)
			continue
		}
		// fall back to TypeId
		var t TypeId
		if err := t.UnmarshalText([]byte(tok)); err == nil {
			tys = append(tys, t)
			continue
		}
	}

	var subs []SubtypeId
	if len(parts) > 1 {
		for _, tok := range strings.Fields(strings.TrimSpace(parts[1])) {
			var st SubtypeId
			if err := st.UnmarshalText([]byte(tok)); err == nil {
				subs = append(subs, st)
			}
		}
	}

	return TypeLineConfig{
		Supertypes: sups,
		Types:      tys,
		Subtypes:   subs,
	}
}

func (tl TypeLineComponent) String() string {
	var parts []string

	// supertypes
	for i := Supertype(0); i < supertypeCount; i++ {
		if tl.SuperMask&(1<<i) != 0 {
			parts = append(parts, i.String())
		}
	}

	// types
	for t := Artifact; t < typeCount; t++ {
		if tl.TypeMask&(1<<t) != 0 {
			parts = append(parts, t.String())
		}
	}

	out := strings.Join(parts, " ")
	// subtypes
	if len(tl.Subtypes) > 0 {
		var subs []string
		for _, st := range tl.Subtypes {
			subs = append(subs, st.String())
		}
		out += " — " + strings.Join(subs, " ")
	}
	return out
}
