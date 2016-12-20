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
	background = color.NRGBA{0xff, 0x69, 0x00, 0xff}

	posX float64 = 64
	posY float64 = 360
)

func Update(screen *ebiten.Image) error {
	// Fill the screen with #FF0000 color
	screen.Fill(background)
	// Create an 16x16 image
	square, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)

	// Fill the square with the white color
	square.Fill(color.White)

	// Create an empty option struct
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
	screen.DrawImage(square, opts)

	if err := ebitenutil.DebugPrint(screen, "zBubble"); err != nil {
		return err
	}
	return nil
}
