package entity

import "fmt"

type EntityId struct {
	Index      uint32
	Generation uint32
}

type Entity struct {
	id   EntityId
	mask ComponentMask
}

func (e Entity) ID() EntityId {
	return e.id
}

func (e Entity) Mask() ComponentMask {
	return e.mask
}

func (e *Entity) AddComponent(compID uint) {
	e.mask.Set(compID)
}

func (e *Entity) RemoveComponent(compID uint) {
	e.mask.Clear(compID)
}

func (e Entity) HasComponent(compID uint) bool {
	return e.mask.Has(compID)
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity{Index:%d,Gen:%d}", e.id.Index, e.id.Generation)
}
