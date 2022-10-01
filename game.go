package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ScreenHeight = 480
	ScreenWidth  = 640
)

type Game struct{}

var (
	mplusNormalFont font.Face
	startTime       time.Time
	player          Player
	asteroids       Asteroids
	elapsed         time.Duration
	counter         float64
)

func init() {

	data, err := os.ReadFile("./assets/Arcade.ttf")
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    2,
		DPI:     1080,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	startTime = time.Now()

	player = Player{X: 40, Y: 80, Life: 3, Shield: false, Velocity: 0}

	asteroids.Generate()

	counter = 0
}

func (g *Game) Update() error {

	player.Update()
	asteroids.Update()

	elapsed = time.Since(startTime)
	if math.Mod(elapsed.Seconds(), 10*counter) < 1 {
		asteroids.Generate()
		fmt.Println("hello")
		counter++
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	elapsed := time.Since(startTime).Seconds()

	// Statusbar Draw
	msg := fmt.Sprintf("Life: %d", player.Life)
	text.Draw(screen, msg, mplusNormalFont, ScreenWidth-100, 30, color.White)
	msg = fmt.Sprintf("Time: %0.0f", elapsed)
	text.Draw(screen, msg, mplusNormalFont, 10, 30, color.White)
	ebitenutil.DrawRect(screen, 0, 30, ScreenWidth, 3, color.Opaque)

	// Player Draw
	ebitenutil.DrawRect(screen, float64(player.X), float64(player.Y), PlayerWidth, PlayerHeight, color.Opaque)

	// Asteroids Draw
	asteroids.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
