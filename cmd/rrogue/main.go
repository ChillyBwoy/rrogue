package main

import (
	_ "image/png"
	"log"

	"github.com/ChillyBwoy/rrogue/internal/rrogue"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := rrogue.NewGame()
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowTitle("Tower")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
