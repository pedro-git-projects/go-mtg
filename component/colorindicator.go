package component

import (
	"fmt"
	"strings"
)

type Color uint8

const (
	ColorColorless Color = 0
	ColorBlack     Color = 1 << iota
	ColorWhite
	ColorGreen
	ColorRed
	ColorBlue
)

const AllColors = ColorBlack | ColorWhite | ColorGreen | ColorRed | ColorBlue

var _orderedColors = []struct {
	bit  Color
	name string
}{
	{ColorWhite, "White"},
	{ColorBlue, "Blue"},
	{ColorBlack, "Black"},
	{ColorRed, "Red"},
	{ColorGreen, "Green"},
}
var _fmtColorNames = map[Color]string{
	ColorBlack: "Black",
	ColorWhite: "White",
	ColorGreen: "Green",
	ColorRed:   "Red",
	ColorBlue:  "Blue",
}

func (c Color) Has(bit Color) bool     { return c&bit != 0 }
func (c Color) Add(bit Color) Color    { return c | bit }
func (c Color) Remove(bit Color) Color { return c &^ bit }
func (c Color) IsColorless() bool      { return c == ColorColorless }

func (c Color) String() string {
	if c.IsColorless() {
		return "Colorless"
	}
	var parts []string
	for _, entry := range _orderedColors {
		if c.Has(entry.bit) {
			parts = append(parts, entry.name)
		}
	}
	return strings.Join(parts, "+")
}

func (c Color) MarshalTOML() ([]byte, error) {
	// emit as a quoted string, e.g. "Red+Blue"
	return []byte(fmt.Appendf(nil, "%q", c.String())), nil
}

func (c *Color) UnmarshalTOML(v any) error {
	switch val := v.(type) {
	case string:
		return c.parseFromString(val)
	case []any:
		// handle ["Red","Blue"] like JSON-array style
		var parts []string
		for _, e := range val {
			if s, ok := e.(string); ok {
				parts = append(parts, s)
			} else {
				return fmt.Errorf("invalid element %T in TOML color list", e)
			}
		}
		return c.parseFromString(strings.Join(parts, "+"))
	default:
		return fmt.Errorf("unexpected type %T for Color in TOML", v)
	}
}

func (c *Color) parseFromString(s string) error {
	if s == "Colorless" {
		*c = ColorColorless
		return nil
	}
	var mask Color
	for part := range strings.SplitSeq(s, "+") {
		matched := false
		for bit, name := range _fmtColorNames {
			if name == part {
				mask |= bit
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Errorf("unknown color %q in TOML", part)
		}
	}
	*c = mask
	return nil
}

type ColorIndicatorComponent struct {
	Value Color `toml:"value"`
}
