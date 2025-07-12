package component

type ComponentId uint

const (
	LifeTotal ComponentId = iota
	Name
	ManaCost
	Illustration
	ColorIndicator
	TypeLine
	PowerAndToughness
	Loyalty
	Defense
	HandModifier
	LifeModifier
	ExpansionSymbol
	IllustrationCredit
	LegalText
	CollectorNumber
	ZoneType
	Contains
	Owner
)
