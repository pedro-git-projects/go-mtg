package game

import (
	"fmt"
	"strings"

	"github.com/pedro-git-projects/go-mtg/component"
)

func (g *Game) PlayerStr() string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	var sb strings.Builder
	for idx, p := range g.Players {
		ltComp := g.lifeTotals[p.ID().Index]
		sb.WriteString(fmt.Sprintf(
			"Player %d (Entity%d): %d life\n",
			idx+1,
			p.ID().Index,
			ltComp.Value,
		))
	}
	return sb.String()
}

func (g *Game) LibraryStr() string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	var sb strings.Builder
	for idx, player := range g.Players {
		sb.WriteString(fmt.Sprintf(
			"Player %d (Entity%d) Library:\n",
			idx+1,
			player.ID().Index,
		))

		found := false
		for zoneIdx, zComp := range g.zoneTypes {
			if zComp.Value != component.Library {
				continue
			}
			if g.owners[zoneIdx].Value != player.ID() {
				continue
			}
			found = true

			for _, cardID := range g.contains[zoneIdx].Value {
				name := g.names[cardID.Index].Value
				cost := g.manaCosts[cardID.Index]
				var parts []string
				// generic mana (any color)
				if cost.Generic > 0 {
					parts = append(parts, fmt.Sprintf("{%d}", cost.Generic))
				}
				// colorless
				for i := uint8(0); i < cost.Colorless; i++ {
					parts = append(parts, "{C}")
				}
				// colored mana
				for i := uint8(0); i < cost.White; i++ {
					parts = append(parts, "{W}")
				}
				for i := uint8(0); i < cost.Blue; i++ {
					parts = append(parts, "{U}")
				}
				for i := uint8(0); i < cost.Black; i++ {
					parts = append(parts, "{B}")
				}
				for i := uint8(0); i < cost.Red; i++ {
					parts = append(parts, "{R}")
				}
				for i := uint8(0); i < cost.Green; i++ {
					parts = append(parts, "{G}")
				}
				// phyrexian
				for i := uint8(0); i < cost.PhyWhite; i++ {
					parts = append(parts, "{W/P}")
				}
				for i := uint8(0); i < cost.PhyBlue; i++ {
					parts = append(parts, "{U/P}")
				}
				for i := uint8(0); i < cost.PhyBlack; i++ {
					parts = append(parts, "{B/P}")
				}
				for i := uint8(0); i < cost.PhyRed; i++ {
					parts = append(parts, "{R/P}")
				}
				for i := uint8(0); i < cost.PhyGreen; i++ {
					parts = append(parts, "{G/P}")
				}
				// hybrid
				for i := uint8(0); i < cost.HybridWU; i++ {
					parts = append(parts, "{W/U}")
				}
				for i := uint8(0); i < cost.HybridWB; i++ {
					parts = append(parts, "{W/B}")
				}
				for i := uint8(0); i < cost.HybridWG; i++ {
					parts = append(parts, "{W/G}")
				}
				for i := uint8(0); i < cost.HybridWR; i++ {
					parts = append(parts, "{W/R}")
				}
				for i := uint8(0); i < cost.HybridUB; i++ {
					parts = append(parts, "{U/B}")
				}
				for i := uint8(0); i < cost.HybridUG; i++ {
					parts = append(parts, "{U/G}")
				}
				for i := uint8(0); i < cost.HybridUR; i++ {
					parts = append(parts, "{U/R}")
				}
				for i := uint8(0); i < cost.HybridBG; i++ {
					parts = append(parts, "{B/G}")
				}
				for i := uint8(0); i < cost.HybridBR; i++ {
					parts = append(parts, "{B/R}")
				}
				for i := uint8(0); i < cost.HybridGR; i++ {
					parts = append(parts, "{G/R}")
				}
				costStr := strings.Join(parts, "")

				color := g.colorIndicators[cardID.Index].Value.String()

				// --- new: type‐line ---
				tl := g.typeLines[cardID.Index]
				// collect supertypes
				var superParts []string
				for s := component.Legendary; s <= component.Ongoing; s++ {
					if tl.HasSuper(s) {
						superParts = append(superParts, s.String())
					}
				}
				// collect types
				var typeParts []string
				for t := component.Artifact; t <= component.Vanguard; t++ {
					if tl.HasType(t) {
						typeParts = append(typeParts, t.String())
					}
				}
				// merge supers+types
				var lineParts []string
				lineParts = append(lineParts, superParts...)
				lineParts = append(lineParts, typeParts...)
				mainLine := strings.Join(lineParts, " ")

				// collect subtypes
				var subParts []string
				for _, st := range tl.Subtypes {
					subParts = append(subParts, st.String())
				}

				var typeLineStr string
				if len(subParts) > 0 {
					// e.g. "Legendary Creature — Elf Warrior"
					typeLineStr = fmt.Sprintf("%s — %s", mainLine, strings.Join(subParts, " "))
				} else {
					typeLineStr = mainLine
				}

				sb.WriteString(fmt.Sprintf(
					"  – %s %s [%s] (%s)\n",
					name, costStr, color, typeLineStr,
				))
			}
		}

		if !found {
			sb.WriteString("  – <no library found>\n")
		}
	}
	return sb.String()
}
