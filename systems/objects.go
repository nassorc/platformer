package systems

import (
	"platformer/components"
	dresolv "platformer/resolv"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateObjects(ecs *ecs.ECS) {
	components.Object.Each(ecs.World, func(e *donburi.Entry) {
		obj := dresolv.GetObject(e)
		obj.Update()
	})
}
