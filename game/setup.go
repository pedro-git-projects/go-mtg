package game

func SetupStandardGameFromTOML(path string) (*Game, error) {
	cfg, err := LoadConfig(path)
	if err != nil {
		return nil, err
	}

	// compute a rough capacity: players + total cards + one zone per player
	totalCards := 0
	for _, lib := range cfg.Libraries {
		totalCards += len(lib.Cards)
	}
	g := NewGame(cfg.Players + totalCards + cfg.Players)

	cardSlices := make([][]CardConfig, len(cfg.Libraries))
	for i, lib := range cfg.Libraries {
		cardSlices[i] = lib.Cards
	}

	g.SpawnPlayersWithLibraries(cfg.Players, cfg.InitialLife, cardSlices)
	return g, nil
}
