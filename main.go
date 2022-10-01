package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Ludum Date 51 - Every Ten Seconds")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
