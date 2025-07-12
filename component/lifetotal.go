package component

import "fmt"

type LifeTotalComponent struct {
	Value int `toml:"value" json:"value"`
}

func (l LifeTotalComponent) MarshalTOML() ([]byte, error) {
	data := fmt.Appendf(nil, "value = %d\n", l.Value)
	return data, nil
}

func (l *LifeTotalComponent) UnmarshalTOML(data any) error {
	switch v := data.(type) {
	case int64:
		l.Value = int(v)
		return nil
	case float64:
		// TOML numeric can be float; convert safely
		l.Value = int(v)
		return nil
	default:
		return fmt.Errorf("invalid type %T for LifeTotalComponent", data)
	}
}
