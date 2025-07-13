package component

import (
	"encoding/json"
	"fmt"
)

type NameComponent struct {
	Value string `json:"name" toml:"name"`
}

func (n *NameComponent) UnmarshalTOML(data any) error {
	s, ok := data.(string)
	if !ok {
		return fmt.Errorf("invalid type %T for NameComponent", data)
	}
	n.Value = s
	return nil
}

func (n *NameComponent) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	n.Value = s
	return nil
}

func ParseName(s string) NameComponent {
	return NameComponent{Value: s}
}
