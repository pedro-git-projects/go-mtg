package component

import (
	"encoding/json"
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

// map from Scryfall’s single rune codes to Color bits
var _abbrToColor = map[string]Color{
	"W": ColorWhite,
	"U": ColorBlue,
	"B": ColorBlack,
	"R": ColorRed,
	"G": ColorGreen,
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

//	json.Unmarshal([]byte(`["B","R"]`), &c)
//
// and get c == ColorBlack|ColorRed.
func (c *Color) UnmarshalJSON(b []byte) error {
	// try an array of strings first
	var arr []string
	if err := json.Unmarshal(b, &arr); err != nil {
		// maybe it was just a single string
		var s string
		if err2 := json.Unmarshal(b, &s); err2 != nil {
			return fmt.Errorf("Color.UnmarshalJSON: invalid JSON %s", string(b))
		}
		arr = []string{s}
	}
	var mask Color
	for _, abbrev := range arr {
		bit, ok := _abbrToColor[abbrev]
		if !ok {
			return fmt.Errorf("Color.UnmarshalJSON: unknown color code %q", abbrev)
		}
		mask |= bit
	}
	*c = mask
	return nil
}

// MarshalJSON emits the same array form, e.g. ["R","G"].
func (c Color) MarshalJSON() ([]byte, error) {
	if c == ColorColorless {
		return json.Marshal([]string{})
	}
	var parts []string
	for abbr, bit := range _abbrToColor {
		if c.Has(bit) {
			parts = append(parts, abbr)
		}
	}
	return json.Marshal(parts)
}

func (ci *ColorIndicatorComponent) UnmarshalJSON(b []byte) error {
	return ci.Value.UnmarshalJSON(b)
}

func (ci ColorIndicatorComponent) MarshalJSON() ([]byte, error) {
	return ci.Value.MarshalJSON()
}

// ColorFromStrings returns the bitmask for a slice of single-letter color codes.
// An empty slice -> ColorColorless.
func ColorFromStrings(abbrevs []string) Color {
	if len(abbrevs) == 0 {
		return ColorColorless
	}
	var mask Color
	for _, a := range abbrevs {
		if bit, ok := _abbrToColor[a]; ok {
			mask |= bit
		}
	}
	return mask
}
