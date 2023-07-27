package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(700, 500)
	ebiten.SetWindowTitle("snakegame")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
