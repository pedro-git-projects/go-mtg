package component

import (
	"fmt"
)

type ManaCostComponent struct {
	Generic   uint8 `toml:"generic,omitempty"`
	Colorless uint8 `toml:"colorless,omitempty"`

	White uint8 `toml:"white,omitempty"`
	Blue  uint8 `toml:"blue,omitempty"`
	Black uint8 `toml:"black,omitempty"`
	Red   uint8 `toml:"red,omitempty"`
	Green uint8 `toml:"green,omitempty"`

	PhyWhite uint8 `toml:"phy_white,omitempty"` // {W/P}
	PhyBlue  uint8 `toml:"phy_blue,omitempty"`  // {U/P}
	PhyBlack uint8 `toml:"phy_black,omitempty"` // {B/P}
	PhyRed   uint8 `toml:"phy_red,omitempty"`   // {R/P}
	PhyGreen uint8 `toml:"phy_green,omitempty"` // {G/P}

	HybridWU uint8 `toml:"hybrid_wu,omitempty"` // {W/U}
	HybridWB uint8 `toml:"hybrid_wb,omitempty"` // {W/B}
	HybridWG uint8 `toml:"hybrid_wg,omitempty"` // {W/G}
	HybridWR uint8 `toml:"hybrid_wr,omitempty"` // {W/R}
	HybridUB uint8 `toml:"hybrid_ub,omitempty"` // {U/B}
	HybridUG uint8 `toml:"hybrid_ug,omitempty"` // {U/G}
	HybridUR uint8 `toml:"hybrid_ur,omitempty"` // {U/R}
	HybridBG uint8 `toml:"hybrid_bg,omitempty"` // {B/G}
	HybridBR uint8 `toml:"hybrid_br,omitempty"` // {B/R}
	HybridGR uint8 `toml:"hybrid_gr,omitempty"` // {G/R}
}

func (m ManaCostComponent) MarshalTOML() ([]byte, error) {
	var buf []byte
	// generic
	if m.Generic > 0 {
		buf = fmt.Appendf(buf, "generic = %d\n", m.Generic)
	}
	// colorless
	if m.Colorless > 0 {
		buf = fmt.Appendf(buf, "colorless = %d\n", m.Colorless)
	}
	// colored
	if m.White > 0 {
		buf = fmt.Appendf(buf, "white = %d\n", m.White)
	}
	if m.Blue > 0 {
		buf = fmt.Appendf(buf, "blue  = %d\n", m.Blue)
	}
	if m.Black > 0 {
		buf = fmt.Appendf(buf, "black = %d\n", m.Black)
	}
	if m.Red > 0 {
		buf = fmt.Appendf(buf, "red   = %d\n", m.Red)
	}
	if m.Green > 0 {
		buf = fmt.Appendf(buf, "green = %d\n", m.Green)
	}
	// phyrexian
	if m.PhyWhite > 0 {
		buf = fmt.Appendf(buf, "phy_white = %d\n", m.PhyWhite)
	}
	if m.PhyBlue > 0 {
		buf = fmt.Appendf(buf, "phy_blue  = %d\n", m.PhyBlue)
	}
	if m.PhyBlack > 0 {
		buf = fmt.Appendf(buf, "phy_black = %d\n", m.PhyBlack)
	}
	if m.PhyRed > 0 {
		buf = fmt.Appendf(buf, "phy_red   = %d\n", m.PhyRed)
	}
	if m.PhyGreen > 0 {
		buf = fmt.Appendf(buf, "phy_green = %d\n", m.PhyGreen)
	}
	// hybrid
	if m.HybridWU > 0 {
		buf = fmt.Appendf(buf, "hybrid_wu = %d\n", m.HybridWU)
	}
	if m.HybridWB > 0 {
		buf = fmt.Appendf(buf, "hybrid_wb = %d\n", m.HybridWB)
	}
	if m.HybridWG > 0 {
		buf = fmt.Appendf(buf, "hybrid_wg = %d\n", m.HybridWG)
	}
	if m.HybridWR > 0 {
		buf = fmt.Appendf(buf, "hybrid_wr = %d\n", m.HybridWR)
	}
	if m.HybridUB > 0 {
		buf = fmt.Appendf(buf, "hybrid_ub = %d\n", m.HybridUB)
	}
	if m.HybridUG > 0 {
		buf = fmt.Appendf(buf, "hybrid_ug = %d\n", m.HybridUG)
	}
	if m.HybridUR > 0 {
		buf = fmt.Appendf(buf, "hybrid_ur = %d\n", m.HybridUR)
	}
	if m.HybridBG > 0 {
		buf = fmt.Appendf(buf, "hybrid_bg = %d\n", m.HybridBG)
	}
	if m.HybridBR > 0 {
		buf = fmt.Appendf(buf, "hybrid_br = %d\n", m.HybridBR)
	}
	if m.HybridGR > 0 {
		buf = fmt.Appendf(buf, "hybrid_gr = %d\n", m.HybridGR)
	}
	return buf, nil
}

func (m *ManaCostComponent) UnmarshalTOML(data any) error {
	tbl, ok := data.(map[string]any)
	if !ok {
		return fmt.Errorf("invalid TOML type for ManaCostComponent: %T", data)
	}
	// helper to read int64 into uint8
	read := func(key string, ptr *uint8) {
		if v, ok := tbl[key].(int64); ok {
			*ptr = uint8(v)
		}
	}
	read("generic", &m.Generic)
	read("colorless", &m.Colorless)
	read("white", &m.White)
	read("blue", &m.Blue)
	read("black", &m.Black)
	read("red", &m.Red)
	read("green", &m.Green)
	read("phy_white", &m.PhyWhite)
	read("phy_blue", &m.PhyBlue)
	read("phy_black", &m.PhyBlack)
	read("phy_red", &m.PhyRed)
	read("phy_green", &m.PhyGreen)
	read("hybrid_wu", &m.HybridWU)
	read("hybrid_wb", &m.HybridWB)
	read("hybrid_wg", &m.HybridWG)
	read("hybrid_wr", &m.HybridWR)
	read("hybrid_ub", &m.HybridUB)
	read("hybrid_ug", &m.HybridUG)
	read("hybrid_ur", &m.HybridUR)
	read("hybrid_bg", &m.HybridBG)
	read("hybrid_br", &m.HybridBR)
	read("hybrid_gr", &m.HybridGR)
	return nil
}
