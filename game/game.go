package game

import (
	"image/color"
	"strconv"

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
	posX, posY  float64
}

type arrow struct {
	posX, posY float64
}

type bubble struct {
	posX, posY     float64
	speedX, speedY float64
	size           int
}

type state struct {
	player  player
	arrows  []*arrow
	bubbles []*bubble
}

var (
	blue   = color.NRGBA{0x00, 0xad, 0xef, 0xff}
	orange = color.NRGBA{0xff, 0x69, 0x00, 0xff}

	imagePlayer     *ebiten.Image
	imagePlayerFlip *ebiten.Image
	imageArrow      *ebiten.Image
	imageFloor      *ebiten.Image
	imageBubble     []*ebiten.Image
)

func NewGame() state {
	return state{
		player{true, false, 64, 200},
		make([]*arrow, 0, 10),
		[]*bubble{{100, 120, 1, 0, 0}, {200, 120, 1, 0, 0}},
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

	imageBubble = make([]*ebiten.Image, 2)
	for k, _ := range imageBubble {
		imageBubble[k], _, err = ebitenutil.NewImageFromFile("images/box_"+strconv.Itoa(k)+".png", ebiten.FilterNearest)
		if err != nil {
			panic(err)
		}
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
	self.arrows = append(self.arrows, &arrow{p.posX + 8, WinY - 30})
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
	arrows := self.arrows[:0]
	for _, arrow := range self.arrows {
		if dead := arrow.update(); !dead {
			arrows = append(arrows, arrow)
		}
	}
	self.arrows = arrows
}

func (self *state) updateBubbles() {
	for _, bubble := range self.bubbles {
		bubble.posX += bubble.speedX
		bubble.posY += bubble.speedY
		if bubble.posY >= self.player.posY {
			bubble.speedY = -bubble.speedY
		} else {
			bubble.speedY += 0.2
		}
		if bubble.posX >= WinX-27 || bubble.posX <= 0 {
			bubble.speedX = -bubble.speedX
		}
	}
}

func (self *state) draw(screen *ebiten.Image) {
	screen.Fill(blue)

	// Draw arrows
	for _, arrow := range self.arrows {
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

	// Draw bubbles
	for _, bubble := range self.bubbles {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Translate(bubble.posX, bubble.posY)
		screen.DrawImage(imageBubble[bubble.size], o)
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
	self.updateBubbles()
	self.draw(screen)

	return nil
}
