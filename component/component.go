package component

type ComponentId uint

const (
	LifeTotal ComponentId = iota
	Name
	ManaCost
	ConvertedManaCost
	TypeLine
	OracleText
	Power
	Toughness
	Loyalty
	ColorIndicator
	ProducedMana
	Illustration
	Defense
	CardData
	HandModifier
	LifeModifier
	ZoneType
	Contains
	Owner
	Transform
)
