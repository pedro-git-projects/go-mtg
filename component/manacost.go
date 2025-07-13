package component

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type ManaCostComponent struct {
	// numeric/generic
	Generic   uint8
	Colorless uint8

	// colored
	White uint8
	Blue  uint8
	Black uint8
	Red   uint8
	Green uint8

	// phyrexian per color
	PhyWhite uint8
	PhyBlue  uint8
	PhyBlack uint8
	PhyRed   uint8
	PhyGreen uint8

	// hybrid two color
	HybridWU uint8
	HybridWB uint8
	HybridWG uint8
	HybridWR uint8
	HybridUB uint8
	HybridUG uint8
	HybridUR uint8
	HybridBG uint8
	HybridBR uint8
	HybridGR uint8

	// X/Y/P edge cases:
	X           uint8 // {X}
	Y           uint8 // {Y}
	Z           uint8 // {Z}
	H           uint8 // {H}  (phyrexian generic)
	HalfGeneric uint8 // {½}

	// half‐color
	HW uint8 // {HW}
	HR uint8 // {HR}

	// special mana types
	Snow            uint8 // {S}
	LegendarySource uint8 // {L}
}

// ParseManaCost does a single‐pass, zero‐allocation parse of a string like "{2}{W/U}{1/2}{S}{L}".
func ParseManaCost(s string) (m ManaCostComponent, _ error) {
	for i := 0; i < len(s); {
		if s[i] != '{' {
			i++
			continue
		}
		// find matching '}'
		j := i + 1
		for j < len(s) && s[j] != '}' {
			j++
		}
		if j >= len(s) {
			return m, fmt.Errorf("ParseManaCost: unmatched '{' in %q", s)
		}

		// length of content = j-(i+1)
		switch length := j - (i + 1); {
		case length == 0:
			// empty `{}` → ignore

		case length == 1:
			// single‐rune tokens: digit, letter or symbol
			r, _ := utf8.DecodeRuneInString(s[i+1 : j])
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// treat all digits as generic
				m.Generic += uint8(r - '0')
			case 'W':
				m.White++
			case 'U':
				m.Blue++
			case 'B':
				m.Black++
			case 'R':
				m.Red++
			case 'G':
				m.Green++
			case 'C':
				m.Colorless++
			case 'X':
				m.X++
			case 'Y':
				m.Y++
			case 'Z':
				m.Z++
			case 'H':
				m.H++
			case 'S':
				m.Snow++
			case 'L':
				m.LegendarySource++
			case '½':
				m.HalfGeneric++
			}

		case length == 2:
			// two‐char tokens: "HW", "HR"
			a := s[i+1]
			b := s[i+2]
			if a == 'H' && b == 'W' {
				m.HW++
			} else if a == 'H' && b == 'R' {
				m.HR++
			}

		case length == 3:
			// "A/B" or "P/W" etc. but without braces and slash at pos+1.
			a := s[i+1]
			b := s[i+3]
			if s[i+2] != '/' {
				break
			}
			switch {
			// phyrexian
			case a == 'W' && b == 'P':
				m.PhyWhite++
			case a == 'U' && b == 'P':
				m.PhyBlue++
			case a == 'B' && b == 'P':
				m.PhyBlack++
			case a == 'R' && b == 'P':
				m.PhyRed++
			case a == 'G' && b == 'P':
				m.PhyGreen++

			// hybrid two color
			case a == 'W' && b == 'U':
				m.HybridWU++
			case a == 'W' && b == 'B':
				m.HybridWB++
			case a == 'W' && b == 'G':
				m.HybridWG++
			case a == 'W' && b == 'R':
				m.HybridWR++
			case a == 'U' && b == 'B':
				m.HybridUB++
			case a == 'U' && b == 'G':
				m.HybridUG++
			case a == 'U' && b == 'R':
				m.HybridUR++
			case a == 'B' && b == 'G':
				m.HybridBG++
			case a == 'B' && b == 'R':
				m.HybridBR++
			case a == 'G' && b == 'R':
				m.HybridGR++
			}

		default:
			// all‐digits >1 rune: "{12}", "{1000000}" -> accumulate into Generic
			// (since 2 -> infinite counts as generic mana)
			var v int
			valid := true
			for k := i + 1; k < j; k++ {
				if d := s[k] - '0'; d >= 0 && d <= 9 {
					v = v*10 + int(d)
				} else {
					valid = false
					break
				}
			}
			if valid {
				m.Generic += uint8(v)
			}
		}

		i = j + 1
	}
	return m, nil
}

