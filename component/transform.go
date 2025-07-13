package component

type TransformComponent struct {
	X     float64 `toml:"x,omitempty" json:"x"`
	Y     float64 `toml:"y,omitempty" json:"y"`
	Scale float64 `toml:"scale,omitempty" json:"scale"`
}
