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
		// build path by adding .png
		filename := ill.AssetKey + ".png"
		path := filepath.Join(assetDir, filename)

		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("illustration not found %q: %w", path, err)
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		if _, format, err := image.DecodeConfig(f); err != nil {
			f.Close()
			return fmt.Errorf("cannot decode %q: %w", path, err)
		} else {
			fmt.Printf("Loading art %q as %s\n", path, format)
		}
		f.Close()

		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return fmt.Errorf("ebiten load %q: %w", path, err)
		}

		tmp := ill
		tmp.Image = img
		g.Illustrations[uint32(idx)] = tmp
	}
	return nil
}
