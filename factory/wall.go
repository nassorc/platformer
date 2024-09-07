package factory

import (
	"platformer/archetypes"
	dresolv "platformer/resolv"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateWall(ecs *ecs.ECS, obj *resolv.Object) *donburi.Entry {
	wall := archetypes.Wall.Spawn(ecs)
	dresolv.SetObject(wall, obj)
	return wall
}
