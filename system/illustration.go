package system

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pedro-git-projects/go-mtg/game"
)

func LoadIllustrations(g *game.Game, assetDir string) error {
	for idx, ill := range g.Illustrations {
		path := filepath.Join(assetDir, ill.AssetKey)
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("illustration %q not found: %w", path, err)
		}
		// peek format
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		cfg, format, err := image.DecodeConfig(f)
		_ = cfg
		f.Close()
		if err != nil {
			return fmt.Errorf("cannot decode %q: %w", path, err)
		}
		fmt.Printf("Loading %q as format %s\n", path, format)

		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return fmt.Errorf("ebiten load %q: %w", path, err)
		}
		comp := ill
		comp.Image = img
		g.Illustrations[uint32(idx)] = comp
	}
	return nil
}
