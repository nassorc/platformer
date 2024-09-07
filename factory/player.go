package factory

import (
	"platformer/archetypes"
	"platformer/assets"
	"platformer/components"
	dresolv "platformer/resolv"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

// handles the details
func CreatePlayer(ecs *ecs.ECS) *donburi.Entry {
	player := archetypes.Player.Spawn(ecs)

  // add object
  // width and height probably match the tile width and height of the animation
	obj := resolv.NewObject(32, 128, 16, 20)
	dresolv.SetObject(player, obj)
	components.Player.SetValue(player, components.PlayerData{
		FacingRight: true,
	})

	obj.SetShape(resolv.NewRectangle(0, 0, 24, 24))

  // add animation
  components.Animation.SetValue(player, components.AnimationData{
    CurrentFrame: 0,
    LastUpdateTime: 0,
    Data: &assets.DinoGreenIdle,
  })

	return player
}
