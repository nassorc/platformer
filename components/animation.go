package components

import (
	"platformer/assets"

	"github.com/yohamta/donburi"
)

type AnimationData struct {
  CurrentFrame int
  LastUpdateTime float32
  Data *assets.Animation 
}

var Animation = donburi.NewComponentType[AnimationData]() 
