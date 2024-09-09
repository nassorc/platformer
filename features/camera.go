package features

import (
	"fmt"
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

  // make camera target player position
  x := float32(obj.X)
  y := float32(obj.Y)
  sw := c.camera.Viewport.Max.X / float32(c.camera.Zoom)
	sh := c.camera.Viewport.Max.Y / float32(c.camera.Zoom)

  // subtract player's object half size to make origin center
	x -= sw / 2 - 8
	y -= sh / 2 - 10

  trapOffset := 0

  // camera trap
  if player.FacingRight {
    trapOffset = 25
  } else {
    trapOffset = -25
  }

  // Lerp(c.camera.Position.X, x+float32(old), 0.5)

  // follow player 
  // c.camera.Position.X = float32(Lerp(float64(x), float64(x + float32(trapOffset)), 0.1))

  // finalX := playerPos + trapoffset
  // cameraX = Lerp(cameraX, finalX, 0.5)

  // maybe twice the trapoffset
  // newFinalX = float32(c.camera.Position.X + float32(trapOffset))

  finalX := float32(obj.X + float64(trapOffset))
  cameraX := Lerp(float64(c.camera.Position.X), float64(finalX), 0.05)

  // c.camera.Position.X = x + float32(trapOffset)
  c.camera.Position.X = float32(cameraX)
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
    drawColor := color.RGBA{255, 20, 20, 255}
    fmt.Println("drawing camera @", center.X, center.Y)
    vector.DrawFilledCircle(screen, center.X, center.Y, 4, drawColor, false)
  }
}
