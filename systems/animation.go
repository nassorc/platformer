package systems

import (
	"platformer/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type Animation struct {
	tick int
}

func NewAnimation() Animation {
	return Animation{
		tick: 0,
	}
}

func (sys *Animation) UpdateAnimation(ecs *ecs.ECS) {
	components.Animation.Each(ecs.World, func(e *donburi.Entry) {
		anim := components.Animation.Get(e)
		anim.NextFrame(float32(sys.tick))
	})
	sys.tick += 1
}

func DrawAnimation(ecs *ecs.ECS, screen *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(components.Object, components.Animation))
	query.Each(ecs.World, func(e *donburi.Entry) {
		// extract data
		anim := components.Animation.Get(e)
		sprite := anim.CurrentSprite() // sprite is a sub image of the texture based off the current frame
		offset := anim.Offset
		entityObj := components.Object.Get(e)
		player := components.MustFindPlayer(ecs.World) // Will only work with the player entity
		ops := &ebiten.DrawImageOptions{}

		// Reflecting the sprite horizontally when the player is facing to the left will flip the image at
		// origin which is the top left corner. This will offset the sprite by -(SpriteWidth) so we must
		// translate the sprite back to its original position.

		// Since we only have one moveable entity in the game world, we can directly use the player entry to
		// determine the direction it's facing.
		if !player.FacingRight {
			ops.GeoM.Scale(-1, 1)
			ops.GeoM.Translate(float64(anim.CurrentSprite().Bounds().Dx()), 0)
		}

		// set the position of the sprite to be the object's position
		ops.GeoM.Translate(float64(entityObj.X)+float64(offset.X), float64(entityObj.Y)+float64(offset.Y))

		screen.DrawImage(sprite, ops)
	})
}
