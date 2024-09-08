package features

import (
	"fmt"
  "math"
	"image/color"
	"platformer/components"
	"platformer/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	dmath "github.com/nassorc/go-codebase/lib/math"
	"github.com/yohamta/donburi/ecs"
)

type Camera struct {
  Zoom float64
  View *ebiten.Image
  Viewport dmath.AABB
  Position dmath.Vec2
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

type CameraFollowPlayer struct {
  camera *Camera
  // Different from camera.Zoom
  // used to implement smooth zooming
  Zoom float64
}

func NewCameraFollowPlayer(camera *Camera) CameraFollowPlayer {
  return CameraFollowPlayer{ camera, camera.Zoom }
}

func (c *CameraFollowPlayer) Update(ecs *ecs.ECS) {
  player := components.MustFindPlayer(ecs.World)
  obj := components.Object.Get(player)

  // make camera target player position
  x := float32(obj.X)
  y := float32(obj.Y)
  sw := c.camera.Viewport.Max.X / float32(c.camera.Zoom)
	sh := c.camera.Viewport.Max.Y / float32(c.camera.Zoom)

	x -= sw / 2
	y -= sh / 2

  c.camera.Position.X = x
  c.camera.Position.Y = y

  // bound camera
	if c.camera.Position.X < 0 {
		c.camera.Position.X = 0
	}
	if (c.camera.Position.X + sw) > float32(config.C.Width) {
		c.camera.Position.X = float32(config.C.Width) - sw
	}
	if c.camera.Position.Y < 0 {
		c.camera.Position.Y = 0
	}
	if (c.camera.Position.Y + sh) > float32(config.C.Height) {
		c.camera.Position.Y = float32(config.C.Height) - sh
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		c.Zoom += 1
		c.camera.Zoom = math.Pow(1.01, c.Zoom)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		c.Zoom -= 1
		c.camera.Zoom = math.Pow(1.01, c.Zoom)
	}
}
