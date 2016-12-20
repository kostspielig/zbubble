package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

const (
	WinX = 500
	WinY = 400
)

var (
	blue   = color.NRGBA{0x00, 0xad, 0xef, 0xff}
	orange = color.NRGBA{0xff, 0x69, 0x00, 0xff}

	posX float64 = 64
	posY float64 = 330
)

var imagePlayer *ebiten.Image

func init() {
	var err error
	imagePlayer, _, err = ebitenutil.NewImageFromFile("images/monkey_s.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
}

func Update(screen *ebiten.Image) error {
	// Fill the screen with #FF0000 color
	screen.Fill(blue)
	// Create an 16x16 image
	floor, _ := ebiten.NewImage(WinX, 40, ebiten.FilterNearest)

	// Fill the square with the white color
	floor.Fill(orange)

	// Create an empty option struct
	fopts := &ebiten.DrawImageOptions{}
	fopts.GeoM.Translate(0, 360)

	screen.DrawImage(floor, fopts)

	opts := &ebiten.DrawImageOptions{}

	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		posX += 2
	}

	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		posX -= 2
	}

	opts.GeoM.Translate(posX, posY)

	// Draw the square image to the screen with an empty option
	screen.DrawImage(imagePlayer, opts)

	if err := ebitenutil.DebugPrint(screen, "zBubble"); err != nil {
		return err
	}
	return nil
}
