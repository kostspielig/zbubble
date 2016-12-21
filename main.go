package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/kostspielig/zbubble/game"
)

func main() {
	state := game.NewGame()

	if err := ebiten.Run(state.Update, game.WinX, game.WinY, 3, "Hello zalandooo!"); err != nil {
		log.Fatal(err)
	}
}
