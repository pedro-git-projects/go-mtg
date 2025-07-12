package component

import (
	"fmt"

	"github.com/pedro-git-projects/go-mtg/entity"
)

type OwnerComponent struct {
	Value entity.EntityId `toml:"owner" json:"value"`
}

func (o OwnerComponent) MarshalTOML() ([]byte, error) {
	buf := fmt.Appendf(nil, "owner = { index = %d, generation = %d }\n", o.Value.Index, o.Value.Generation)
	return buf, nil
}

func (o *OwnerComponent) UnmarshalTOML(v any) error {
	m, ok := v.(map[string]any)
	if !ok {
		return fmt.Errorf("invalid type %T for OwnerComponent", v)
	}
	idx, ok1 := m["index"].(int64)
	gen, ok2 := m["generation"].(int64)
	if !ok1 || !ok2 {
		return fmt.Errorf("invalid index/generation for OwnerComponent: %v", m)
	}
	o.Value = entity.EntityId{Index: uint32(idx), Generation: uint32(gen)}
	return nil
}
