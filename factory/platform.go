package factory

import (
	"platformer/archetypes"
	"platformer/components"
	dresolv "platformer/resolv"

	"github.com/solarlune/resolv"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlatform(ecs *ecs.ECS, object *resolv.Object) *donburi.Entry {
	platform := archetypes.Platform.Spawn(ecs)
	dresolv.SetObject(platform, object)

	return platform
}

func CreateFloatingPlatform(ecs *ecs.ECS, object *resolv.Object) *donburi.Entry {
	platform := archetypes.FloatingPlatform.Spawn(ecs)
	dresolv.SetObject(platform, object)

	// The floating platform moves using a *gween.Sequence sequence of tweens, moving it back and forth.
	tw := gween.NewSequence()
	obj := components.Object.Get(platform)
	tw.Add(
		gween.New(float32(obj.Y), float32(obj.Y-128), 2, ease.Linear),
		gween.New(float32(obj.Y-128), float32(obj.Y), 2, ease.Linear),
	)
	components.Tween.Set(platform, tw)

	return platform
}
