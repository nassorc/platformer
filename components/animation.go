package components

import (
	"platformer/assets"

	"github.com/nassorc/go-codebase/lib/math"
	"github.com/yohamta/donburi"
)

type AnimationData struct {
  CurrentFrame int
  LastUpdateTime float32
  Offset math.Vec2
  CurrentState string
  StateMap map[string]*assets.Animation
}

func (anim *AnimationData) MustChangeState(state string) {
  if _, ok := anim.StateMap[state]; !ok {
    panic("Animation state not found")
  }
  if state == anim.CurrentState {
    return
  }
  anim.CurrentState = state
  anim.CurrentFrame = 0
  anim.LastUpdateTime = 0
}

var Animation = donburi.NewComponentType[AnimationData]() 
