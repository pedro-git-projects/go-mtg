package component

import "fmt"

type NameComponent struct {
	Value string `toml:"name" json:"value"`
}

func (n NameComponent) MarshalTOML() ([]byte, error) {
	data := fmt.Appendf(nil, "name = %q\n", n.Value)
	return data, nil
}

func (n *NameComponent) UnmarshalTOML(data any) error {
	s, ok := data.(string)
	if !ok {
		return fmt.Errorf("invalid type %T for NameComponent", data)
	}
	n.Value = s
	return nil
}
