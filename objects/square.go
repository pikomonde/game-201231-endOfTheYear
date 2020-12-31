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
	Degree   float64
	Speed    float64
}

func (obj *Square) Create() {
	rand.Seed(time.Now().Unix())
	obj.Degree = rand.Float64() * 360
	obj.Size = tools.Vector2{8, 8}
	obj.Position = tools.Vector2{
		rand.Float64() * (obj.Room.LayoutSize.X - obj.Size.X),
		rand.Float64() * (obj.Room.LayoutSize.Y - obj.Size.Y),
	}
	obj.Speed = 1
}

func (obj *Square) Update() {
	speedVector := tools.Vector2{obj.Speed, 0}
	speedVector.Rotate((obj.Degree / 360) * 2 * math.Pi)
	obj.Position.X += speedVector.X
	obj.Position.Y += speedVector.Y
	if (obj.Position.X < 0) || (obj.Position.X+obj.Size.X > obj.Room.LayoutSize.X) {
		// obj.Speed.X = -obj.Speed.X
		if (obj.Degree >= 0) && (obj.Degree < 180) {
			obj.Degree = 180 - obj.Degree
		} else {
			obj.Degree = 540 - obj.Degree
		}
	}
	if (obj.Position.Y < 0) || (obj.Position.Y+obj.Size.Y > obj.Room.LayoutSize.Y) {
		// obj.Speed.Y = -obj.Speed.Y
		obj.Degree = 360 - obj.Degree
	}
	if obj.Speed < 1 {
		obj.Speed += 0.01
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(ebiten.MouseButtonLeft)) {
		mouseX, mouseY := ebiten.CursorPosition()
		if (mouseX >= int(obj.Position.X)) && (mouseX <= int(obj.Position.X+obj.Size.X)) &&
			(mouseY >= int(obj.Position.Y)) && (mouseY <= int(obj.Position.Y+obj.Size.Y)) {
			obj.Speed = 0.6 * obj.Speed
			obj.Room.Rule.SndMouseClick.Rewind()
			obj.Room.Rule.SndMouseClick.Play()
			if obj.Speed <= 0.2 {
				obj.Create()
				obj.Room.Rule.Score++
			}
		}
	}
}

func (obj *Square) Draw() *ebiten.Image {
	img := ebiten.NewImage(int(obj.Size.X), int(obj.Size.Y))
	revSpeed := 1 - obj.Room.Square.Speed
	img.Fill(color.RGBA{255 - uint8(145*revSpeed), 255 - uint8(255*revSpeed), 255 - uint8(255*revSpeed), 255})
	return img
}
