package features

import (
	"image/color"
	"math"
	"platformer/components"
	"platformer/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	dmath "github.com/nassorc/go-codebase/lib/math"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
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
  query *donburi.Query
}

func NewCameraFollowPlayer(camera *Camera) CameraFollowPlayer {
  query := donburi.NewQuery(filter.Contains(components.Player))

  return CameraFollowPlayer{ camera, camera.Zoom, query }
}

func Lerp(x, y, val float64) float64 {
  return (y-x) * val + x
}

func (c *CameraFollowPlayer) Update(ecs *ecs.ECS) {
  entry, _ := c.query.First(ecs.World)
  player := components.Player.Get(entry)
  obj := components.Object.Get(entry)

  min, max := obj.Shape.Bounds()
  ohw := float32(max.X() - min.X()) / 2
  ohh := float32(max.Y() - min.Y()) / 2

  // center camera at target's x and y
  x := float32(obj.X)
  y := float32(obj.Y)

  // get viewport width and height, then normalize it by the zoom factor
  sw := c.camera.Viewport.Max.X / float32(c.camera.Zoom)
	sh := c.camera.Viewport.Max.Y / float32(c.camera.Zoom)

  // subtract half width and height to get the camera position centered at the target
	x -= sw / 2 - ohw
	y -= sh / 2 - ohh

  trapOffset := float32(16)

  // camera trap
  if !player.FacingRight {
    trapOffset *= -1
  }

  lerpX := 0.08 // [0.0, 1.0]
  lerpY := 0.08 // [0.0, 1.0]
  trapX := float32(x + trapOffset)
  cameraX := float32(Lerp(float64(c.camera.Position.X), float64(trapX), lerpX))
  cameraY := float32(Lerp(float64(c.camera.Position.Y), float64(y), lerpY))

  if !player.FacingRight && x < cameraX {
    cameraX = float32(Lerp(float64(cameraX), float64(x), lerpX))
  }

  if player.FacingRight && x > cameraX {
    cameraX = float32(Lerp(float64(cameraX), float64(x), lerpX))
  }

  c.camera.Position.X = cameraX
  c.camera.Position.Y = cameraY

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

type DebugCamera struct {
  camera *Camera
}

func NewDebugCamera(camera *Camera) DebugCamera {
  return DebugCamera{ camera }
}

func (c *DebugCamera) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
  settings := components.MustFindSettings(ecs.World)
  if settings.Debug {
    center := c.camera.Viewport.Center()
    drawColor := color.RGBA{20, 255, 20, 255}
    vector.DrawFilledCircle(c.camera.View, center.X, center.Y, 4, drawColor, false)
  }
}
