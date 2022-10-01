package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	AsteroidHeight  = 40
	AsteroidWidth   = 40
	DestroyableZone = -75
)

type Asteroid struct {
	X           float64
	Y           float64
	Velocity    float64
	Destroyable bool
	Colliding   bool
}

type Asteroids struct {
	obstacles []Asteroid
}

type Point struct {
	X float64
	Y float64
}

func (ast *Asteroid) calculateEuclidianDistance(p Point) float64 {
	return math.Sqrt(math.Pow(p.X-ast.X, 2) + math.Pow(p.Y-ast.Y, 2))
}

func (ast *Asteroid) isColliding() {
	var A, B, C, D Point
	A = Point{X: player.X, Y: player.Y}
	B = Point{X: player.X + PlayerWidth, Y: player.Y}
	C = Point{X: player.X + PlayerWidth, Y: player.Y + PlayerHeight}
	D = Point{X: player.X, Y: player.Y + PlayerHeight}

	distances := []float64{ast.calculateEuclidianDistance(A),
		ast.calculateEuclidianDistance(B),
		ast.calculateEuclidianDistance(C),
		ast.calculateEuclidianDistance(D)}

	for _, el := range distances {
		if el < (AsteroidHeight / 2) {
			ast.Colliding = true
			ast.Destroyable = true
			player.Life--
			fmt.Println("Collision")
		}
	}
}

func (a *Asteroid) isDestroyable() {
	if a.X < DestroyableZone {
		a.Destroyable = true
	}
}

func (a *Asteroid) Update() {
	a.X -= a.Velocity
	a.isDestroyable()
	a.isColliding()
}

func (a *Asteroids) Update() {

	for i := 0; i < len(a.obstacles); i++ {
		a.obstacles[i].Update()

		if a.obstacles[i].Destroyable {
			a.obstacles = append(a.obstacles[:i], a.obstacles[i+1:]...) // Deleting destroyable asteroids
		}
	}
}

func (a *Asteroids) Generate() {

	var randomYValue, randomXValue, randomVelocity int
	var x Asteroid

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)

	asteroidGenerated := randomGenerator.Intn(10) + 5

	for i := 0; i < asteroidGenerated; i++ {
		randomYValue = randomGenerator.Intn(ScreenHeight-AsteroidHeight-18) + 18 + AsteroidHeight/2
		randomXValue = randomGenerator.Intn(200) + ScreenWidth
		randomVelocity = randomGenerator.Intn(5) + 1
		x = Asteroid{X: float64(randomXValue), Y: float64(randomYValue), Velocity: float64(randomVelocity)}
		a.obstacles = append(a.obstacles, x)
	}
}

func (a *Asteroids) Draw(screen *ebiten.Image) {

	for i := 0; i < len(a.obstacles); i++ {
		//ebitenutil.DrawRect(screen, float64(a.obstacles[i].X), float64(a.obstacles[i].Y), AsteroidWidth, AsteroidHeight, color.Opaque)
		ebitenutil.DrawCircle(screen, float64(a.obstacles[i].X), float64(a.obstacles[i].Y), AsteroidHeight/2, color.Opaque)
		ebitenutil.DrawCircle(screen, float64(a.obstacles[i].X), float64(a.obstacles[i].Y), 1, color.Black)
	}

}
