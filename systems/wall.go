package systems

import (
	"image/color"

	"platformer/assets"
	dresolv "platformer/resolv"
	"platformer/tags"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawWall(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Wall.Each(ecs.World, func(e *donburi.Entry) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{255, 60, 60, 255}
		ebitenutil.DrawRect(screen, o.X, o.Y, o.W, o.H, drawColor)
	})
}

func DrawLevel(ecs *ecs.ECS, screen *ebiten.Image) {
  opt := &ebiten.DrawImageOptions{}
  screen.DrawImage(assets.PlatformLevel.Background, opt)
}
