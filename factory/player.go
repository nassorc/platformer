package factory

import (
	"platformer/archetypes"
	"platformer/assets"
	"platformer/components"
	"platformer/config"
	dresolv "platformer/resolv"

	"github.com/nassorc/go-codebase/lib/math"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type AnimationState = string

const (
	Idle      AnimationState = "Idle"
	Run       AnimationState = "Run"
	Jump      AnimationState = "Jump"
	WallClimb AnimationState = "WallClimb"
)

// handles the details
func CreatePlayer(ecs *ecs.ECS) *donburi.Entry {
	player := archetypes.Player.Spawn(ecs)

	// add object
	// width and height probably match the tile width and height of the animation
	obj := resolv.NewObject(config.C.SpawnX, config.C.SpawnY, 8, 16)
	dresolv.SetObject(player, obj)
	components.Player.SetValue(player, components.PlayerData{
		FacingRight: true,
	})

	obj.SetShape(resolv.NewRectangle(0, 0, 8, 16))

	components.Animation.SetValue(player, components.AnimationData{
		Offset:       math.Vec2{X: -8, Y: -5},
		CurrentState: Idle, // initial state
		StateMap: map[string]*assets.Animation{
			Idle:      &assets.DinoGreenIdle,
			Run:       &assets.DinoGreenRun,
			Jump:      &assets.DinoGreenJumpAsc,
			WallClimb: &assets.DinoGreenWallClimb,
		},
	})

	return player
}
