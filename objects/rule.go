package objects

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// Rule is a rule
type Rule struct {
	Room          *Room
	SndMouseClick *audio.Player
	Score         int
}

func (obj *Rule) Create() {
	defaultContext := audio.NewContext(44100)

	soundFile, err := ebitenutil.OpenFile("media/snd_mouse_click.ogg")
	if err != nil {
		log.Fatal(err)
	}

	soundDecoded, err := vorbis.Decode(defaultContext, soundFile)
	if err != nil {
		log.Fatal(err)
	}

	obj.SndMouseClick, err = audio.NewPlayer(defaultContext, soundDecoded)
	if err != nil {
		log.Fatal(err)
	}
}

func (obj *Rule) Update() {
	// obj.Position.X += obj.Speed.X
	// obj.Position.Y += obj.Speed.Y
	// if (obj.Position.X < 0) || (obj.Position.X+obj.Size.X > obj.Room.LayoutSize.X) {
	// 	obj.Speed.X = -obj.Speed.X
	// }
	// if (obj.Position.Y < 0) || (obj.Position.Y+obj.Size.Y > obj.Room.LayoutSize.Y) {
	// 	obj.Speed.Y = -obj.Speed.Y
	// }
	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(ebiten.MouseButtonLeft)) {
	// 	mouseX, mouseY := ebiten.CursorPosition()
	// 	if (mouseX >= int(obj.Position.X)) && (mouseX <= int(obj.Position.X+obj.Size.X)) &&
	// 		(mouseY >= int(obj.Position.Y)) && (mouseY <= int(obj.Position.Y+obj.Size.Y)) {
	// 		obj.Create()
	// 	}
	// }
}

func (obj *Rule) Draw() *ebiten.Image {
	font := basicfont.Face7x13
	txt := fmt.Sprintf("%d", obj.Score)
	rect := text.BoundString(font, txt)

	img := ebiten.NewImage(rect.Dx(), rect.Dy())
	img.Fill(color.RGBA{10, 10, 10, 255})
	text.Draw(img, txt, font, 0, rect.Dy(), color.White)
	return img
}
