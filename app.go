package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pikomonde/game-201231-endOfTheYear/objects"
	"github.com/pikomonde/game-201231-endOfTheYear/tools"
)

type Game struct {
	Room *objects.Room
}

func (g *Game) Update() error {
	g.Room.Square.Update()
	g.Room.Rule.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	{
		img := g.Room.Square.Draw()
		geom := ebiten.GeoM{}
		geom.Translate(g.Room.Square.Position.X, g.Room.Square.Position.Y)
		screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geom})
	}
	{
		img := g.Room.Rule.Draw()
		// geom := ebiten.GeoM{}
		// geom.Translate(g.Room.Rule.Position.X, g.Room.Rule.Position.Y)
		screen.DrawImage(img, &ebiten.DrawImageOptions{})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.Room.LayoutSize.X), int(g.Room.LayoutSize.Y)
}

func main() {
	ebiten.SetWindowSize(380, 600)
	game := &Game{
		Room: &objects.Room{
			LayoutSize: tools.Vector2{190, 300},
			Rule:       &objects.Rule{},
			Square:     &objects.Square{},
		},
	}
	game.Room.Rule.Room = game.Room
	game.Room.Square.Room = game.Room

	// Create all objects
	game.Room.Rule.Create()
	game.Room.Square.Create()

	ebiten.RunGame(game)
}
