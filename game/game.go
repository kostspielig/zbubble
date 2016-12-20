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

var (
	blue   = color.NRGBA{0x00, 0xad, 0xef, 0xff}
	orange = color.NRGBA{0xff, 0x69, 0x00, 0xff}

	posX float64 = 64
	posY float64 = 200
)

var (
	imagePlayer        *ebiten.Image
	imagePlayerFlip    *ebiten.Image
	currentImagePlayer *ebiten.Image
	imageArrow         *ebiten.Image
)

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
	currentImagePlayer = imagePlayer

	imageArrow, _, err = ebitenutil.NewImageFromFile("images/arrow_s.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
}

func Update(screen *ebiten.Image) error {
	// Fill the screen with #FF0000 color
	screen.Fill(blue)

	floor, _ := ebiten.NewImage(WinX, 40, ebiten.FilterNearest)
	floor.Fill(orange)

	fopts := &ebiten.DrawImageOptions{}
	fopts.GeoM.Translate(0, 224)

	screen.DrawImage(floor, fopts)

	opts := &ebiten.DrawImageOptions{}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if posX < WinX-30 {
			posX += 2
		}
		currentImagePlayer = imagePlayerFlip
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if posX > 10 {
			posX -= 2
		}
		currentImagePlayer = imagePlayer
	}

	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		// TODO : show arrow
	}

	opts.GeoM.Translate(posX, posY)
	screen.DrawImage(currentImagePlayer, opts)

	if err := ebitenutil.DebugPrint(screen, "zBubble"); err != nil {
		return err
	}
	return nil
}
