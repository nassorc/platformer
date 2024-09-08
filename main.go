package main

import (
	_ "embed"

	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"platformer/assets"
	"platformer/config"
	// "platformer/fonts"
	"platformer/scenes"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	bounds image.Rectangle
	scene  Scene
}

func NewGame() *Game {
  assets.MustLoadAssets()

	g := &Game{
		bounds: image.Rectangle{},
		// scene:  &scenes.PlatformerScene{},
		scene:  scenes.NewPlatformScene(),
	}

	return g
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func main() {
	ebiten.SetWindowSize(config.C.Width, config.C.Height)
	ebiten.SetWindowResizable(false)
	rand.Seed(time.Now().UTC().UnixNano())
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
