package systems

import (
	"platformer/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

func DrawLevel(ecs *ecs.ECS, screen *ebiten.Image) {
  opt := &ebiten.DrawImageOptions{}
  screen.DrawImage(assets.PlatformLevel.Background, opt)
}
