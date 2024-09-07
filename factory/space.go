package factory

import (
	"platformer/archetypes"
	"platformer/components"
	"platformer/config"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSpace(ecs *ecs.ECS) *donburi.Entry {
	space := archetypes.Space.Spawn(ecs)

	cfg := config.C
	spaceData := resolv.NewSpace(cfg.Width, cfg.Height, 16, 16)
	components.Space.Set(space, spaceData)

	return space
}
