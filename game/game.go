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

func (self *player) moveRight() {
	if self.posX < WinX-30 {
		self.posX += 2
	}
	self.facingRight = true
}

func (self *player) moveLeft() {
	if self.posX > 10 {
		self.posX -= 2
	}
	self.facingRight = false
}

func (self *state) throwArrow(p player) {
	self.arrow = append(self.arrow, &arrow{p.posX + 8, WinY - 30})
}

func (self *state) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		self.player.moveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		self.player.moveLeft()
	}

	keyUp := ebiten.IsKeyPressed(ebiten.KeyUp)
	if !self.player.lastKeyUp && keyUp {
		self.throwArrow(self.player)
	}
	self.player.lastKeyUp = keyUp
}

func (self *state) updateArrows() {
	arrows := self.arrow[:0]
	for _, arrow := range self.arrow {
		if dead := arrow.update(); !dead {
			arrows = append(arrows, arrow)
		}
	}
	self.arrow = arrows
}

func (self *state) draw(screen *ebiten.Image) {
	screen.Fill(blue)

	// Draw arrows
	for _, arrow := range self.arrow {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Translate(arrow.posX, arrow.posY)
		screen.DrawImage(imageArrow, o)
	}

	// Draw main player
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(self.player.posX, self.player.posY)
	if self.player.facingRight {
		screen.DrawImage(imagePlayerFlip, opts)
	} else {
		screen.DrawImage(imagePlayer, opts)
	}

	// Draw floor
	fopts := &ebiten.DrawImageOptions{}
	fopts.GeoM.Translate(0, 224)
	screen.DrawImage(imageFloor, fopts)

	ebitenutil.DebugPrint(screen, "zBubble")
}

func (self *state) Update(screen *ebiten.Image) error {
	self.handleInput()
	self.updateArrows()
	self.draw(screen)

	return nil
}
