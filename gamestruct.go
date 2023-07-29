package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

var globalX float32 = 0
var globalY float32 = 0

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, globalX, globalY, 10, 10, color.White, false)

	limitX := screen.Bounds().Dx()
	limitY := screen.Bounds().Dy()
	globalX = float32(int(globalX+1) % (limitX))
	globalY = float32(int(globalY+1) % (limitY))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
