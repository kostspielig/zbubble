package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

const (
	WinX = 320
	WinY = 240
)

type player struct {
	facingRight bool
	posX        float64
	posY        float64
}

type state struct {
	player player
}

var (
	blue   = color.NRGBA{0x00, 0xad, 0xef, 0xff}
	orange = color.NRGBA{0xff, 0x69, 0x00, 0xff}
)

var (
	imagePlayer     *ebiten.Image
	imagePlayerFlip *ebiten.Image
	imageArrow      *ebiten.Image
)

func NewGame() state {
	return state{player{true, 64, 200}}
}

func init() {
	var err error
	imagePlayer, _, err = ebitenutil.NewImageFromFile("images/monkey_s.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	imagePlayerFlip, _, err = ebitenutil.NewImageFromFile("images/monkey_sf.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	imageArrow, _, err = ebitenutil.NewImageFromFile("images/arrow_s.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
}

func (self *state) Update(screen *ebiten.Image) error {
	// Fill the screen with #FF0000 color
	screen.Fill(blue)

	floor, _ := ebiten.NewImage(WinX, 40, ebiten.FilterNearest)
	floor.Fill(orange)

	fopts := &ebiten.DrawImageOptions{}
	fopts.GeoM.Translate(0, 224)

	screen.DrawImage(floor, fopts)

	opts := &ebiten.DrawImageOptions{}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if self.player.posX < WinX-30 {
			self.player.posX += 2
		}
		self.player.facingRight = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if self.player.posX > 10 {
			self.player.posX -= 2
		}
		self.player.facingRight = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		// TODO : show arrow
	}

	opts.GeoM.Translate(self.player.posX, self.player.posY)
	if self.player.facingRight {
		screen.DrawImage(imagePlayerFlip, opts)
	} else {
		screen.DrawImage(imagePlayer, opts)
	}
	if err := ebitenutil.DebugPrint(screen, "zBubble"); err != nil {
		return err
	}
	return nil
}
