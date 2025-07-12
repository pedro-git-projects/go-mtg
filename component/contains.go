package component

import (
	"fmt"

	"github.com/pedro-git-projects/go-mtg/entity"
)

type ContainsComponent struct {
	Value []entity.EntityId `toml:"contains" json:"value"`
}

// MarshalTOML emits this component as a TOML array of tables:
// [[contains]]
// index = X
// generation = Y
func (c ContainsComponent) MarshalTOML() ([]byte, error) {
	var buf []byte
	for _, id := range c.Value {
		buf = fmt.Appendf(buf, "[[contains]]\n")
		buf = fmt.Appendf(buf, "index = %d\n", id.Index)
		buf = fmt.Appendf(buf, "generation = %d\n\n", id.Generation)
	}
	return buf, nil
}

func (c *ContainsComponent) UnmarshalTOML(v any) error {
	arr, ok := v.([]any)
	if !ok {
		return fmt.Errorf("invalid type %T for ContainsComponent", v)
	}
	c.Value = make([]entity.EntityId, 0, len(arr))
	for _, item := range arr {
		m, ok := item.(map[string]any)
		if !ok {
			return fmt.Errorf("invalid element %T in ContainsComponent", item)
		}
		idx, ok1 := m["index"].(int64)
		gen, ok2 := m["generation"].(int64)
		if !ok1 || !ok2 {
			return fmt.Errorf("invalid index/generation in ContainsComponent element: %v", m)
		}
		c.Value = append(c.Value, entity.EntityId{Index: uint32(idx), Generation: uint32(gen)})
	}
	return nil
}
