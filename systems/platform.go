package systems

import (
	"image/color"

	"platformer/components"
	dresolv "platformer/resolv"
	"platformer/tags"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateFloatingPlatform(ecs *ecs.ECS) {
	tags.FloatingPlatform.Each(ecs.World, func(e *donburi.Entry) {
		tw := components.Tween.Get(e)
		// Platform movement needs to be done first to make sure there's no space between the top and the player's bottom; otherwise, an alternative might
		// be to have the platform detect to see if the Player's resting on it, and if so, move the player up manually.
		y, _, seqDone := tw.Update(1.0 / 60.0)

		obj := dresolv.GetObject(e)
		obj.Y = float64(y)
		if seqDone {
			tw.Reset()
		}
	})
}

func DrawPlatform(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Platform.Each(ecs.World, func(e *donburi.Entry) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{180, 100, 0, 255}
		vector.DrawFilledRect(screen, float32(o.X), float32(o.Y), float32(o.W), float32(o.H), drawColor, false)
	})
}

func DrawFloatingPlatform(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.FloatingPlatform.Each(ecs.World, func(e *donburi.Entry) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{180, 100, 0, 255}
		vector.DrawFilledRect(screen, float32(o.X), float32(o.Y), float32(o.W), float32(o.H), drawColor, false)
	})
}
