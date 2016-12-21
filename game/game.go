package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	WinX = 320
	WinY = 240
)

type player struct {
	facingRight bool
	lastKeyUp   bool
	posX        float64
	posY        float64
}

type arrow struct {
	posX float64
	posY float64
}

type state struct {
	player player
	arrow  []*arrow
}

var (
	blue   = color.NRGBA{0x00, 0xad, 0xef, 0xff}
	orange = color.NRGBA{0xff, 0x69, 0x00, 0xff}

	imagePlayer     *ebiten.Image
	imagePlayerFlip *ebiten.Image
	imageArrow      *ebiten.Image
	imageFloor      *ebiten.Image
)

func NewGame() state {
	return state{
		player{true, false, 64, 200},
		make([]*arrow, 0, 10),
	}
}

func (self *arrow) update() bool {
	self.posY -= 2
	return self.posY <= 0
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

	imageFloor, _ = ebiten.NewImage(WinX, 40, ebiten.FilterNearest)
	imageFloor.Fill(orange)
}

func (self *state) Update(screen *ebiten.Image) error {
	screen.Fill(blue)

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

	keyUp := ebiten.IsKeyPressed(ebiten.KeyUp)
	if !self.player.lastKeyUp && keyUp {
		self.arrow = append(self.arrow, &arrow{self.player.posX + 8, WinY - 30})
	}
	self.player.lastKeyUp = keyUp

	arrows := self.arrow[:0]
	for _, arrow := range self.arrow {
		dead := arrow.update()
		if !dead {
			arrows = append(arrows, arrow)
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Translate(arrow.posX, arrow.posY)
			screen.DrawImage(imageArrow, o)
		}
	}
	self.arrow = arrows

	opts.GeoM.Translate(self.player.posX, self.player.posY)
	if self.player.facingRight {
		screen.DrawImage(imagePlayerFlip, opts)
	} else {
		screen.DrawImage(imagePlayer, opts)
	}

	fopts := &ebiten.DrawImageOptions{}
	fopts.GeoM.Translate(0, 224)
	screen.DrawImage(imageFloor, fopts)
	return nil
}
