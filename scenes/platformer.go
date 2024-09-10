package scenes

import (
	"image/color"
	"sync"

	"platformer/config"
	"platformer/factory"
	"platformer/features"
	"platformer/layers"
	dresolv "platformer/resolv"
	"platformer/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/nassorc/go-codebase/lib/math"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type PlatformerScene struct {
	ecs  *ecs.ECS
	once sync.Once
  camera features.Camera
  worldView *ebiten.Image
}

func NewPlatformScene() *PlatformerScene {
  return &PlatformerScene{
    camera: features.Camera{
      Zoom: 1,
      View: ebiten.NewImage(config.C.Width, config.C.Height),
      Viewport: math.AABB{
        Min: math.Vec2{X: 0, Y: 0},
        Max: math.Vec2{X: float32(config.C.Width), Y: float32(config.C.Height)},
      },
    },
  }
}

func (ps *PlatformerScene) Update() {
	ps.once.Do(ps.configure)
	ps.ecs.Update()
}

func (ps *PlatformerScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255})
  ps.worldView.Clear()
	ps.ecs.Draw(ps.worldView)
  ps.camera.Draw(ps.worldView, screen)
}

func (ps *PlatformerScene) configure() {
  // camera view
  ps.worldView = ebiten.NewImage(config.C.Width, config.C.Height)

	lecs := ecs.NewECS(donburi.NewWorld())
  animation := systems.NewAnimation()
  cameraFollowPlayer := features.NewCameraFollowPlayer(&ps.camera)
  debugCamera := features.NewDebugCamera(&ps.camera)

	lecs.AddSystem(systems.UpdateFloatingPlatform)
	lecs.AddSystem(systems.UpdatePlayer)
	lecs.AddSystem(systems.UpdateObjects)
	lecs.AddSystem(cameraFollowPlayer.Update)
	lecs.AddSystem(animation.UpdateAnimation)
	lecs.AddSystem(systems.UpdateSettings)

	lecs.AddRenderer(layers.Default, systems.DrawWall)
	lecs.AddRenderer(layers.Default, systems.DrawPlatform)
	lecs.AddRenderer(layers.Default, systems.DrawRamp)
	lecs.AddRenderer(layers.Default, systems.DrawFloatingPlatform)
	lecs.AddRenderer(layers.Default, systems.DrawPlayer)
	lecs.AddRenderer(layers.Default, systems.DrawAnimation)
	lecs.AddRenderer(layers.Default, systems.DrawDebug)
	lecs.AddRenderer(layers.Default, systems.DrawHelp)
	lecs.AddRenderer(layers.Top, debugCamera.Draw)
  lecs.AddRenderer(layers.Bottom, func (ecs *ecs.ECS, screen *ebiten.Image) {
    vector.StrokeRect(screen, 0, 0, 32, 32, 2, color.White, false)
  })


	ps.ecs = lecs

	gw, gh := float64(config.C.Width), float64(config.C.Height)

	// Define the world's Space. Here, a Space is essentially a grid (the game's width and height, or 640x360), made up of 16x16 cells. Each cell can have 0 or more Objects within it,
	// and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects. This is a broad, simplified approach to collision
	// detection.
	space := factory.CreateSpace(ps.ecs)

	dresolv.Add(space,
		// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells,
		// as it all is in this platformer example.
		factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, 16, gh, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-16, 0, 16, gh, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, gw, 16, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(0, gh-24, gw, 32, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(160, gh-56, 160, 32, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(320, 64, 32, 160, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(64, 128, 16, 160, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, 64, 128, 16, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, gh-88, 128, 16, "solid")),
		// Create the Player. NewPlayer adds it to the world's Space.
		factory.CreatePlayer(ps.ecs),
		// Non-moving floating Platforms.
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64, 48, 8, "platform")),
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64+64, 48, 8, "platform")),
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64+128, 48, 8, "platform")),
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64+192, 48, 8, "platform")),
		// Create the floating platform.
		factory.CreateFloatingPlatform(ps.ecs, resolv.NewObject(128, gh-32, 128, 8, "platform")),
    factory.CreateFloatingPlatform(ps.ecs, resolv.NewObject(410, 210, 16, 16, "platform")),
    // factor.CreateFloatingPlatform(ps.ecs, resolv)
		// A ramp, which is unique as it has a non-rectangular shape. For this, we will specify a different shape for collision testing.
		factory.CreateRamp(ps.ecs, resolv.NewObject(320, gh-56, 64, 32, "ramp")),
	)
}
