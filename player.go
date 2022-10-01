package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PlayerHeight = 20
	PlayerWidth  = 20
)

type Player struct {
	X        float64
	Y        float64
	Life     int
	Shield   bool
	Velocity int
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.Y < ScreenHeight-PlayerHeight {
		p.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.Y > 35 {
		p.Y -= 1
	}
}
