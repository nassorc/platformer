package features

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nassorc/go-codebase/lib/math"
)

type Camera struct {
  Zoom float64
  View *ebiten.Image
  Viewport math.AABB
  Position math.Vec2
}

func (c Camera) Draw(world *ebiten.Image, screen *ebiten.Image) {
  c.View.Clear()

	cameraOps := &ebiten.DrawImageOptions{}
  // move camera in game world
	cameraOps.GeoM.Translate(-float64(c.Position.X), -float64(c.Position.Y))
	cameraOps.GeoM.Scale(c.Zoom, c.Zoom)

	viewportOps := &ebiten.DrawImageOptions{}
  viewportOps.GeoM.Translate(float64(c.Viewport.Min.X), float64(c.Viewport.Min.Y))

  // use camera's view to draw the world
  c.View.DrawImage(world, cameraOps)
  // draw camera on screen
	screen.DrawImage(c.View, viewportOps)
}
