package objects

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pikomonde/game-201231-endOfTheYear/tools"
)

// Square is a square
type Square struct {
	Room     *Room
	Size     tools.Vector2
	Position tools.Vector2
	Speed    tools.Vector2
}

func (obj *Square) Create() {
	rand.Seed(time.Now().Unix())
	degree := rand.Float64() * 2 * math.Pi
	obj.Size = tools.Vector2{8, 8}
	// fmt.Println(obj.Room.LayoutSize)
	obj.Position = tools.Vector2{
		rand.Float64() * (obj.Room.LayoutSize.X - obj.Size.X),
		rand.Float64() * (obj.Room.LayoutSize.Y - obj.Size.Y),
	}
	obj.Speed = tools.Vector2{1, 0}
	obj.Speed.Rotate(degree)
}

func (obj *Square) Update() {
	obj.Position.X += obj.Speed.X
	obj.Position.Y += obj.Speed.Y
	if (obj.Position.X < 0) || (obj.Position.X+obj.Size.X > obj.Room.LayoutSize.X) {
		obj.Speed.X = -obj.Speed.X
	}
	if (obj.Position.Y < 0) || (obj.Position.Y+obj.Size.Y > obj.Room.LayoutSize.Y) {
		obj.Speed.Y = -obj.Speed.Y
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(ebiten.MouseButtonLeft)) {
		mouseX, mouseY := ebiten.CursorPosition()
		if (mouseX >= int(obj.Position.X)) && (mouseX <= int(obj.Position.X+obj.Size.X)) &&
			(mouseY >= int(obj.Position.Y)) && (mouseY <= int(obj.Position.Y+obj.Size.Y)) {
			obj.Create()
			obj.Room.Rule.Score++
			// obj.Room.Rule.AudioPlayer.Rewind()
			// obj.Room.Rule.AudioPlayer.Play()
			obj.Room.Rule.SndMouseClick.Rewind()
			obj.Room.Rule.SndMouseClick.Play()
		}
	}
}

func (obj *Square) Draw() *ebiten.Image {
	img := ebiten.NewImage(int(obj.Size.X), int(obj.Size.Y))
	img.Fill(color.White)
	return img
}