func (m ManaCostComponent) String() string {
	var parts []string

	// Generic (printed all at once as {N})
	if m.Generic > 0 {
		parts = append(parts, fmt.Sprintf("{%d}", m.Generic))
	}

	// Colorless symbols (C)
	for i := uint8(0); i < m.Colorless; i++ {
		parts = append(parts, "{C}")
	}

	// Single‐color symbols
	for i := uint8(0); i < m.White; i++ {
		parts = append(parts, "{W}")
	}
	for i := uint8(0); i < m.Blue; i++ {
		parts = append(parts, "{U}")
	}
	for i := uint8(0); i < m.Black; i++ {
		parts = append(parts, "{B}")
	}
	for i := uint8(0); i < m.Red; i++ {
		parts = append(parts, "{R}")
	}
	for i := uint8(0); i < m.Green; i++ {
		parts = append(parts, "{G}")
	}

	// Phyrexian per‐color
	for i := uint8(0); i < m.PhyWhite; i++ {
		parts = append(parts, "{W/P}")
	}
	for i := uint8(0); i < m.PhyBlue; i++ {
		parts = append(parts, "{U/P}")
	}
	for i := uint8(0); i < m.PhyBlack; i++ {
		parts = append(parts, "{B/P}")
	}
	for i := uint8(0); i < m.PhyRed; i++ {
		parts = append(parts, "{R/P}")
	}
	for i := uint8(0); i < m.PhyGreen; i++ {
		parts = append(parts, "{G/P}")
	}

	// Hybrid two-color
	for i := uint8(0); i < m.HybridWU; i++ {
		parts = append(parts, "{W/U}")
	}
	for i := uint8(0); i < m.HybridWB; i++ {
		parts = append(parts, "{W/B}")
	}
	for i := uint8(0); i < m.HybridWG; i++ {
		parts = append(parts, "{W/G}")
	}
	for i := uint8(0); i < m.HybridWR; i++ {
		parts = append(parts, "{W/R}")
	}
	for i := uint8(0); i < m.HybridUB; i++ {
		parts = append(parts, "{U/B}")
	}
	for i := uint8(0); i < m.HybridUG; i++ {
		parts = append(parts, "{U/G}")
	}
	for i := uint8(0); i < m.HybridUR; i++ {
		parts = append(parts, "{U/R}")
	}
	for i := uint8(0); i < m.HybridBG; i++ {
		parts = append(parts, "{B/G}")
	}
	for i := uint8(0); i < m.HybridBR; i++ {
		parts = append(parts, "{B/R}")
	}
	for i := uint8(0); i < m.HybridGR; i++ {
		parts = append(parts, "{G/R}")
	}

	// “X/Y/P” edge-cases
	for i := uint8(0); i < m.X; i++ {
		parts = append(parts, "{X}")
	}
	for i := uint8(0); i < m.Y; i++ {
		parts = append(parts, "{Y}")
	}
	for i := uint8(0); i < m.Z; i++ {
		parts = append(parts, "{Z}")
	}
	for i := uint8(0); i < m.H; i++ {
		parts = append(parts, "{H}")
	}
	for i := uint8(0); i < m.HalfGeneric; i++ {
		parts = append(parts, "{½}")
	}

	// Half‐color
	for i := uint8(0); i < m.HW; i++ {
		parts = append(parts, "{HW}")
	}
	for i := uint8(0); i < m.HR; i++ {
		parts = append(parts, "{HR}")
	}

	// Special mana types
	for i := uint8(0); i < m.Snow; i++ {
		parts = append(parts, "{S}")
	}
	for i := uint8(0); i < m.LegendarySource; i++ {
		parts = append(parts, "{L}")
	}

	return strings.Join(parts, "")
}
