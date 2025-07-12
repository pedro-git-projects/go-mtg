package component

import "fmt"

type Zone int

const (
	Library Zone = iota
	Hand
	Battlefield
	Graveyard
	Stack
	Exile
	CommandZone
)

type ZoneComponent struct {
	Value Zone `toml:"zone" json:"value"`
}

func (z ZoneComponent) MarshalTOML() ([]byte, error) {
	data := fmt.Appendf(nil, "zone = %d\n", z.Value)
	return data, nil
}

func (z *ZoneComponent) UnmarshalTOML(data any) error {
	switch v := data.(type) {
	case int64:
		z.Value = Zone(v)
	case float64:
		z.Value = Zone(int(v))
	default:
		return fmt.Errorf("invalid type %T for ZoneComponent", data)
	}
	return nil
}
