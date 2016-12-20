package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/kostspielig/zbubble/game"
)

func main() {
	if err := ebiten.Run(game.Update, game.WinX, game.WinY, 2, "Hello zalandooo!"); err != nil {
		log.Fatal(err)
	}
}
