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
	facingRight  bool
	lastKeyUp    bool
	lastKeySpace bool
	posX, posY   float64
}

type arrow struct {
	posX, posY float64
}

type bubble struct {
	posX, posY     float64
	speedX, speedY float64
	kind           int
}

type bubbleKind struct {
	next        int
	imageIndex  int
	scale       float64
	size        float64
	speedX      float64
	gravity     float64
	bounce      float64
	spawnBounce float64
}

type state struct {
	player   player
	arrows   []*arrow
	bubbles  []*bubble
	gameOver bool
	pause    bool
}

var (
	blue   = color.NRGBA{0x00, 0xad, 0xef, 0xff}
	orange = color.NRGBA{0xff, 0x69, 0x00, 0xff}

	imagePlayer     *ebiten.Image
	imagePlayerFlip *ebiten.Image
	imageArrow      *ebiten.Image
	imageFloor      *ebiten.Image
	imageBubble     []*ebiten.Image

	bubbleKinds []bubbleKind = []bubbleKind{
		{-1, 0, 1, 27, 1, 0.2, 5, 5},
		{0, 0, 2, 54, 1, 0.2, 7, 5},
		{-1, 1, 1, 25, 1, 0.2, 6, 5},
		{2, 1, 2, 50, 1, 0.2, 8, 5},
	}

	keyMap map[string]ebiten.Key = map[string]ebiten.Key{
		"pause": ebiten.KeySpace,
		"right": ebiten.KeyRight,
		"left":  ebiten.KeyLeft,
		"shoot": ebiten.KeyUp,
	}
)

func newBubble(kind int, x, y float64, dir float64) *bubble {
	k := bubbleKinds[kind]
	return &bubble{x, y, dir * k.speedX, -k.spawnBounce, kind}
}

func NewGame() state {
	return state{
		player{true, false, false, 64, 200},
		make([]*arrow, 0, 10),
		[]*bubble{
			newBubble(0, 100, 120, 1),
			newBubble(1, 200, 120, -1),
			newBubble(2, 150, 120, 1),
			newBubble(3, 250, 120, -1),
		},
		false,
		false,
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
	if !self.pause && ebiten.IsKeyPressed(keyMap["right"]) {
		self.player.moveRight()
	}

	if !self.pause && ebiten.IsKeyPressed(keyMap["left"]) {
		self.player.moveLeft()
	}

	keyUp := ebiten.IsKeyPressed(keyMap["shoot"])
	if !self.player.lastKeyUp && keyUp {
		self.throwArrow(self.player)
	}
	self.player.lastKeyUp = keyUp

	keySpace := ebiten.IsKeyPressed(keyMap["pause"])
	if !self.player.lastKeySpace && keySpace {
		self.pause = !self.pause
	}
	self.player.lastKeySpace = keySpace
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
		k := bubbleKinds[bubble.kind]
		bubble.posX += bubble.speedX
		bubble.posY += bubble.speedY
		if bubble.posY >= self.player.posY {
			bubble.speedY = -k.bounce
		} else {
			bubble.speedY += k.gravity
		}
		if bubble.posX >= WinX-k.size || bubble.posX <= 0 {
			bubble.speedX = -bubble.speedX
		}
	}
}

func rectangleCollision(r1x, r1y, r1w, r1h, r2x, r2y, r2w, r2h float64) bool {
	if r1x < r2x+r2w &&
		r1x+r1w > r2x &&
		r1y < r2y+r2h &&
		r1h+r1y > r2y {
		return true // collision detected!
	}
	return false
}

func (self *state) detectCollisions(screen *ebiten.Image) {
	// Collision between arrows and bubbles
	bubbles := self.bubbles[:0]
	newBubbles := make([]*bubble, 0)
	for _, b := range self.bubbles {
		k := bubbleKinds[b.kind]
		collided := false
		arrows := self.arrows[:0]
		for _, a := range self.arrows {
			if rectangleCollision(b.posX, b.posY, k.size, k.size, a.posX, a.posY, 7, WinY) {
				collided = true
			} else {
				arrows = append(arrows, a)
			}
		}
		self.arrows = arrows
		if !collided {
			bubbles = append(bubbles, b)
		} else if k.next != -1 {
			newBubbles = append(newBubbles,
				newBubble(k.next, b.posX, b.posY, -1),
				newBubble(k.next, b.posX, b.posY, 1))
		}
	}
	bubbles = append(bubbles, newBubbles...)
	self.bubbles = bubbles

	// Collision between bubbles and the player
	for _, b := range self.bubbles {
		k := bubbleKinds[b.kind]
		if rectangleCollision(b.posX, b.posY, k.size, k.size, self.player.posX, self.player.posY, 20, 20) {
			self.gameOver = true
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
		k := bubbleKinds[bubble.kind]
		o.GeoM.Translate(-k.size/k.scale, -k.size/k.scale)
		o.GeoM.Scale(k.scale, k.scale)
		o.GeoM.Translate(k.size, k.size/k.scale)
		o.GeoM.Translate(bubble.posX, bubble.posY)
		screen.DrawImage(imageBubble[k.imageIndex], o)
	}

	// Draw floor
	fopts := &ebiten.DrawImageOptions{}
	fopts.GeoM.Translate(0, 224)
	screen.DrawImage(imageFloor, fopts)

	ebitenutil.DebugPrint(screen, "zBubble")
}

func (self *state) Update(screen *ebiten.Image) error {

	if !self.gameOver {
		self.handleInput()

		if !self.pause {
			self.updateArrows()
			self.updateBubbles()
			self.detectCollisions(screen)
		}
	}
	self.draw(screen)

	if self.gameOver {
		ebitenutil.DebugPrint(screen, "gameOver")
	}
	return nil
}
