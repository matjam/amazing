package main

import (
	"github.com/rs/zerolog/log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/matjam/amazing/internal/game"
)

const (
	sWidth  = 480
	sHeight = 270
)

func main() {
	ebiten.SetWindowSize(sWidth, sHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetFullscreen(true)
	ebiten.SetMaxTPS(60)
	ebiten.SetVsyncEnabled(true)

	g := game.NewGame(sWidth, sHeight)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal().Err(err).Msg("fatal error")
	}
}
