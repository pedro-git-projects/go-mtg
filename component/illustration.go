package component

import "github.com/hajimehoshi/ebiten/v2"

type IllustrationComponent struct {
	AssetKey string        `toml:"asset_key"` // e.g. "snapcaster_mage.png" or "cards/snapcaster_mage"
	Image    *ebiten.Image // populated by the resource loader system
}
