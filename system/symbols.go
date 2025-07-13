package system

var symbolFiles = map[string]string{
	// generic numeric
	"One":   "One.png",
	"Two":   "Two.png",
	"Three": "Three.png",
	"Four":  "Four.png",
	"X":     "X.png",

	// colorless
	"C": "C.png",

	// colored
	"W": "W.png",
	"U": "U.png",
	"B": "B.png",
	"R": "R.png",
	"G": "G.png",

	// phyrexian
	"WP": "WP.png",
	"UP": "UP.png",
	"BP": "BP.png",
	"RP": "RP.png",
	"GP": "GP.png",

	// hybrids
	"WU": "WU.png",
	"WB": "WB.png",
	"WG": "GW.png",
	"WR": "RW.png",
	"UB": "UB.png",
	"UG": "GU.png",
	"UR": "UR.png",
	"BG": "BG.png",
	"BR": "BR.png",
	"GR": "RG.png",
}

// func LoadSymbols(g *game.Game, assetDir string) error {
// 	g.Symbols = make(map[string]*ebiten.Image, len(symbolFiles))
// 	for key, fname := range symbolFiles {
// 		path := filepath.Join(assetDir, fname)
// 		img, _, err := ebitenutil.NewImageFromFile(path)
// 		if err != nil {
// 			return fmt.Errorf("loading symbol %q from %q: %w", key, path, err)
// 		}
// 		g.Symbols[key] = img
// 	}
// 	return nil
// }
