package rrogue

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map         GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	world, tags := InitializeWorld(g.Map.CurrentLevel)
	g.WorldTags = tags
	g.World = world
	g.Turn = PlayerTurn
	g.TurnCounter = 0
	return g
}

func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 5 {
		TakePlayerAction(g)
	}
	if g.Turn == MonsterTurn {
		UpdateMonster(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	gd := NewGameData()

	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}
